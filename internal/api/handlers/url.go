package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"themotka/shortener/internal/api/entities"
)

func (h *Handler) GetUrlHandler(c *gin.Context) {
	strId := c.Param("id")
	if _, ok := h.hashTable.Table[strId]; ok {
		c.JSON(http.StatusOK, h.hashTable.Table[strId])
	} else {
		c.JSON(http.StatusNoContent, gin.H{"status": "not found"})
	}
}

func (h *Handler) PostUrlHandler(c *gin.Context) {
	var input entities.InputEntity
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	shortened := h.hashTable.WriteTo(input.Data)
	c.JSON(http.StatusOK, shortened)
}
