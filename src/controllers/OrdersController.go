package controllers

import (
	. "ebusiness-go/src/app"
	. "ebusiness-go/src/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetOneOrder(context echo.Context) error {
	order := new(Order)
	orderId := context.Param("orderId")
	if result := DB.First(&order, orderId); result.Error != nil {
		return context.NoContent(http.StatusNotFound)
	}
	return context.JSON(http.StatusOK, order)
}

func GetAllOrders(context echo.Context) error {
	orders := new([]Order)
	if result := DB.Find(&orders); result.Error != nil {
		return context.NoContent(http.StatusNotFound)
	}
	return context.JSON(http.StatusOK, orders)
}

func GetOrdersForUser(context echo.Context) error {
	orders := new([]Order)
	userId := context.Param("userId")
	if result := DB.Where("UserId == ?", userId).Find(&orders); result.Error != nil {
		return context.NoContent(http.StatusNotFound)
	}
	return context.JSON(http.StatusOK, orders)
}

func SaveOrder(context echo.Context) error {
	order := new(Order)
	if err := context.Bind(order); err != nil {
		return context.JSON(http.StatusBadRequest, "Cannot create order with given data")
	}
	if order.Finished {
		return context.JSON(http.StatusBadRequest, "Cannot create order with given data")
	}
	if result := DB.Create(order); result.Error != nil {
		return context.JSON(http.StatusBadRequest, "Cannot create order with given data")
	}
	return context.JSON(http.StatusOK, order)
}

func FinishOrder(context echo.Context) error {
	order := new(Order)
	orderId := context.Param("orderId")
	if result := DB.Find(order, orderId); result.Error != nil {
		return context.JSON(http.StatusBadRequest, "Cannot finish order with given data")
	}
	order.Finished = true
	if result := DB.Save(order); result.Error != nil {
		return context.JSON(http.StatusBadRequest, "Cannot finish order with given data")
	}
	return context.JSON(http.StatusOK, order)
}
