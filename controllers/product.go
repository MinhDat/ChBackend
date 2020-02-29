package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/gin-gonic/gin/binding"

	"ChGo/db"
	"ChGo/models"
)

func GetProducts(c *gin.Context) {

	var tasks []models.Product
	db := db.GetDB()
	db.Find(&tasks)
	c.JSON(200, tasks)
}

func CreateProduct(c *gin.Context) {
	var task models.Product
	var db = db.GetDB()

	if err := c.BindJSON(&task); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	db.Create(&task)
	c.JSON(http.StatusOK, &task)
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var task models.Product

	db := db.GetDB()
	if err := db.Where("id = ?", id).First(&task).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.BindJSON(&task)
	db.Save(&task)
	c.JSON(http.StatusOK, &task)
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	var task models.Product
	db := db.GetDB()

	if err := db.Where("id = ?", id).First(&task).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	db.Delete(&task)
}
