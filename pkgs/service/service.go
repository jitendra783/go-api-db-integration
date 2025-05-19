package service

import (
	"go-api/pkgs/logger"
	"go-api/pkgs/service/devices"
	"go-api/pkgs/service/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ServiceGroupLayer interface {
	Status(c *gin.Context)
	user.UserRegisterLayer
	devices.DeviceRegisterLayer
}

type serviceObj struct {
	user.UserRegisterLayer
	devices.DeviceRegisterLayer
}

func NewServiceGroupObject(pgsql *gorm.DB) ServiceGroupLayer {
	return &serviceObj{
		user.UserService(pgsql),
		devices.DeviceService(pgsql),
	}
}

func (s *serviceObj) Status(c *gin.Context) {
	logger.Log().Info("Status API called")
	c.String(http.StatusOK, "Working!")
}
