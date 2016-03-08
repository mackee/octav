package main

import (
	"bytes"
	"errors"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/lestrrat/go-pdebug"
)

var ErrAnnotatedStructNotFound = errors.New("annotated struct was not found")

func snakeCase(s string) string {
	ret := []rune{}
	wasLower := false
	for len(s) > 0 {
		r, n := utf8.DecodeRuneInString(s)
		if r == utf8.RuneError {
			panic("yikes")
		}

		s = s[n:]
		if unicode.IsUpper(r) {
			if wasLower {
				ret = append(ret, '_')
			}
			wasLower = false
		} else {
			wasLower = true
		}

		ret = append(ret, unicode.ToLower(r))
	}
	return string(ret)
}

type Processor struct {
	Types []string
	Dir   string
}

func skipGenerated(fi os.FileInfo) bool {
	switch {
	case strings.HasSuffix(fi.Name(), "gendb.go"):
		return false
	case strings.HasSuffix(fi.Name(), "_gen.go"):
		return false
	}
	return true
}

func (p *Processor) Do() error {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, p.Dir, skipGenerated, parser.ParseComments)
	if err != nil {
		return err
	}

	if len(pkgs) == 0 {
		return errors.New("no packages to process...")
	}

	for _, pkg := range pkgs {
		if err := p.ProcessPkg(pkg); err != nil {
			return err
		}
	}

	return nil
}

func (p *Processor) ShouldProceed(s Struct) bool {
	if len(p.Types) == 0 {
		return true
	}

	for _, t := range p.Types {
		if t == s.Name {
			return true
		}
	}
	return false
}

func (p *Processor) ProcessPkg(pkg *ast.Package) error {
	if pdebug.Enabled {
		g := pdebug.Marker("ProcessPkg %s", pkg.Name)
		defer g.End()
	}
	for fn, f := range pkg.Files {
		pdebug.Printf("Checking file %s", fn)
		for _, s := range p.ExtractStructs(pkg, f) {
			if pdebug.Enabled {
				pdebug.Printf("Checking struct %s", s.Name)
			}
			if !p.ShouldProceed(s) {
				if pdebug.Enabled {
					pdebug.Printf("Skipping struct %s", s.Name)
				}
				continue
			}

			if err := p.ProcessStruct(s); err != nil {
				return err
			}
		}
	}
	return nil
}

func (p *Processor) ProcessStruct(s Struct) error {
	if pdebug.Enabled {
		g := pdebug.Marker("ProcessStruct %s", s.Name)
		defer g.End()
	}

	buf := bytes.Buffer{}

	buf.WriteString("// Automatically generated by gendb utility. DO NOT EDIT!\n")
	buf.WriteString("package ")
	buf.WriteString(s.PackageName)
	buf.WriteString("\n\n")
	buf.WriteString("\nimport (")
	buf.WriteString("\n" + `"errors"`)
	buf.WriteString("\n)")
	varname := unicode.ToLower(rune(s.Name[0]))

	scols := bytes.Buffer{}
	icols := bytes.Buffer{}
	sfields := bytes.Buffer{}   // Scan fields
	ifields := bytes.Buffer{}   // Scan fields (everything except for OID)
	ipholders := bytes.Buffer{} // Insert place holders (everything except for OID)

	hasEID := false
	sansOidFields := make([]StructField, 0, len(s.Fields)-1)
	for _, f := range s.Fields {
		if f.Name == "OID" {
			continue
		}
		if f.Name == "EID" {
			hasEID = true
		}
		sansOidFields = append(sansOidFields, f)
	}

	for i, f := range sansOidFields {
		icols.WriteString(f.ColumnName)
		ipholders.WriteRune('?')
		ifields.WriteRune(varname)
		ifields.WriteRune('.')
		ifields.WriteString(f.Name)
		if i < len(sansOidFields)-1 {
			icols.WriteString(", ")
			ipholders.WriteString(", ")
			ifields.WriteString(", ")
		}
	}

	for i, f := range s.Fields {
		scols.WriteString(f.ColumnName)
		sfields.WriteRune('&')
		sfields.WriteRune(varname)
		sfields.WriteRune('.')
		sfields.WriteString(f.Name)
		if i < len(s.Fields)-1 {
			scols.WriteString(", ")
			sfields.WriteString(", ")
		}
	}

	fmt.Fprintf(&buf, "\nvar %sTable = %s\n", s.Name, strconv.Quote(s.Tablename))
	fmt.Fprintf(&buf, "\nfunc (%c *%s) Scan(scanner interface { Scan(...interface{}) error }) error {", varname, s.Name)
	fmt.Fprintf(&buf, "\nreturn scanner.Scan(%s)", sfields.String())
	buf.WriteString("\n}\n")

	if hasEID {
		fmt.Fprintf(&buf, "\nfunc (%c *%s) LoadByEID(tx *Tx, eid string) error {", varname, s.Name)
		fmt.Fprintf(&buf, "\nrow := tx.QueryRow(`SELECT %s FROM ` + %sTable + ` WHERE eid = ?`, eid)", scols.String(), s.Name)
		fmt.Fprintf(&buf, "\nif err := %c.Scan(row); err != nil {", varname)
		buf.WriteString("\nreturn err")
		buf.WriteString("\n}")
		buf.WriteString("\nreturn nil")
		buf.WriteString("\n}\n")
	}

	fmt.Fprintf(&buf, "\nfunc (%c *%s) Create(tx *Tx) error {", varname, s.Name)
	if hasEID {
		fmt.Fprintf(&buf, "\nif %c.EID == "+`""`+" {", varname)
		buf.WriteString("\n" + `return errors.New("create: non-empty EID required")`)
		buf.WriteString("\n}\n\n")
	}
	fmt.Fprintf(&buf, "\nresult, err := tx.Exec(`INSERT INTO ` + %sTable + ` (%s) VALUES (%s)`, %s)", s.Name, icols.String(), ipholders.String(), ifields.String())
	buf.WriteString("\nif err != nil {")
	buf.WriteString("\nreturn err")
	buf.WriteString("\n}\n")
	buf.WriteString("\nlii, err := result.LastInsertId()")
	buf.WriteString("\nif err != nil {")
	buf.WriteString("\nreturn err")
	buf.WriteString("\n}\n")
	fmt.Fprintf(&buf, "\n%c.OID = lii", varname)
	buf.WriteString("\nreturn nil")
	buf.WriteString("\n}\n")

	fmt.Fprintf(&buf, "\nfunc (%c %s) Delete(tx *Tx) error {", varname, s.Name)
	fmt.Fprintf(&buf, "\nif %c.OID != 0 {", varname)
	fmt.Fprintf(&buf, "\n_, err := tx.Exec(`DELETE FROM ` + %sTable + ` WHERE oid = ?`, %c.OID)", s.Name, varname)
	buf.WriteString("\nreturn err")
	buf.WriteString("\n}\n")
	if hasEID {
		fmt.Fprintf(&buf, "\nif %c.EID != %s {", varname, `""`)
		fmt.Fprintf(&buf, "\n_, err := tx.Exec(`DELETE FROM ` + %sTable + ` WHERE eid = ?`, %c.EID)", s.Name, varname)
		buf.WriteString("\nreturn err")
		buf.WriteString("\n}\n")
	}

	if hasEID {
		fmt.Fprintf(&buf, "\nreturn errors.New(%s)", strconv.Quote("either OID/EID must be filled"))
		buf.WriteString("\n}\n")
	} else {
		fmt.Fprintf(&buf, "\nreturn errors.New(%s)", strconv.Quote("column OID must be filled"))
		buf.WriteString("\n}\n")
	}

	fsrc, err := format.Source(buf.Bytes())
	if err != nil {
		log.Printf("%s", buf.Bytes())
		return err
	}

	fn := filepath.Join(p.Dir, snakeCase(s.Name)+"_gen.go")
	if pdebug.Enabled {
		pdebug.Printf("Generating file %s", fn)
	}
	fi, err := os.Create(fn)
	if err != nil {
		return err
	}
	defer fi.Close()

	if _, err := fi.Write(fsrc); err != nil {
		return err
	}

	return nil
}

func (p *Processor) ExtractStructs(pkg *ast.Package, f *ast.File) []Struct {
	ctx := &InspectionCtx{
		Marker:  "+DB",
		Package: pkg.Name,
	}

	ast.Inspect(f, ctx.ExtractStructs)
	return ctx.Structs
}

func (c *InspectionCtx) ExtractStructs(n ast.Node) bool {
	var decl *ast.GenDecl
	var ok bool
	var err error

	if decl, ok = n.(*ast.GenDecl); !ok {
		return true
	}

	if decl.Tok != token.TYPE {
		return true
	}

	if err = c.ExtractStructsFromDecl(decl); err != nil {
		return true
	}

	return false
}

func (ctx *InspectionCtx) ExtractStructsFromDecl(decl *ast.GenDecl) error {
	marked := false
	cacheEnabled := true
	noScanner := false
	tablename := ""
	preCreate := ""
	postCreate := ""
	cacheExpires := "1800"
	if cg := decl.Doc; cg != nil {
		for _, c := range cg.List {
			txt := strings.TrimSpace(strings.TrimLeft(c.Text, "//"))
			if strings.HasPrefix(txt, ctx.Marker) {
				marked = true
				tag := reflect.StructTag(strings.TrimPrefix(txt, ctx.Marker))
				tablename = tag.Get("tablename")
				preCreate = tag.Get("pre_create")
				postCreate = tag.Get("post_create")
				if tag.Get("cache") == "false" {
					cacheEnabled = false
				}
				if ce := tag.Get("cache_expires"); ce != "" {
					cacheExpires = ce
				}
				if tag.Get("scanner") == "false" {
					noScanner = true
				}
				break
			}
		}
	}

	if !marked {
		return ErrAnnotatedStructNotFound
	}

	for _, spec := range decl.Specs {
		var t *ast.TypeSpec
		var s *ast.StructType
		var ok bool
		var ident *ast.Ident

		if t, ok = spec.(*ast.TypeSpec); !ok {
			return ErrAnnotatedStructNotFound
		}

		if s, ok = t.Type.(*ast.StructType); !ok {
			return ErrAnnotatedStructNotFound
		}

		if tablename == "" {
			tablename = fmt.Sprintf("%s_%s",
				ctx.Package,
				snakeCase(t.Name.Name),
			)
		}

		st := Struct{
			PackageName:  ctx.Package,
			CacheEnabled: cacheEnabled,
			CacheExpires: cacheExpires,
			Fields:       make([]StructField, 0, len(s.Fields.List)),
			Name:         t.Name.Name,
			NoScanner:    noScanner,
			PreCreate:    preCreate,
			PostCreate:   postCreate,
			Tablename:    tablename,
		}

	LoopFields:
		for _, f := range s.Fields.List {
			autoIncrement := false
			primaryKey := false
			unique := false
			converter := ""
			columnName := ""

			if f.Tag != nil {
				v := f.Tag.Value
				if len(v) >= 2 {
					if v[0] == '`' {
						v = v[1:]
					}
					if v[len(v)-1] == '`' {
						v = v[:len(v)-1]
					}
				}

				tags := reflect.StructTag(v).Get("db")
				for _, tag := range strings.Split(tags, ",") {
					switch tag {
					case "ignore":
						continue LoopFields
					case "auto_increment":
						autoIncrement = true
					case "pk":
						primaryKey = true
					case "unique":
						unique = true
					default:
						switch {
						case strings.HasPrefix(tag, "converter="):
							converter = tag[10:]
						case strings.HasPrefix(tag, "column="):
							columnName = tag[7:]
						}
					}
				}
			}

			if columnName == "" {
				columnName = snakeCase(f.Names[0].Name)
			}

			switch f.Type.(type) {
			case *ast.Ident:
				ident = f.Type.(*ast.Ident)
			case *ast.SelectorExpr:
				ident = f.Type.(*ast.SelectorExpr).Sel
			case *ast.StarExpr:
				ident = f.Type.(*ast.StarExpr).X.(*ast.SelectorExpr).Sel
			default:
				fmt.Printf("%#v\n", f.Type)
				panic("field type not supported")
			}

			field := StructField{
				AutoIncrement: autoIncrement,
				Converter:     converter,
				ColumnName:    columnName,
				Name:          f.Names[0].Name,
				Type:          ident.Name,
				Unique:        unique,
			}

			if autoIncrement {
				st.AutoIncrementField = &field
			}

			if primaryKey {
				st.PrimaryKey = &field
			}

			// if auto_increment, this key does not need to
			// be included in the st.Fields list
			st.Fields = append(st.Fields, field)
		}
		ctx.Structs = append(ctx.Structs, st)
	}

	return nil
}
