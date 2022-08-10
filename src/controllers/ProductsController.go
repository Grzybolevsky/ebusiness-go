package controllers

import (
	. "ebusiness-go/src/app"
	. "ebusiness-go/src/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetOneProduct(context echo.Context) error {
	product := new(Product)
	productId := context.Param("productId")
	if result := DB.Where("id = ?", productId).First(&product); result.Error != nil {
		return context.NoContent(http.StatusNotFound)
	}
	return context.JSON(http.StatusOK, product)
}

func GetAllProducts(context echo.Context) error {
	products := new([]Product)
	if result := DB.Find(&products); result.Error != nil {
		return context.NoContent(http.StatusNotFound)
	}
	return context.JSON(http.StatusOK, products)
}

func GetProductsForShop(context echo.Context) error {
	products := new([]Product)
	shopId := context.Param("shopId")
	if result := DB.Where("ShopId = ?", shopId).Find(&products); result.Error != nil {
		return context.NoContent(http.StatusNotFound)
	}
	return context.JSON(http.StatusOK, products)
}

func SaveProduct(context echo.Context) error {
	product := new(Product)
	if err := context.Bind(product); err != nil {
		return context.JSON(http.StatusBadRequest, "Cannot create product with given data")
	}
	if result := DB.Create(product); result.Error != nil {
		return context.JSON(http.StatusBadRequest, "Cannot create product with given data")
	}
	return context.JSON(http.StatusOK, product)
}
