package api

import (
	"go-api/pkgs/middleware"
	"go-api/pkgs/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RouterInit(obj service.ServiceGroupLayer, logger *zap.Logger) *gin.Engine {
	router := gin.New()
	router.Use(middleware.CustomLogger(logger))
	router.Use(gin.Recovery())
	router.GET("/health", obj.Status)
	router.GET("/static", obj.Status)
	router.POST("/signup", obj.UserRegister)
	//router.Use(middleware.AuthMiddleware()) // This middleware is not defined in the snippet

	userGroup := router.Group("user")
	{
		userGroup.POST("/login", obj.Login)
		userGroup.POST("/forgot", obj.UserRegister)
		userGroup.POST("/reset", obj.UserRegister)
		userGroup.POST("/logout", obj.UserRegister)
		userGroup.GET("/getdetail/:id", obj.GetUserByID)
		userGroup.PUT("/update", obj.UpdateUserDetails)
		userGroup.DELETE("/deregister", obj.UserRegister)

	}
	deviceGroup := router.Group("devices")
	{
		deviceGroup.POST("/register", obj.DeviceRegister)
		deviceGroup.GET("/details", obj.GetDevices)
		deviceGroup.GET("/details/:id", obj.GetDeviceByID)
		deviceGroup.PUT("/update", obj.DeviceRegister)
		deviceGroup.DELETE("/deregister", obj.DeviceRegister)
	}
	return router
}
