package controllers

import (
	"strconv"

	"muzza.server/models"

	"github.com/astaxie/beego"
)

type RadioController struct {
	beego.Controller
}

func (c *RadioController) Get() {
	id := c.Ctx.Input.Param(":id")
	intid, _ := strconv.ParseInt(id, 10, 64)

	radio, err := models.GetRadioByID(intid)
	checkErr(err)

	c.Data["json"] = &radio
	c.ServeJson()
}
