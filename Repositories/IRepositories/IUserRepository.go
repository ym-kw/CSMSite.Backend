package IRepositories

import (
	"CSMSite.Backend/Domain/Dtos"
	"CSMSite.Backend/Entities"
	"github.com/jinzhu/gorm"
)

type IUserRepository interface {
	FindAll(db *gorm.DB) (user []Dtos.UserResponse, err error)
	FindById(db *gorm.DB, id int) (user Dtos.UserResponse, err error)
	Create(db *gorm.DB, req Entities.User) (user Entities.User, err error)
}
