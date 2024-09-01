package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"themotka/shortener/internal/url"
)

func (h *Handler) PostUrlAndGetKey(c *gin.Context) {
	var input url.URL
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	key, err := h.service.GetCurrentOrGenerateKey(input.Data)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, key)
	}
}

func (h *Handler) GetUrlByKey(c *gin.Context) {
	strId := c.Param("id")
	res, err := h.service.GetUrlIfExist(strId)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, res)
	}
}
