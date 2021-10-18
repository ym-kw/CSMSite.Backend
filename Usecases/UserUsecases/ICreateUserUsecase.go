package UserUsecases

import (
	"CSMSite.Backend/Domain/Dtos"
	"github.com/gin-gonic/gin"
)

type ICreateUserUsecase interface {
	CreateUser(c *gin.Context) (user Dtos.UserResponse, err error)
}
