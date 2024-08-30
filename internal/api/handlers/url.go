package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"themotka/shortener/internal/api/entities"
)

func (h *Handler) GetUrlHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GET request received",
	})
}

func (h *Handler) PostUrlHandler(c *gin.Context) {
	var input entities.InputEntity
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "POST request received",
		"data":    input.Data,
	})
}
