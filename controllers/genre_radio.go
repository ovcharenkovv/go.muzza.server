package controllers

import (
	"strconv"

	"muzza.server/models"

	"github.com/astaxie/beego"
)

type GenreRadioController struct {
	beego.Controller
}

func (c *GenreRadioController) Get() {
	id := c.Ctx.Input.Param(":id")
	intid, _ := strconv.ParseInt(id, 10, 64)

	genre, err := models.GetGenreByID(intid)
	checkErr(err)

	radios, err := models.GetRadiosByGenre(*genre)
	checkErr(err)

	c.Data["json"] = &radios
	c.ServeJson()
}
