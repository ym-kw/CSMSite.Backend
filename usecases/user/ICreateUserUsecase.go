package usecase

import (
	"CSMSite.Backend/domain/dtos"
)

type ICreateUserUsecase interface {
	createUser(dtos.UserRequest) dtos.UserResponse
}
