package database

import (
	"log"
)

func (db Database) CreateVideo(name string, uploaded string) {
	stmt, err := db.SqlDB.Prepare("INSERT INTO videos(name, uploaded, show) VALUES (?,?,?)")
	if err != nil {
		log.Panic(err)
	}
	_, err = stmt.Exec(name, uploaded, false)
	if err != nil {
		log.Panic(err)
	}
}
