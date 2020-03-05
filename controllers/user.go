package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	helper "ChGo/helpers"
	"ChGo/models"
)

func Register(c *gin.Context) {
	var user models.User
	var db = helper.GetDB()

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	db.Create(&user)
	c.JSON(http.StatusOK, &user)
}
