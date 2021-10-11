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

	user = Dtos.UserResponse{
		Id:          foundUser.Id,
		UserName:    foundUser.UserName,
		Password:    foundUser.Password,
		Email:       foundUser.Email,
		Admin:       foundUser.Admin,
		DisableFlag: foundUser.DisableFlag,
		CreatedAt:   foundUser.CreatedAt,
		UpdatedAt:   foundUser.UpdatedAt,
	}

	return user, nil
}

func (interactor *ReadUserInteractor) GetAllUsers() (users []Dtos.UserResponse, err error) {
	db := interactor.Db.Connect()

	foundUsers, err := interactor.User.FindAll(db)
	if err != nil {
		return []Dtos.UserResponse{}, err
	}

	for _, v := range foundUsers {
		users = append(users, Dtos.UserResponse{
			Id:          v.Id,
			UserName:    v.UserName,
			Password:    v.Password,
			Email:       v.Email,
			Admin:       v.Admin,
			DisableFlag: v.DisableFlag,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		})
	}

	return users, nil
}
