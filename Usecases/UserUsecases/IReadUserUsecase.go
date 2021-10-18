package UserUsecases

import (
	"CSMSite.Backend/Domain/Dtos"
)

type IReadUserUsecase interface {
	GetAllUsers() (users []Dtos.UserResponse, err error)
	GetUser(id int) (user Dtos.UserResponse, err error)
}
