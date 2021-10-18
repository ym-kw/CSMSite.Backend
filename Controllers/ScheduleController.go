package Controllers

import (
	"strconv"

	ScheduleInteractor "CSMSite.Backend/Domain/Schedule"
	"CSMSite.Backend/Infrastructure/InfrastructureInterface"
	"CSMSite.Backend/Repositories"
	"CSMSite.Backend/Usecases/ScheduleUsecases"
	"github.com/gin-gonic/gin"
)

type ScheduleController struct {
	readScheduleInteractor   ScheduleUsecases.IReadScheduleUsecase
	createScheduleInteractor ScheduleUsecases.ICreateScheduleUsecase
}

func NewScheduleController(db InfrastructureInterface.IDb) *ScheduleController {
	return &ScheduleController{
		readScheduleInteractor: &ScheduleInteractor.ReadScheduleInteractor{
			Db:                 &Repositories.DbRepository{DbRepository: db},
			ScheduleRepository: &Repositories.ScheduleRepository{},
		},
		createScheduleInteractor: &ScheduleInteractor.CreateScheduleInteractor{
			Db:                 &Repositories.DbRepository{DbRepository: db},
			ScheduleRepository: &Repositories.ScheduleRepository{},
		},
	}
}

func (controller *ScheduleController) Get(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Params.ByName("userId"))

	schedule, err := controller.readScheduleInteractor.GetMySchedules(userId)
	if err != nil {
		c.AbortWithStatus(404)
	}
	c.JSON(200, schedule)
}

func (controller *ScheduleController) GetAll(c *gin.Context) {
	schedule, err := controller.readScheduleInteractor.GetAllSchedules()
	if err != nil {
		c.AbortWithStatus(404)
	}
	c.JSON(200, schedule)
}

func (controller *ScheduleController) Post(c *gin.Context) {
	schedule, err := controller.createScheduleInteractor.CreateSchedule(c)
	if err != nil {
		c.AbortWithStatus(404)
	}
	c.JSON(200, schedule)
}
