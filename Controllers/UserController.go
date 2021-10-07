package Controllers

import (
	"strconv"

	UserInteractor "CSMSite.Backend/Domain/User"
	"CSMSite.Backend/Infrastructure/InfrastructureInterface"
	"CSMSite.Backend/Repositories"
)

type UserController struct {
	readUserInteractor   UserInteractor.ReadUserInteractor
	createUserInteractor UserInteractor.CreateUserInteractor
}

func NewUserController(db InfrastructureInterface.IDb) *UserController {
	return &UserController{
		readUserInteractor: UserInteractor.ReadUserInteractor{
			Db:   &Repositories.DbRepository{DbRepository: db},
			User: &Repositories.UserRepository{},
		},
		createUserInteractor: UserInteractor.CreateUserInteractor{
			Db:   &Repositories.DbRepository{DbRepository: db},
			User: &Repositories.UserRepository{},
		},
	}
}

func (controller *UserController) Get(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := controller.readUserInteractor.GetUser(id)
	if err != nil {
		c.JSON(501, nil)
		return
	}
	c.JSON(200, NewH("success", user))
}
