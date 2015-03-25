package main

import (
	"muzza.server/controllers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
	}

	beego.Router("/genres", &controllers.GenreController{})
	beego.Router("/genres/:id/radios", &controllers.GenreRadioController{})

	beego.Router("/radios/:id", &controllers.RadioController{})
	beego.Router("/radios/:id/tracks", &controllers.RadioTracksController{})

	beego.Run()
}
