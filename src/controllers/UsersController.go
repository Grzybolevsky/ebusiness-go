package controllers

import (
	. "ebusiness-go/src/app"
	. "ebusiness-go/src/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetUser(context echo.Context) error {
	user := new(User)
	userId := context.Param("userId")
	if result := DB.First(&user, userId); result.Error != nil {
		return context.NoContent(http.StatusNotFound)
	}
	return context.JSON(http.StatusOK, user)
}

func GetAllUsers(context echo.Context) error {
	users := new([]User)
	if result := DB.Find(&users); result.Error != nil {
		return context.NoContent(http.StatusNotFound)
	}
	return context.JSON(http.StatusOK, users)
}

func SaveUser(context echo.Context) error {
	user := new(User)
	if err := context.Bind(user); err != nil {
		return context.JSON(http.StatusBadRequest, "Cannot create user with given data")
	}
	if result := DB.Create(user); result.Error != nil {
		return context.JSON(http.StatusBadRequest, "Cannot create user with given data")
	}
	return context.JSON(http.StatusCreated, user)
}
