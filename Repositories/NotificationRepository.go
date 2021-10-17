package Repositories

import (
	"errors"

	"CSMSite.Backend/Entities"
	"github.com/jinzhu/gorm"
)

type NotificationRepository struct{}

func (repo *NotificationRepository) FindById(db *gorm.DB, id int) (notification Entities.Notification, err error) {
	notification = Entities.Notification{}
	db.First(&notification, id)
	if notification.Id <= 0 {
		return Entities.Notification{}, errors.New("notification is not found")
	}
	return notification, nil
}

func (repo *NotificationRepository) FindListByUserId(db *gorm.DB, userId int) (notifications []Entities.Notification, err error) {
	notifications = []Entities.Notification{}
	result := db.Where("id = ? AND id = NULL", userId).Find(&notifications)
	return notifications, result.Error
}

func (repo *NotificationRepository) Create(db *gorm.DB, req Entities.Notification) (notification Entities.Notification, err error) {
	result := db.Create(&req)

	return req, result.Error
}
