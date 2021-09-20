package Repositories

import (
	"CSMSite.Backend/Domain/Dtos"
	"github.com/jinzhu/gorm"
)

type IUserRepository interface {
	FindById(db *gorm.DB, id int) (user Dtos.UserResponse, err error)
}
