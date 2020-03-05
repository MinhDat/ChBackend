package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	helper "ChGo/helpers"
	"ChGo/models"
)

func GetCategories(c *gin.Context) {

	var categorys []models.Category
	db := helper.GetDB()
	db.Find(&categorys)
	c.JSON(200, categorys)
}

func CreateCategory(c *gin.Context) {
	var category models.Category
	var db = helper.GetDB()

	if err := c.BindJSON(&category); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// TODO: Will create category translation in here - next feature

	db.Create(&category)
	c.JSON(http.StatusOK, &category)
}

func UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.Category

	db := helper.GetDB()
	if err := db.Where("uuid = ?", id).First(&category).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.BindJSON(&category)
	db.Save(&category)
	c.JSON(http.StatusOK, &category)
}

func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.Category
	db := helper.GetDB()

	if err := db.Where("uuid = ?", id).First(&category).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	db.Delete(&category)
}
