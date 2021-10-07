package Repositories

import (
	"errors"

	"CSMSite.Backend/Domain/Dtos"
	"CSMSite.Backend/Entities"
	"github.com/jinzhu/gorm"
)

type UserRepository struct{}

func (repo *UserRepository) FindAll(db *gorm.DB) (user []Dtos.UserResponse, err error) {
	result := db.Find(&user)

	return user, result.Error
}

func (repo *UserRepository) FindById(db *gorm.DB, id int) (user Dtos.UserResponse, err error) {
	user = Dtos.UserResponse{}
	db.First(&user, id)
	if user.Id <= 0 {
		return Dtos.UserResponse{}, errors.New("user is not found")
	}
	return user, nil
}

func (repo *UserRepository) Create(db *gorm.DB, req Entities.User) (user Entities.User, err error) {
	result := db.Create(&req)

	return req, result.Error
}
