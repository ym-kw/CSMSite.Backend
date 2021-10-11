package ScheduleUsecases

import "CSMSite.Backend/Domain/Dtos"

type IReadScheduleUsecase interface {
	getScheduleList() []Dtos.ScheduleResponse
	getSchedule(int) Dtos.ScheduleResponse
}
