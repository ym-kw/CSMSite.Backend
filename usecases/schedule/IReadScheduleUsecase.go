package usecase

import "CSMSite.Backend/domain/dtos"

type IReadScheduleUsecase interface {
	getScheduleList() []dtos.ScheduleResponse
	getSchedule(int) dtos.ScheduleResponse
}
