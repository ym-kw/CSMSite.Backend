package ScheduleInteractor

import (
	"encoding/json"
	"time"

	"CSMSite.Backend/Domain/Dtos"
	"CSMSite.Backend/Entities"
	"CSMSite.Backend/Repositories/IRepositories"
	"github.com/gin-gonic/gin"
)

type CreateScheduleInteractor struct {
	Db                 IRepositories.IDbRepository
	ScheduleRepository IRepositories.IScheduleRepository
}

func (interactor *CreateScheduleInteractor) CreateSchedule(c *gin.Context) (schedule Dtos.ScheduleResponse, err error) {
	db := interactor.Db.Connect()

	body := make([]byte, c.Request.ContentLength)
	c.Request.Body.Read(body)
	var req Dtos.ScheduleRequest
	json.Unmarshal(body, &req)

	now := time.Now().Format("2006-01-02T15:04:05+09:00")

	newSchedule := Entities.Schedule{
		StartDateTime: req.StartDateTime,
		EndDateTime:   req.EndDateTime,
		UserId:        req.UserId,
		DisableFlag:   false,
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	createdSchedule, err := interactor.ScheduleRepository.Create(db, newSchedule)
	if err != nil {
		return Dtos.ScheduleResponse{}, err
	}

	schedule = Dtos.ScheduleResponse{
		Id:            createdSchedule.Id,
		StartDateTime: createdSchedule.StartDateTime,
		EndDateTime:   createdSchedule.EndDateTime,
		UserId:        createdSchedule.UserId,
		DisableFlag:   createdSchedule.DisableFlag,
		CreatedAt:     createdSchedule.CreatedAt,
		UpdatedAt:     createdSchedule.UpdatedAt,
	}
	return schedule, err
}
