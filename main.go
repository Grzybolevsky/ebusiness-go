package main

import (
	"ebusiness-go/src/app"
	"ebusiness-go/src/controllers"
	"ebusiness-go/src/models"
)

func main() {
	models.SetupDb()
	controllers.SetupControllers()
	app.Start(":1323")
}
