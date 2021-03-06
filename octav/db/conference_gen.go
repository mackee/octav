// Automatically generated by gendb utility. DO NOT EDIT!
package db

import (
	"errors"
	"strconv"
	"time"
)

var ConferenceTable = "conferences"

func (c *Conference) Scan(scanner interface {
	Scan(...interface{}) error
}) error {
	return scanner.Scan(&c.OID, &c.EID, &c.Slug, &c.Title, &c.SubTitle, &c.CreatedOn, &c.ModifiedOn)
}

func (c *Conference) LoadByEID(tx *Tx, eid string) error {
	row := tx.QueryRow(`SELECT oid, eid, slug, title, sub_title, created_on, modified_on FROM `+ConferenceTable+` WHERE eid = ?`, eid)
	if err := c.Scan(row); err != nil {
		return err
	}
	return nil
}

func (c *Conference) Create(tx *Tx) error {
	if c.EID == "" {
		return errors.New("create: non-empty EID required")
	}

	c.CreatedOn = time.Now()
	result, err := tx.Exec(`INSERT INTO `+ConferenceTable+` (eid, slug, title, sub_title, created_on, modified_on) VALUES (?, ?, ?, ?, ?, ?)`, c.EID, c.Slug, c.Title, c.SubTitle, c.CreatedOn, c.ModifiedOn)
	if err != nil {
		return err
	}

	lii, err := result.LastInsertId()
	if err != nil {
		return err
	}

	c.OID = lii
	return nil
}

func (c Conference) Update(tx *Tx) error {
	if c.OID != 0 {
		_, err := tx.Exec(`UPDATE `+ConferenceTable+` SET eid = ?, slug = ?, title = ?, sub_title = ? WHERE oid = ?`, c.EID, c.Slug, c.Title, c.SubTitle, c.OID)
		return err
	}
	if c.EID != "" {
		_, err := tx.Exec(`UPDATE `+ConferenceTable+` SET slug = ?, title = ?, sub_title = ? WHERE eid = ?`, c.Slug, c.Title, c.SubTitle, c.EID)
		return err
	}
	return errors.New("either OID/EID must be filled")
}

func (c Conference) Delete(tx *Tx) error {
	if c.OID != 0 {
		_, err := tx.Exec(`DELETE FROM `+ConferenceTable+` WHERE oid = ?`, c.OID)
		return err
	}

	if c.EID != "" {
		_, err := tx.Exec(`DELETE FROM `+ConferenceTable+` WHERE eid = ?`, c.EID)
		return err
	}

	return errors.New("either OID/EID must be filled")
}

func (v *ConferenceList) LoadSinceEID(tx *Tx, since string, limit int) error {
	var s int64
	if id := since; id != "" {
		vdb := Conference{}
		if err := vdb.LoadByEID(tx, id); err != nil {
			return err
		}

		s = vdb.OID
	}
	return v.LoadSince(tx, s, limit)
}

func (v *ConferenceList) LoadSince(tx *Tx, since int64, limit int) error {
	rows, err := tx.Query(`SELECT oid, eid, slug, title, sub_title, created_on, modified_on FROM `+ConferenceTable+` WHERE oid > ? ORDER BY oid ASC LIMIT `+strconv.Itoa(limit), since)
	if err != nil {
		return err
	}
	res := make([]Conference, 0, limit)
	for rows.Next() {
		vdb := Conference{}
		if err := vdb.Scan(rows); err != nil {
			return err
		}
		res = append(res, vdb)
	}
	*v = res
	return nil
}
