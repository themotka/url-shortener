package handlers

import (
	"github.com/gin-gonic/gin"
	"themotka/shortener/internal/api/middleware"
)

type Router struct {
	hashTable *middleware.HashTable
}

func NewRouter(hashTable *middleware.HashTable) *Router {
	return &Router{hashTable: hashTable}
}

func (h *Router) InitRoutes() *gin.Engine {
	router := gin.New()
	url := router.Group("/")
	url.POST("/post", h.PostUrlHandler)
	if h.hashTable.Repo == nil {
		url.GET("/:id", h.GetUrlHandler)
	} else {
		url.GET("/:id", h.GetUrlHandlerDB)
	}
	return router
}
