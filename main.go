package main

import (
	"database/sql"
	"fmt"
	"go-restapi/database"
	"log"
	"net/http"

	//"go-restapi/structs"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

/*func uploadVideo(w http.ResponseWriter, r *http.Request) {
	//Video Uploading

}*/

func getVideo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Video: %v", vars["id"])
}

func main() {

	//Router
	r := mux.NewRouter()

	//SQLite
	sqlObj, err := sql.Open("sqlite3", "videos.sqlite")
	if err != nil {
		log.Panic(err)
	}
	data := database.Database{
		SqlDB: sqlObj,
	}

	createStmt := "CREATE TABLE IF NOT EXISTS videos (name VARCHAR NOT NULL, uploaded VARCHAR, show BOOLEAN NOT NULL);"
	_, err = data.SqlDB.Exec(createStmt)
	if err != nil {
		log.Panic(err)
	}
	//data.CreateSQLVideo("TestVideo", "1.1.1")
	video := data.GetSQLVideo(1)
	fmt.Println(video.GetName())

	//Endpoints
	//r.HandleFunc("/api/videos", getVideos).Methods("GET")
	r.HandleFunc("/api/video/{id}", getVideo).Methods("GET")
	//r.HandleFunc("/api/video", uploadVideo).Methods("POST")
	//r.HandleFunc("/api/video/{id}", updateVideo).Methods("PUT")
	//r.HandleFunc("/api/video/{id}", deleteVideo).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
