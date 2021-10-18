package IRepositories

import (
	"CSMSite.Backend/Entities"
	"github.com/jinzhu/gorm"
)

type IScheduleRepository interface {
	FindAll(db *gorm.DB) (shchedules []Entities.Schedule, err error)
	FindByUserId(db *gorm.DB, userId int) (shchedules []Entities.Schedule, err error)
	Create(db *gorm.DB, req Entities.Schedule) (schedule Entities.Schedule, err error)
}
