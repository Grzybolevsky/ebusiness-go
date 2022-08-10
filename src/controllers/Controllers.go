package controllers

import (
	"ebusiness-go/src/app"
	"github.com/labstack/echo/v4"
	"net/http"
)

func SetupControllers() {
	app.Echo.GET("/hello", hello)
	setupOrdersController()
	setupProductsController()
	setupShopsController()
	setupOrderItemsController()
	setupUsersController()
}

func setupOrdersController() {
	app.Echo.GET("/orders/:orderId", GetOneOrder)
	app.Echo.GET("/orders", GetAllOrders)
	app.Echo.GET("/orders/user/:userId", GetOrdersForUser)
	app.Echo.POST("/orders", SaveOrder)
	app.Echo.POST("/orders/:orderId/finish", FinishOrder)
}

func setupOrderItemsController() {
	app.Echo.GET("/orders/:orderId/items", GetAllItemsForOrder)
	app.Echo.POST("/orders/:orderId/items", SaveOrderItem)
}

func setupProductsController() {
	app.Echo.GET("/products/:productId", GetOneProduct)
	app.Echo.GET("/products", GetAllProducts)
	app.Echo.POST("/products", SaveProduct)
}

func setupShopsController() {
	app.Echo.GET("/shops", GetAllShops)
	app.Echo.GET("/shops/:shopId/products", GetProductsForShop)
	app.Echo.POST("/shops", SaveShop)
}

func setupUsersController() {
	app.Echo.GET("/users/:userId", GetUser)
	app.Echo.GET("/users", GetAllUsers)
	app.Echo.POST("/users", SaveUser)
}

func hello(context echo.Context) error {
	return context.JSON(http.StatusOK, "Hello, World!")
}
