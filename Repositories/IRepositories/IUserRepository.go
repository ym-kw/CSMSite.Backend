package IRepositories

import (
	"CSMSite.Backend/Entities"
	"github.com/jinzhu/gorm"
)

type IUserRepository interface {
	FindAll(db *gorm.DB) (user []Entities.User, err error)
	FindById(db *gorm.DB, id int) (user Entities.User, err error)
	Create(db *gorm.DB, req Entities.User) (user Entities.User, err error)
}
