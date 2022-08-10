package controllers

import (
	. "ebusiness-go/src/app"
	. "ebusiness-go/src/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func GetAllItemsForOrder(context echo.Context) error {
	orderItems := new([]OrderItem)
	orderId := context.Param("orderId")
	DB.Where("OrderId = ?", orderId).Find(orderItems)
	return context.JSON(http.StatusOK, orderItems)
}

func SaveOrderItem(context echo.Context) error {
	order := new(Order)
	orderItem := new(OrderItem)
	orderId := context.Param("orderId")
	if result := DB.First(order, orderId); order.Finished || result.Error != nil {
		return context.JSON(http.StatusBadRequest, "Cannot create order item with given data")
	}
	if err := context.Bind(orderItem); err != nil {
		return context.JSON(http.StatusBadRequest, "Cannot create order item with given data")
	}
	id, errParseInt := strconv.Atoi(orderId)
	if errParseInt != nil {
		return context.JSON(http.StatusBadRequest, "Cannot create order item with given data")
	}
	orderItem.OrderId = id
	if result := DB.Create(orderItem); result.Error != nil {
		return context.JSON(http.StatusBadRequest, "Cannot create order item with given data")
	}
	return context.JSON(http.StatusOK, orderItem)
}
