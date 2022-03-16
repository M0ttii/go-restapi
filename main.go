package main

import (
	"database/sql"
	"fmt"
	"go-restapi/database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type Video struct {
	Name     string `json:"name"`
	Uploaded string `json:"uploaded"`
	Show     bool   `json:"show"`
}

/*func uploadVideo(w http.ResponseWriter, r *http.Request) {
	//Video Uploading

}*/

func getVideo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Video: %v", vars["id"])
}

func main() {
	fmt.Println("dfgdf")

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

	createStmt := "CREATE TABLE IF NOT EXISTS videos (video_id INT PRIMARY KEY, name VARCHAR NOT NULL, uploaded VARCHAR, show BOOLEAN NOT NULL);"
	_, err = data.SqlDB.Exec(createStmt)
	if err != nil {
		log.Panic(err)
	}
	data.CreateVideo("TestVideo", "1.1.1")

	//Endpoints
	//r.HandleFunc("/api/videos", getVideos).Methods("GET")
	r.HandleFunc("/api/video/{id}", getVideo).Methods("GET")
	//r.HandleFunc("/api/video", uploadVideo).Methods("POST")
	//r.HandleFunc("/api/video/{id}", updateVideo).Methods("PUT")
	//r.HandleFunc("/api/video/{id}", deleteVideo).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
