package Repository

import (
	"CSMSite.Backend/domain/dtos"
	"github.com/jinzhu/gorm"
)

type IScheduleRepository interface {
	FindByUserId(db *gorm.DB, id int) (shchedule dtos.ScheduleResponse, err error)
	getList() (shchedule dtos.ScheduleResponse, err error)
}
