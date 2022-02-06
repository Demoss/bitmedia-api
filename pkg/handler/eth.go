package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetTransactionByHash(c *gin.Context) {
	hash := c.Param("h")

	item, err := h.services.GetTransactionsByHash(c, hash)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "invalid param")
		return
	}
	c.JSON(http.StatusOK, item)
}
func (h *Handler) GetTransactionsByUserFrom(c *gin.Context) {
	from := c.Param("from")
	page, err := getPage(c)
	item, err := h.services.Eth.GetTransactionsByUserFrom(c, from, int64(page))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "invalid param")
		return
	}
	c.JSON(http.StatusOK, item)
}
func (h *Handler) GetTransactionsByBlock(c *gin.Context) {
	block := c.Param("block")
	page, err := getPage(c)
	item, err := h.services.GetTransactionsByBlock(c, block, int64(page))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "invalid param")
		return
	}
	c.JSON(http.StatusOK, item)
}
func (h *Handler) GetTransactionsByUserTo(c *gin.Context) {
	to := c.Param("to")
	page, err := getPage(c)
	item, err := h.services.GetTransactionsByUserTo(c, to, int64(page))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "invalid param")
		return
	}
	c.JSON(http.StatusOK, item)
}
func (h *Handler) GetTransactionsByTimestamp(c *gin.Context) {
	timestamp := c.Param("tm")
	page, err := getPage(c)
	item, err := h.services.GetTransactionsByTimestamp(c, timestamp, int64(page))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "invalid param")
		return
	}
	c.JSON(http.StatusOK, item)
}

func getPage(c *gin.Context) (int, error) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		return 0, err
	}
	return page, nil
}
