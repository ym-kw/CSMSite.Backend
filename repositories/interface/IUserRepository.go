package repository

import (
	"CSMSite.Backend/domain/dtos"
	"github.com/jinzhu/gorm"
)

type IUserRepository interface {
	FindById(db *gorm.DB, id int) (user dtos.UserResponse, err error)
}
