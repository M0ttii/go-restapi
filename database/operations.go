package database

import (
	"database/sql"
	"fmt"
	"go-restapi/structs"
	"log"
)

func (db Database) CreateSQLVideo(name string, uploaded string) {
	stmt, err := db.SqlDB.Prepare("INSERT INTO videos(name, uploaded, show) VALUES (?,?,?)")
	if err != nil {
		log.Panic(err)
	}
	_, err = stmt.Exec(name, uploaded, false)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Video created")
}

func (db Database) GetSQLVideo(rowid int) structs.Video {
	var video structs.Video
	sqlStmt := "SELECT rowid, name, uploaded, show FROM videos WHERE rowid=$1"
	var name string
	var uploaded string
	var show bool

	row := db.SqlDB.QueryRow(sqlStmt, rowid)
	switch err := row.Scan(&rowid, &name, &uploaded, &show); err {
	case sql.ErrNoRows:
		fmt.Println("Video with id " + string(rowid) + " does not exist.")
	case nil:
		video = structs.Video{
			Name:     name,
			Uploaded: uploaded,
			Show:     show,
		}
		return video
	default:
		log.Panic(err)
	}
	return video

}
