package Usecase

import "CSMSite.Backend/Domain/Dtos"

type ICreateScheduleUsecase interface {
	createSchedule(Dtos.ScheduleRequest) Dtos.ScheduleResponse
}
