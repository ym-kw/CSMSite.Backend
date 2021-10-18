package ScheduleUsecases

import (
	"CSMSite.Backend/Domain/Dtos"
	"github.com/gin-gonic/gin"
)

type ICreateScheduleUsecase interface {
	CreateSchedule(c *gin.Context) (schedule Dtos.ScheduleResponse, err error)
}
