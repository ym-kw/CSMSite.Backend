package Repository

import (
	"CSMSite.Backend/Domain/Dtos"
	"github.com/jinzhu/gorm"
)

type INotificationRepository interface {
	FindById(db *gorm.DB, id int) (notification Dtos.NotificationResponse, err error)
	FindListById(db *gorm.DB, id int) (notification Dtos.NotificationResponse, err error)
}
