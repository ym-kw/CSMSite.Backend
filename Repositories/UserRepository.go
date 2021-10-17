package Repositories

import (
	"errors"

	"CSMSite.Backend/Entities"
	"github.com/jinzhu/gorm"
)

type UserRepository struct{}

func (repo *UserRepository) FindAll(db *gorm.DB) (users []Entities.User, err error) {
	result := db.Find(&users)

	return users, result.Error
}

func (repo *UserRepository) FindById(db *gorm.DB, id int) (user Entities.User, err error) {
	user = Entities.User{}
	db.First(&user, id)
	if user.Id <= 0 {
		return Entities.User{}, errors.New("user is not found")
	}
	return user, nil
}

func (repo *UserRepository) Create(db *gorm.DB, req Entities.User) (user Entities.User, err error) {
	result := db.Create(&req)

	return req, result.Error
}
