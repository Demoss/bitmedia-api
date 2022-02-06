package handler

import (
	"github.com/bitmedia-api/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/hash/:h", h.GetTransactionByHash)
	router.GET("/from/:from", h.GetTransactionsByUserFrom)
	router.GET("/block/:block", h.GetTransactionsByBlock)
	router.GET("/to/:to", h.GetTransactionsByUserTo)
	router.GET("/timestamp/:tm", h.GetTransactionsByTimestamp)
	return router
}
