package usecase

import (
	"CSMSite.Backend/domain/dtos"
)

type IReadUserUsecase interface {
	getUserList() []dtos.UserResponse
	getUser(int) dtos.UserResponse
}
