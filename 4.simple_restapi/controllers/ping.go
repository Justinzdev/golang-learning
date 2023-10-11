package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PangOut(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
