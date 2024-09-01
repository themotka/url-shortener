package handlers

import (
	"github.com/gin-gonic/gin"
	"themotka/shortener/internal/url"
)

type Handler struct {
	service url.Service
}

func NewHandler(service *url.Service) *Handler {
	return &Handler{service: *service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	routerGroup := router.Group("/")
	routerGroup.POST("/post", h.PostUrlAndGetKey)
	routerGroup.GET("/:id", h.GetUrlByKey)
	return router
}
