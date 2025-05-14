package controllers

import (
	"net/http"

	"github.com/frozenshield/first_go/database"
	"github.com/frozenshield/first_go/models"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.BindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	database.DB.Create(&post)
	c.JSON(http.StatusOK, post)
}

func UpdatePost(c *gin.Context) {
	var post models.Post
	id := c.Param("id")
	if err := database.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		return
	}

	var UpdatePost models.Post
	if err := c.BindJSON(&UpdatePost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message:": err.Error()})
		return
	}

	database.DB.Model(&post).Updates(UpdatePost)
	c.JSON(http.StatusOK, gin.H{"message:": "Updated Successfully"})
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	if err := database.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message:": err.Error()})
	}

	database.DB.Delete(&post)
	c.JSON(http.StatusOK, gin.H{"message:": "Deleted Successfully"})
}

func GetPosts(c *gin.Context) {
	var post []models.Post
	database.DB.Find(&post)
	c.JSON(http.StatusOK, post)
}

func GetPost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	if err := database.DB.First(&post, id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message:": "Cant Find"})
	}

	c.JSON(http.StatusOK, post)
}
