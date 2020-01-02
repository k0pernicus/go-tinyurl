package db

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func AddRecord(db *sql.DB, id string, redirection string, deadline time.Time, hasDeadline bool) error {
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into urls (id,redirection,deadline,has_deadline) values (?,?,?,?)")
	if _, err := stmt.Exec(id, redirection, deadline, hasDeadline); err != nil {
		fmt.Printf("Cannot exec INSERT with id %s: %s\n", id, err.Error())
		return err
	}
	if err := tx.Commit(); err != nil {
		fmt.Printf("Cannot commit INSERT with id %s: %s\n", id, err.Error())
		return err
	}
	return nil
}

func GetRecord(db *sql.DB, ID string) (Record, error) {
	rows, err := db.Query("select * from urls")
	if err != nil {
		fmt.Printf("Cannot exec SELECT: %s\n", err.Error())
		return Record{}, err
	}
	for rows.Next() {
		var tempRecord Record
		var createdDate string
		err = rows.Scan(&tempRecord.ID, &tempRecord.Redirection, &createdDate, &tempRecord.HasDeadline)
		if err != nil {
			fmt.Printf("Cannot exec scan when SELECT: %s\n", err.Error())
			return Record{}, err
		}
		if tempRecord.ID == ID {
			t, _ := time.Parse("2006-01-02 15:04:05", createdDate)
			tempRecord.Deadline = t
			return tempRecord, nil
		}
	}
	return Record{}, errors.New("Not found")
}

func DeleteRecord(db *sql.DB, ID string) {
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("delete from urls where id=?")
	if _, err := stmt.Exec(ID); err != nil {
		fmt.Printf("Cannot exec DELETE with id %s: %s\n", ID, err.Error())
		return
	}
	tx.Commit()
}
