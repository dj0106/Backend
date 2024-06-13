package router

import (
	"InvitKaro/Middleware"
	"InvitKaro/handler"
	"github.com/gin-gonic/gin"
)

func RouteDispatcher() {
	router := gin.Default()

	router.Use(Middleware.CORSMiddleware())

	apis := router.Group("/apis/")
	userApis := apis.Group("/service/v1")

	userApis.Use(Middleware.AuthMiddleware())
	{
		router.POST("/login", handler.LoginHandler)
		router.POST("/register", handler.RegisterHandler)
		//router.GET("/dashboard", handler.Dashboard)
		//router.GET("/nearby", handler.Localities)
		//router.POST("/host", handler.BeHost)
		router.GET("/filters")
		router.POST("/apply/filters", handler.Filter)
		router.POST("/filter/host", handler.GenerateHost)
		router.POST("/generate/host_id", handler.GenerateHostCredential)
		router.GET("/generate/otp")
		router.GET("/type/user")
		//router.GET("/type")
	}

	userApis.Use(Middleware.NoAuth())
	{
		router.GET("/organisations")
		router.GET("/location/organisations")

	}
}
