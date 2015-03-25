package controllers

import (
	"strconv"

	"github.com/astaxie/beego"
	"muzza.server/models"
)

type RadioTracksController struct {
	beego.Controller
}

func (c *RadioTracksController) Get() {
	id := c.Ctx.Input.Param(":id")
	intid, _ := strconv.ParseInt(id, 10, 64)

	tracks, err := models.GetTracksByRadioId(intid)
	checkErr(err)

	c.Data["json"] = &tracks
	c.ServeJson()
}
