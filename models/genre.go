package models

import (
	"database/sql"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

type Genre struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Bg   string `json:"bg"`
}

func GetAllGenres() ([]*Genre, error) {

	db, err := sql.Open("mysql", beego.AppConfig.String("dbSource"))
	checkErr(err)
	defer db.Close()

	out, err := db.Query("SELECT id, name, bg FROM genres ORDER BY name")
	checkErr(err)
	defer out.Close()

	var id int64
	var name, bg string

	genres := make([]*Genre, 0)

	for out.Next() {
		err := out.Scan(&id, &name, &bg)
		checkErr(err)
		genres = append(genres, &Genre{id, name, bg})
	}

	return genres, nil
}

func GetGenreByID(id int64) (*Genre, error) {

	db, err := sql.Open("mysql", beego.AppConfig.String("dbSource"))
	checkErr(err)
	defer db.Close()

	out, err := db.Prepare("SELECT id, name, bg FROM genres WHERE id=?")
	checkErr(err)
	defer out.Close()

	genre := new(Genre)
	err = out.QueryRow(id).Scan(&genre.Id, &genre.Name, &genre.Bg)
	checkErr(err)

	return genre, nil
}

func GetGenreByName(name string) (*Genre, error) {

	db, err := sql.Open("mysql", beego.AppConfig.String("dbSource"))
	checkErr(err)
	defer db.Close()

	out, err := db.Prepare("SELECT id, name, bg FROM genres WHERE name=?")
	checkErr(err)
	defer out.Close()

	genre := new(Genre)
	err = out.QueryRow(name).Scan(&genre.Id, &genre.Name, &genre.Bg)
	checkErr(err)

	return genre, nil
}
