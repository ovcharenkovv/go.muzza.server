package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"muzza.server/controllers"
)

func main() {

	// CROS FILTERS
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "DELETE", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
	}

	beego.Router("/genres", &controllers.GenreController{})
	beego.Router("/genres/:name", &controllers.GetGenreController{})
	beego.Router("/genres/:id/radios", &controllers.GenreRadioController{})

	beego.Router("/radios/:id", &controllers.RadioController{})
	beego.Router("/radios/:id/tracks", &controllers.RadioTracksController{})

	beego.Run()
}
