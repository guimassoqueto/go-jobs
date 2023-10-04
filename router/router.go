package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initializer Router
	router := gin.Default()
	InitializeRoute(router)
	// Init Routes
	router.Run(":8080")
}