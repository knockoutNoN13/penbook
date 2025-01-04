package handler

import (
	"pentbook/pkg/service"

	"github.com/gin-contrib/cors"
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
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"http://127.0.0.1:8080"}, // Разрешенные источники
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		ExposeHeaders: []string{"Content-Length"},
	}))
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
