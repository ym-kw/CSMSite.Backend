package Repositories

import (
	"CSMSite.Backend/Entities"
	"github.com/jinzhu/gorm"
)

type ScheduleRepository struct{}

func (repo ScheduleRepository) FindAll(db *gorm.DB) (schedules []Entities.Schedule, err error) {
	result := db.Find(&schedules)

	return schedules, result.Error
}

func (repo ScheduleRepository) FindByUserId(db *gorm.DB, userId int) (schedules []Entities.Schedule, err error) {
	schedules = []Entities.Schedule{}
	result := db.Where("userId = ?", userId).Find(&schedules)

	return schedules, result.Error
}

func (repo *ScheduleRepository) Create(db *gorm.DB, req Entities.Schedule) (schedule Entities.Schedule, err error) {
	result := db.Create(&req)

	return req, result.Error
}
