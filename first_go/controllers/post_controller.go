package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/frozenshield/first_go/database"
	"github.com/frozenshield/first_go/models"
	"github.com/frozenshield/first_go/utils"
	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	userId, err := utils.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var post models.Post
	err_bind := c.ShouldBindJSON(&post)
	if err_bind != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message:": "invalid JSON format"})
		return
	}

	if post.Title == "" || post.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error:": "Title and Content are required"})
		return
	}

	if database.DB == nil {
		c.JSON(http.StatusBadGateway, gin.H{"error:": "Database COnnection not stablished"})
		return
	}

	query := "INSERT INTO posts (user_id, title, content) VALUES (?, ?, ?)"
	result, err := database.DB.Exec(query, userId, post.Title, post.Content)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to create post"})
		return
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error:": "failed to get last inserted id"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":      insertedID,
		"user_id": userId,
		"title":   post.Title,
		"content": post.Content,
	})

}

// func GetPost(c *gin.Context) {
// 	query := "SELECT * FROM posts"
// 	result, err := database.DB.Exec()
// 	}

func UpdatePost(c *gin.Context) {

	postID := c.Param("id")

	id, err := strconv.Atoi(postID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error:": "Invalid post ID"})
		return
	}

	userID, err := utils.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized access")
		return
	}

	var post models.Post
	post_err := c.ShouldBindJSON(&post)
	if post_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input Data"})
		return
	}

	if post.Content == "" || post.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error:": "Title or content should not be empty"})
		return
	}

	if database.DB == nil {
		c.JSON(http.StatusBadGateway, gin.H{"error:": "No Database Connection"})
		return
	}

	query := "UPDATE posts SET title = ? , content = ? WHERE id = ? AND user_id = ?"
	result, err := database.DB.Exec(query, post.Title, post.Content, id, userID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error:": "failed to update post"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error:": "Failed to verify update"})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error:": "Post not found you dont have oermission"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message:": "Post updated Successfully",
		"Post ID:": id})

}

func GetPosts(c *gin.Context) {

	userID, err := utils.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Unauthorized Account")
		return
	}

	rows, err := database.DB.Query("SELECT id, Title, Content FROM posts WHERE user_id = ?", userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error:": "Failed to get Post"})
		return
	}

	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.Content, &post.Title)
		fmt.Println(err)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error:": "Error Scanning"})
			return
		}

		posts = append(posts, post)
	}

	c.JSON(http.StatusOK, posts)

}
