package IRepositories

import (
	"CSMSite.Backend/Domain/Dtos"
	"github.com/jinzhu/gorm"
)

type IScheduleRepository interface {
	FindByUserId(db *gorm.DB, id int) (shchedule Dtos.ScheduleResponse, err error)
	getList() (shchedule Dtos.ScheduleResponse, err error)
}
