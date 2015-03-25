package controllers

import (
	"muzza.server/models"

	"github.com/astaxie/beego"
)

type GenreController struct {
	beego.Controller
}

func (c *GenreController) Get() {
	genres, err := models.GetAllGenres()
	checkErr(err)

	c.Data["json"] = &genres
	c.ServeJson()
}
