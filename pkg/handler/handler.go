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
	router.GET("/from/:from", h.GetTransactionByUserFrom)
	router.GET("/block/:block", h.GetTransactionByBlock)
	router.GET("/to/:to", h.GetTransactionByUserTo)
	router.GET("/timestamp/:tm", h.GetTransactionByTimestamp)
	return router
}
