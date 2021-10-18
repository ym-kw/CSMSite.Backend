package ScheduleUsecases

import "CSMSite.Backend/Domain/Dtos"

type IReadScheduleUsecase interface {
	GetAllSchedules() (schedules []Dtos.ScheduleResponse, err error)
	GetMySchedules(userId int) (schedules []Dtos.ScheduleResponse, err error)
}
