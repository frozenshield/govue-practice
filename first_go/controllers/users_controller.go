package controllers

import (
	"net/http"
	"strconv"

	"github.com/frozenshield/first_go/database"
	"github.com/frozenshield/first_go/models"
	"github.com/frozenshield/first_go/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type loginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func CreateUser(c *gin.Context) {
	var user models.Users
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message:": "failed to add"})
		return
	}

	hashPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message:": "Failed to Generate Hash"})
		return
	}
	user.Password = string(hashPass)

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message:": "Failed to Create User"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message:": "Account Created"})
}

func DeleteUser(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message:": "Invalid Input Type"})
		return
	}

	var user models.Users
	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message:": "Record not Found"})
		return
	}

	if err := database.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message:": err.Error()})
		return
	}

	database.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message:": "User Deleted Successfully"})
}

func LoginUser(c *gin.Context) {
	var login loginRequest
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Input"})
		return
	}

	var user models.Users
	if err := database.DB.Where("username = ?", login.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Username Can't Found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Wrong Password"})
		return
	}

	token, err := utils.GenerateJWT(strconv.FormatUint(uint64(user.ID), 10))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to generate"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully Logged in!!",
		"token": token})
}

func UpdateUser(c *gin.Context) {

	id := c.Param("id")
	var user models.Users
	if err := database.DB.Find(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message:": "Id not found"})
		return
	}

	var updateUser models.Users
	if err := c.BindJSON(&updateUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message:": "Error"})
		return
	}

	if err := database.DB.Model(&user).Updates(updateUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message:": "Failed to Update"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message:": "Updated Successfully"})

}
