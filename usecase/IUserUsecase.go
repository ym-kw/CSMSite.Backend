package usecase

import (
	"CSMSite.Backend/domain/dtos"
)

type UserUsecase interface {
	createUser(dtos.UserRequest) dtos.UserResponse
	getUserList() []dtos.UserResponse
	getUser(int) dtos.UserResponse
}
