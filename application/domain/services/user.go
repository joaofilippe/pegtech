package iservices

import (
	irepositories "github.com/joaofilippe/pegtech/application/domain/repositories"
	userusecases "github.com/joaofilippe/pegtech/application/domain/usecases/user"
)

type IUserService interface {
	CreateUser(email string, password string) error
	//UpdateUser(email string, password string) error
	//DeleteUser(email string) error
	//GetUser(email string) (string, error)
}

type UserService struct {
	createUseCase *userusecases.CreateUserCase
}

func NewUserService(repo irepositories.IUserRepository) IUserService {
	createUseCase := userusecases.NewCreateUserCase(repo)

	return &UserService{createUseCase: createUseCase}
}

func (u *UserService) CreateUser(email string, password string) error {
	return u.createUseCase.Execute(email, password)
}
