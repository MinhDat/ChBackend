package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/gin-gonic/gin/binding"

	helper "ChGo/helpers"
	"ChGo/models"
)

func GetProducts(c *gin.Context) {

	var products []models.Product
	db := helper.GetDB()
	db.Find(&products)
	c.JSON(200, products)
}

func CreateProduct(c *gin.Context) {
	var product models.Product
	var db = helper.GetDB()

	if err := c.BindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	db.Create(&product)
	c.JSON(http.StatusOK, &product)
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	db := helper.GetDB()
	if err := db.Where("id = ?", id).First(&product).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.BindJSON(&product)
	db.Save(&product)
	c.JSON(http.StatusOK, &product)
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	db := helper.GetDB()

	if err := db.Where("id = ?", id).First(&product).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	db.Delete(&product)
}
