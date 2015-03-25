package controllers

import (
	"muzza.server/models"

	"github.com/astaxie/beego"
)

type GenreController struct {
	beego.Controller
}

type GetGenreController struct {
	beego.Controller
}

func (c *GenreController) Get() {
	genres, err := models.GetAllGenres()
	checkErr(err)

	c.Data["json"] = &genres
	c.ServeJson()
}

func (c *GetGenreController) Get() {
	name := c.Ctx.Input.Param(":name")

	genres, err := models.GetGenreByName(name)
	checkErr(err)

	c.Data["json"] = &genres
	c.ServeJson()
}
