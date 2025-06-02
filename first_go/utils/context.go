package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func GetUserId(c *gin.Context) (int, error) {

	userID, exist := c.Get("user_id")

	if !exist {
		return 0, errors.New("User ID does not exist")
	}

	id, ok := userID.(int)
	if !ok {
		return 0, errors.New("Invalid ID")
	}

	return id, nil
}
