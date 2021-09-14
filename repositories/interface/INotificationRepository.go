package repository

import (
	"CSMSite.Backend/domain/dtos"
	"github.com/jinzhu/gorm"
)

type INotificationRepository interface {
	FindById(db *gorm.DB, id int) (notification dtos.NotificationResponse, err error)
	FindListById(db *gorm.DB, id int) (notification dtos.NotificationResponse, err error)
}
