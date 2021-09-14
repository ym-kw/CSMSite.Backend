package usecase

import "CSMSite.Backend/domain/dtos"

type ICreateScheduleUsecase interface {
	createSchedule(dtos.ScheduleRequest) dtos.ScheduleResponse
}
