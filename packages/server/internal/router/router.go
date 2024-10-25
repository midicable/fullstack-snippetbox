package router

import (
	"fullstack-snippetbox-backend/internal/snippets/handler"
	"fullstack-snippetbox-backend/internal/snippets/repository"
	"fullstack-snippetbox-backend/internal/snippets/service"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup) {
	SnippetsRegister(router)
}

func SnippetsRegister(router *gin.RouterGroup) {
	db := 
  snippetsRepo := repository.NewSnippetsRepository("")
	snippetsService := service.NewSnippetsService(snippetsRepo)
	snippetsHandler := handler.NewSnippetsHandler(snippetsService)

	SnippetsGroup := router.Group("/snippets")
	
		// DSGroup.GET("/list", snippetsHandler.GetSnippet())
	SnippetsGroup.GET("/:id", snippetsHandler.GetSnippet)
}
