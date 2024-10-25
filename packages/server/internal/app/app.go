package app

import (
	"fullstack-snippetbox-backend/internal/middleware"
	"fullstack-snippetbox-backend/internal/router"

	"github.com/gin-gonic/gin"
)


func Start() {
	engine := gin.Default()
	engine.Use(middleware.CORSMiddleware())

	api := engine.Group("/api")
	router.SetupRoutes(api)

	err := engine.Run(":8080")
	if err != nil {
		return
	}
}