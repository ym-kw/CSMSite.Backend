package UserInteractor

import (
	"encoding/json"
	"time"

	"CSMSite.Backend/Domain/Dtos"
	"CSMSite.Backend/Entities"
	"CSMSite.Backend/Repositories/IRepositories"
	"github.com/gin-gonic/gin"
)

type CreateUserInteractor struct {
	Db   IRepositories.IDbRepository
	UserRepository IRepositories.IUserRepository
}

func (interactor *CreateUserInteractor) CreateUser(c *gin.Context) (user Dtos.UserResponse, err error) {
	db := interactor.Db.Connect()

	body := make([]byte, c.Request.ContentLength)
	c.Request.Body.Read(body)
	var req Dtos.UserRequest
	json.Unmarshal(body, &req)

	now := time.Now().Format("2006-01-02T15:04:05+09:00")

	newUser := Entities.User{
		UserName:    req.UserName,
		Password:    req.Password,
		Email:       req.Email,
		Admin:       false,
		DisableFlag: false,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	createdUser, err := interactor.UserRepository.Create(db, newUser)
	if err != nil {
		return Dtos.UserResponse{}, err
	}

	user = Dtos.UserResponse{
		Id:          createdUser.Id,
		UserName:    createdUser.UserName,
		Password:    createdUser.Password,
		Email:       createdUser.Email,
		Admin:       createdUser.Admin,
		DisableFlag: createdUser.DisableFlag,
		CreatedAt:   createdUser.CreatedAt,
		UpdatedAt:   createdUser.UpdatedAt,
	}
	return user, err
}
