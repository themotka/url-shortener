package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.New()
	router.GET("/", getUrlHandler)
	router.POST("/", postUrlHandler)
	err := router.Run(":8080")
	if err != nil {
		return
	}
}

func getUrlHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GET request received",
	})
}

func postUrlHandler(c *gin.Context) {
	var data map[string]interface{}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "POST request received",
		"data":    data,
	})
}
