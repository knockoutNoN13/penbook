package handler

import (
	"pentbook/pkg/service"

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

	// ================== AUTH API GROUP =========================
	auth := router.Group("/auth")
	{
		auth.POST("/signin", h.signIn)
	}

	// ================== COMMANDS API GROUP =========================

	commands := router.Group("/commands")
	{
		commands.GET("/getall", h.getAllCommands) // dashboard

		command := commands.Group("/:id")
		{
			command.GET("/", h.getSelectedCommand)               // данные по скану
			command.DELETE("/", h.userIdentity, h.deleteCommand) // снести результаты скана со всеми хостами

		}

		commands.POST("/create", h.userIdentity, h.createCommand)

	}

	return router
}
