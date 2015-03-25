package models

import (
	"database/sql"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

type Radio struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Genre        string `json:"genre"`
	Shoutcast_id int64  `json:"shoutcast_id"`
	Stream_url   string `json:"stream_url"`
}

func GetRadioByID(id int64) (*Radio, error) {

	db, err := sql.Open("mysql", beego.AppConfig.String("dbSource"))
	checkErr(err)
	defer db.Close()

	out, err := db.Prepare("SELECT id, name, genre, shoutcast_id, stream_url  FROM radios WHERE id=?")
	checkErr(err)
	defer out.Close()

	radio := new(Radio)
	err = out.QueryRow(id).Scan(&radio.Id, &radio.Name, &radio.Genre, &radio.Shoutcast_id, &radio.Stream_url)

	return radio, nil
}

func GetRadiosByGenre(genre Genre) ([]*Radio, error) {

	db, err := sql.Open("mysql", beego.AppConfig.String("dbSource"))
	checkErr(err)
	defer db.Close()

	out, err := db.Query("SELECT id, name, genre, shoutcast_id, stream_url FROM radios WHERE genre=?", genre.Name)
	checkErr(err)
	defer out.Close()

	var id, shId int64
	var name, genreName, url string

	radios := make([]*Radio, 0)

	for out.Next() {
		err := out.Scan(&id, &name, &genreName, &shId, &url)
		checkErr(err)
		radios = append(radios, &Radio{id, name, genreName, shId, url})
	}

	return radios, nil
}
