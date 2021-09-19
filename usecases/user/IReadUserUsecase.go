package Usecases

import (
	"CSMSite.Backend/Domain/Dtos"
)

type IReadUserUsecase interface {
	getUserList() []Dtos.UserResponse
	getUser(int) Dtos.UserResponse
}
