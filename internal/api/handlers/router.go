package handlers

import (
	"github.com/gin-gonic/gin"
	"themotka/shortener/internal/api/middleware"
)

type Handler struct {
	hashTable *middleware.HashTable
}

func NewHandler(hashTable *middleware.HashTable) *Handler {
	return &Handler{hashTable: hashTable}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	url := router.Group("/")
	{
		url.POST("/post", h.PostUrlHandler)
		url.GET("/:id", h.GetUrlHandler)
	}
	return router
}
