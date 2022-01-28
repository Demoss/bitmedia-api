package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetTransactionByHash(c *gin.Context) {
	hash := c.Param("h")
	item, err := h.services.GetTransactionByHash(c, hash)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "invalid param")
		return
	}
	c.JSON(http.StatusOK, item)
}
func (h *Handler) GetTransactionByUserFrom(c *gin.Context) {
	from := c.Param("from")
	item, err := h.services.GetTransactionByUserFrom(c, from)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "invalid param")
		return
	}
	c.JSON(http.StatusOK, item)
}
func (h *Handler) GetTransactionByBlock(c *gin.Context) {
	block := c.Param("block")
	item, err := h.services.GetTransactionByBlock(c, block)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "invalid param")
		return
	}
	c.JSON(http.StatusOK, item)
}
func (h *Handler) GetTransactionByUserTo(c *gin.Context) {
	to := c.Param("to")
	item, err := h.services.GetTransactionByUserTo(c, to)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "invalid param")
		return
	}
	c.JSON(http.StatusOK, item)
}
func (h *Handler) GetTransactionByTimestamp(c *gin.Context) {
	timestamp := c.Param("tm")
	item, err := h.services.GetTransactionByTimestamp(c, timestamp)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "invalid param")
		return
	}
	c.JSON(http.StatusOK, item)
}
