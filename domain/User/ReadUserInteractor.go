package UserInteractor

import (
	"CSMSite.Backend/Domain/Dtos"
	"CSMSite.Backend/Repositories/IRepositories"
)

type ReadUserInteractor struct {
	Db   IRepositories.IDbRepository
	User IRepositories.IUserRepository
}

func (interactor *ReadUserInteractor) GetUser(id int) (user Dtos.UserResponse, err error) {
	db := interactor.Db.Connect()

	foundUser, err := interactor.User.FindById(db, id)
	if err != nil {
		return Dtos.UserResponse{}, err
	}
	user = Dtos.UserResponse(foundUser)

	return user, nil
}
