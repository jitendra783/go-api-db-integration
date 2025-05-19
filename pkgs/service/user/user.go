package user

import (
	"go-api/pkgs/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRegisterLayer interface {
	UserRegister(c *gin.Context)
	GetUserByID(c *gin.Context)
	UpdateUserDetails(c *gin.Context)
	Login(c *gin.Context)
}

type userObj struct {
	db *gorm.DB
}

func UserService(db *gorm.DB) UserRegisterLayer {
	return &userObj{db: db}
}

func (s *userObj) UserRegister(c *gin.Context) {
	logger.Log().Info("UserRegister API called")
	c.String(http.StatusOK, "UserRegister!")

}
func (s *userObj) GetUserByID(c *gin.Context) {
	logger.Log().Info("GetUserByID API called")
	c.String(http.StatusOK, "GetUserByID!")
}
func (s *userObj) Login(c *gin.Context) {
	logger.Log().Info("Login API called")
	c.String(http.StatusOK, "Login!")
}
func (s *userObj) UpdateUserDetails(c *gin.Context) {
	logger.Log().Info("UpdateUserByID API called")
	c.String(http.StatusOK, "UpdateUserByID!")
}
