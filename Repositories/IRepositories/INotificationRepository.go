package IRepositories

import (
	"CSMSite.Backend/Entities"
	"github.com/jinzhu/gorm"
)

type INotificationRepository interface {
	FindById(db *gorm.DB, id int) (notification Entities.Notification, err error)
	FindListByUserId(db *gorm.DB, userId int) (notifications []Entities.Notification, err error)
	Create(db *gorm.DB, req Entities.Notification) (notification Entities.Notification, err error)
}
