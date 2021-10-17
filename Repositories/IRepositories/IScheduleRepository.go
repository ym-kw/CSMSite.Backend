package IRepositories

import (
	"CSMSite.Backend/Domain/Dtos"
	"CSMSite.Backend/Entities"
	"github.com/jinzhu/gorm"
)

type IScheduleRepository interface {
	FindAll(db *gorm.DB) (shchedule Dtos.ScheduleResponse, err error)
	FindByUserId(db *gorm.DB, userId int) (shchedule Dtos.ScheduleResponse, err error)
	Create(db *gorm.DB, req Entities.Schedule) (schedule Entities.Schedule, err error)
}
