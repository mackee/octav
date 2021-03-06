// Automatically generated by gendb utility. DO NOT EDIT!
package db

import (
	"errors"
	"strconv"
	"time"
)

var RoomTable = "rooms"

func (r *Room) Scan(scanner interface {
	Scan(...interface{}) error
}) error {
	return scanner.Scan(&r.OID, &r.EID, &r.VenueID, &r.Name, &r.Capacity, &r.CreatedOn, &r.ModifiedOn)
}

func (r *Room) LoadByEID(tx *Tx, eid string) error {
	row := tx.QueryRow(`SELECT oid, eid, venue_id, name, capacity, created_on, modified_on FROM `+RoomTable+` WHERE eid = ?`, eid)
	if err := r.Scan(row); err != nil {
		return err
	}
	return nil
}

func (r *Room) Create(tx *Tx) error {
	if r.EID == "" {
		return errors.New("create: non-empty EID required")
	}

	r.CreatedOn = time.Now()
	result, err := tx.Exec(`INSERT INTO `+RoomTable+` (eid, venue_id, name, capacity, created_on, modified_on) VALUES (?, ?, ?, ?, ?, ?)`, r.EID, r.VenueID, r.Name, r.Capacity, r.CreatedOn, r.ModifiedOn)
	if err != nil {
		return err
	}

	lii, err := result.LastInsertId()
	if err != nil {
		return err
	}

	r.OID = lii
	return nil
}

func (r Room) Update(tx *Tx) error {
	if r.OID != 0 {
		_, err := tx.Exec(`UPDATE `+RoomTable+` SET eid = ?, venue_id = ?, name = ?, capacity = ? WHERE oid = ?`, r.EID, r.VenueID, r.Name, r.Capacity, r.OID)
		return err
	}
	if r.EID != "" {
		_, err := tx.Exec(`UPDATE `+RoomTable+` SET venue_id = ?, name = ?, capacity = ? WHERE eid = ?`, r.VenueID, r.Name, r.Capacity, r.EID)
		return err
	}
	return errors.New("either OID/EID must be filled")
}

func (r Room) Delete(tx *Tx) error {
	if r.OID != 0 {
		_, err := tx.Exec(`DELETE FROM `+RoomTable+` WHERE oid = ?`, r.OID)
		return err
	}

	if r.EID != "" {
		_, err := tx.Exec(`DELETE FROM `+RoomTable+` WHERE eid = ?`, r.EID)
		return err
	}

	return errors.New("either OID/EID must be filled")
}

func (v *RoomList) LoadSinceEID(tx *Tx, since string, limit int) error {
	var s int64
	if id := since; id != "" {
		vdb := Room{}
		if err := vdb.LoadByEID(tx, id); err != nil {
			return err
		}

		s = vdb.OID
	}
	return v.LoadSince(tx, s, limit)
}

func (v *RoomList) LoadSince(tx *Tx, since int64, limit int) error {
	rows, err := tx.Query(`SELECT oid, eid, venue_id, name, capacity, created_on, modified_on FROM `+RoomTable+` WHERE oid > ? ORDER BY oid ASC LIMIT `+strconv.Itoa(limit), since)
	if err != nil {
		return err
	}
	res := make([]Room, 0, limit)
	for rows.Next() {
		vdb := Room{}
		if err := vdb.Scan(rows); err != nil {
			return err
		}
		res = append(res, vdb)
	}
	*v = res
	return nil
}
