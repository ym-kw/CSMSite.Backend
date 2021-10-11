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
	User IRepositories.IUserRepository
}

func (interactor *CreateUserInteractor) CreateUser(c *gin.Context) (user Dtos.UserResponse, err error) {
	db := interactor.Db.Connect()

	body := make([]byte, c.Request.ContentLength)
	c.Request.Body.Read(body)
	var req Dtos.UserRequest
	json.Unmarshal(body, &req)

	newUser := Entities.User{
		UserName:    req.UserName,
		Password:    req.Password,
		Email:       req.Email,
		Admin:       false,
		DisableFlag: false,
		CreatedAt:   time.Now().Format("2006-01-02T15:04:05+09:00"),
		UpdatedAt:   time.Now().Format("2006-01-02T15:04:05+09:00"),
	}

	createUser, err := interactor.User.Create(db, newUser)
	if err != nil {
		return Dtos.UserResponse{}, err
	}
	user = Dtos.UserResponse{
		Id:          createUser.Id,
		UserName:    createUser.UserName,
		Password:    createUser.Password,
		Email:       createUser.Email,
		Admin:       createUser.Admin,
		DisableFlag: createUser.DisableFlag,
		CreatedAt:   createUser.CreatedAt,
		UpdatedAt:   createUser.UpdatedAt,
	}
	return user, err
}
