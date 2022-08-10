package controllers

import (
	. "ebusiness-go/src/app"
	. "ebusiness-go/src/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetAllShops(context echo.Context) error {
	shops := new([]Shop)
	if result := DB.Find(&shops); result.Error != nil {
		return context.NoContent(http.StatusNotFound)
	}
	return context.JSON(http.StatusOK, shops)
}

func SaveShop(context echo.Context) error {
	shop := new(Shop)
	if err := context.Bind(shop); err != nil {
		return context.JSON(http.StatusBadRequest, "Cannot create shop with given data")
	}
	if result := DB.Create(shop); result.Error != nil {
		return context.JSON(http.StatusBadRequest, "Cannot create shop with given data")
	}
	return context.JSON(http.StatusOK, shop)
}
