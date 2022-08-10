package app

import (
	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Echo = echo.New()
var DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

func Start(address string) {
	CheckError(err, "failed to connect database")
	Echo.Logger.Fatal(Echo.Start(address))
}
