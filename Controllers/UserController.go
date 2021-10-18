package Controllers

import (
	"strconv"

	UserInteractor "CSMSite.Backend/Domain/User"
	"CSMSite.Backend/Infrastructure/InfrastructureInterface"
	"CSMSite.Backend/Repositories"
	"CSMSite.Backend/Usecases/UserUsecases"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	readUserInteractor   UserUsecases.IReadUserUsecase
	createUserInteractor UserUsecases.ICreateUserUsecase
}

func NewUserController(db InfrastructureInterface.IDb) *UserController {
	return &UserController{
		readUserInteractor: &UserInteractor.ReadUserInteractor{
			Db:             &Repositories.DbRepository{DbRepository: db},
			UserRepository: &Repositories.UserRepository{},
		},
		createUserInteractor: &UserInteractor.CreateUserInteractor{
			Db:             &Repositories.DbRepository{DbRepository: db},
			UserRepository: &Repositories.UserRepository{},
		},
	}
}

func (controller *UserController) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	user, err := controller.readUserInteractor.GetUser(id)
	if err != nil {
		c.AbortWithStatus(404)
		return
	}
	c.JSON(200, user)
}

func (controller *UserController) GetAll(c *gin.Context) {
	users, err := controller.readUserInteractor.GetAllUsers()
	if err != nil {
		c.AbortWithStatus(404)
	}
	c.JSON(200, users)
}

func (controller *UserController) Post(c *gin.Context) {
	user, err := controller.createUserInteractor.CreateUser(c)
	if err != nil {
		c.AbortWithStatus(404)
	}
	c.JSON(200, user)
}
