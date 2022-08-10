package models

import (
	"ebusiness-go/src/app"
)

func SetupDb() {
	err := app.DB.AutoMigrate(&User{})
	app.CheckError(err, "failed to migrate schema")
	err = app.DB.AutoMigrate(&Shop{})
	app.CheckError(err, "failed to migrate schema")
	err = app.DB.AutoMigrate(&Product{})
	app.CheckError(err, "failed to migrate schema")
	err = app.DB.AutoMigrate(&OrderItem{})
	app.CheckError(err, "failed to migrate schema")
	err = app.DB.AutoMigrate(&Order{})
	app.CheckError(err, "failed to migrate schema")
}
