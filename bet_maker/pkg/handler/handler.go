package handler

import (
	"github.com/gin-gonic/gin"
	"microServices/bet_maker/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		events := api.Group("events")
		{
			events.GET("/")
			events.POST("/", h.createEvent)
		}
		bets := router.Group("bets")
		{
			bets.POST("/")
			bets.GET("/")
		}
	}
	return router
}
