package usecase

import "CSMSite.Backend/domain/dtos"

type ScheduleUsecase interface {
	createSchedule(dtos.ScheduleRequest) dtos.ScheduleResponse
	getScheduleList() []dtos.ScheduleResponse
	getSchedule(int) dtos.ScheduleResponse
}
