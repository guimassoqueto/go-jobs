package router

import (
	"github.com/gin-gonic/gin"
	"github.com/guimassoqueto/go-jobs/handler"
	docs "github.com/guimassoqueto/go-jobs/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type User struct {
	Name string
}

func InitializeRoute(router *gin.Engine) {
	// Inititalizer Handler
	handler.InitializeHandler()

	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	
	v1 := router.Group(basePath)
	{
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningHandler)
	}

	// Swagger route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}