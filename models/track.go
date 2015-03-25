package models

import (
	"database/sql"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

type Track struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}

func GetTrackById(id int64) (*Track, error) {
	db, err := sql.Open("mysql", beego.AppConfig.String("dbSource"))
	// utils.checkErr(err)
	defer db.Close()

	out, err := db.Prepare("SELECT id, title FROM radio_tracks WHERE id=?")
	checkErr(err)
	defer out.Close()

	track := new(Track)
	err = out.QueryRow(id).Scan(&track.Id, &track.Title)

	return track, nil
}

func GetTracksByRadioId(id int64) ([]*Track, error) {
	db, err := sql.Open("mysql", beego.AppConfig.String("dbSource"))
	// utils.checkErr(err)
	defer db.Close()

	out, err := db.Query("SELECT id, title FROM radio_tracks WHERE radio_id=? LIMIT 10", id)
	checkErr(err)
	defer out.Close()

	var title string

	tracks := make([]*Track, 0)

	for out.Next() {
		err := out.Scan(&id, &title)
		checkErr(err)
		tracks = append(tracks, &Track{id, title})
	}

	return tracks, nil
}
