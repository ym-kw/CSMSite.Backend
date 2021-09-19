package Usecases

import (
	"CSMSite.Backend/Domain/Dtos"
)

type ICreateUserUsecase interface {
	createUser(Dtos.UserRequest) Dtos.UserResponse
}
