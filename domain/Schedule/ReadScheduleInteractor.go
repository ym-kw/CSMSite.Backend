package ScheduleInteractor

import (
	"CSMSite.Backend/Domain/Dtos"
	"CSMSite.Backend/Repositories/IRepositories"
)

type ReadScheduleInteractor struct {
	Db                 IRepositories.IDbRepository
	ScheduleRepository IRepositories.IScheduleRepository
}

func (interactor *ReadScheduleInteractor) GetAllSchedules() (schedules []Dtos.ScheduleResponse, err error) {
	db := interactor.Db.Connect()

	foundSchedules, err := interactor.ScheduleRepository.FindAll(db)
	if err != nil {
		return []Dtos.ScheduleResponse{}, err
	}

	for _, v := range foundSchedules {
		schedules = append(schedules, Dtos.ScheduleResponse{
			Id:            v.Id,
			StartDateTime: v.StartDateTime,
			EndDateTime:   v.EndDateTime,
			UserId:        v.UserId,
			DisableFlag:   v.DisableFlag,
			CreatedAt:     v.CreatedAt,
			UpdatedAt:     v.UpdatedAt,
		})
	}
	return schedules, nil
}

func (interactor *ReadScheduleInteractor) GetMySchedules(userId int) (schedules []Dtos.ScheduleResponse, err error) {
	db := interactor.Db.Connect()

	foundSchedules, err := interactor.ScheduleRepository.FindByUserId(db, userId)
	if err != nil {
		return []Dtos.ScheduleResponse{}, err
	}

	for _, v := range foundSchedules {
		schedules = append(schedules, Dtos.ScheduleResponse{
			Id:            v.Id,
			StartDateTime: v.StartDateTime,
			EndDateTime:   v.EndDateTime,
			UserId:        v.UserId,
			DisableFlag:   v.DisableFlag,
			CreatedAt:     v.CreatedAt,
			UpdatedAt:     v.UpdatedAt,
		})
	}

	return schedules, nil
}
