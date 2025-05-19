package devices

import (
	"go-api/pkgs/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DeviceRegisterLayer interface {
	GetDeviceByID(c *gin.Context)
	DeviceRegister(c *gin.Context)
	GetDevices(c *gin.Context)
	UpdateDeviceDetails(c *gin.Context)
}

type deviceObj struct {	
	db *gorm.DB	
}
func DeviceService(db *gorm.DB) DeviceRegisterLayer {	
	return &deviceObj{db : db}
}


func (s *deviceObj) GetDeviceByID(c *gin.Context) {
	logger.Log().Info("GetDeviceByID API called")
	c.String(http.StatusOK, "GetDeviceByID!")
}
func (s *deviceObj) DeviceRegister(c *gin.Context) {
	logger.Log().Info("DeviceRegister API called")
	c.String(http.StatusOK, "DeviceRegister!")
}
func (s *deviceObj) GetDevices(c *gin.Context) {
	logger.Log().Info("GetDevices API called")
	c.String(http.StatusOK, "GetDevices!")
}
func (s *deviceObj) UpdateDeviceDetails(c *gin.Context) {
	logger.Log().Info("UpdateDeviceByID API called")
	c.String(http.StatusOK, "UpdateDeviceByID!")
}
