// Automatically generated by genmodel utility. DO NOT EDIT!
package octav

import (
	"encoding/json"
	"github.com/builderscon/octav/octav/db"
	"github.com/lestrrat/go-pdebug"
)

func (v Room) GetPropNames() ([]string, error) {
	l, _ := v.L10N.GetPropNames()
	return append(l, "id", "venue_id", "name", "capacity"), nil
}

func (v Room) GetPropValue(s string) (interface{}, error) {
	switch s {
	case "id":
		return v.ID, nil
	case "venue_id":
		return v.VenueID, nil
	case "name":
		return v.Name, nil
	case "capacity":
		return v.Capacity, nil
	default:
		return v.L10N.GetPropValue(s)
	}
}

func (v Room) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["id"] = v.ID
	m["venue_id"] = v.VenueID
	m["name"] = v.Name
	m["capacity"] = v.Capacity
	buf, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return marshalJSONWithL10N(buf, v.L10N)
}

func (v *Room) UnmarshalJSON(data []byte) error {
	m := make(map[string]interface{})
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	if jv, ok := m["id"]; ok {
		switch jv.(type) {
		case string:
			v.ID = jv.(string)
			delete(m, "id")
		default:
			return ErrInvalidFieldType{Field: "id"}
		}
	}

	if jv, ok := m["venue_id"]; ok {
		switch jv.(type) {
		case string:
			v.VenueID = jv.(string)
			delete(m, "venue_id")
		default:
			return ErrInvalidFieldType{Field: "venue_id"}
		}
	}

	if jv, ok := m["name"]; ok {
		switch jv.(type) {
		case string:
			v.Name = jv.(string)
			delete(m, "name")
		default:
			return ErrInvalidFieldType{Field: "name"}
		}
	}

	if jv, ok := m["capacity"]; ok {
		switch jv.(type) {
		case float64:
			v.Capacity = uint(jv.(float64))
			delete(m, "capacity")
		default:
			return ErrInvalidFieldType{Field: "capacity"}
		}
	}

	if err := ExtractL10NFields(m, &v.L10N, []string{"name"}); err != nil {
		return err
	}
	return nil
}

func (v *Room) Load(tx *db.Tx, id string) error {
	vdb := db.Room{}
	if err := vdb.LoadByEID(tx, id); err != nil {
		return err
	}

	if err := v.FromRow(vdb); err != nil {
		return err
	}
	if err := v.LoadLocalizedFields(tx); err != nil {
		return err
	}
	return nil
}

func (v *Room) LoadLocalizedFields(tx *db.Tx) error {
	ls, err := db.LoadLocalizedStringsForParent(tx, v.ID, "Room")
	if err != nil {
		return err
	}

	if len(ls) > 0 {
		v.L10N = LocalizedFields{}
		for _, l := range ls {
			v.L10N.Set(l.Language, l.Name, l.Localized)
		}
	}
	return nil
}

func (v *Room) FromRow(vdb db.Room) error {
	v.ID = vdb.EID
	v.VenueID = vdb.VenueID
	v.Name = vdb.Name
	v.Capacity = vdb.Capacity
	return nil
}

func (v *Room) Create(tx *db.Tx) error {
	if v.ID == "" {
		v.ID = UUID()
	}

	vdb := db.Room{
		EID:      v.ID,
		VenueID:  v.VenueID,
		Name:     v.Name,
		Capacity: v.Capacity,
	}
	if err := vdb.Create(tx); err != nil {
		return err
	}

	if err := v.L10N.CreateLocalizedStrings(tx, "Room", v.ID); err != nil {
		return err
	}
	return nil
}

func (v *Room) Delete(tx *db.Tx) error {
	if pdebug.Enabled {
		g := pdebug.Marker("Room.Delete (%s)", v.ID)
		defer g.End()
	}

	vdb := db.Room{EID: v.ID}
	if err := vdb.Delete(tx); err != nil {
		return err
	}
	if err := db.DeleteLocalizedStringsForParent(tx, v.ID, "Room"); err != nil {
		return err
	}
	return nil
}

func (v *RoomList) Load(tx *db.Tx, since string, limit int) error {
	vdbl := db.RoomList{}
	if err := vdbl.LoadSinceEID(tx, since, limit); err != nil {
		return err
	}
	res := make([]Room, len(vdbl))
	for i, vdb := range vdbl {
		v := Room{}
		if err := v.FromRow(vdb); err != nil {
			return err
		}
		if err := v.LoadLocalizedFields(tx); err != nil {
			return err
		}
		res[i] = v
	}
	*v = res
	return nil
}
