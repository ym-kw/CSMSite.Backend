package UserUsecases

import (
	"CSMSite.Backend/Domain/Dtos"
)

type IReadUserUsecase interface {
	GetUserList() []Dtos.UserResponse
	GetUser(int) Dtos.UserResponse
}
