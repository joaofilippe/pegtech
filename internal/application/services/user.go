package services

import (
	"github.com/joaofilippe/pegtech/internal/domain/entities"
	"github.com/joaofilippe/pegtech/internal/domain/iservices"
	irepositories "github.com/joaofilippe/pegtech/internal/domain/irepositories"
	userusecases "github.com/joaofilippe/pegtech/internal/domain/usecases/user"
)

type UserService struct {
	createUseCase  *userusecases.CreateUserCase
	getByEmailCase *userusecases.GetUserByEmailCase
	getByIDCase    *userusecases.GetUserByIDCase
	updateUseCase  *userusecases.UpdateUserCase
	deleteUseCase  *userusecases.DeleteUserCase
}

func NewUserService(repo irepositories.UserRepository) iservices.UserService {
	return &UserService{
		createUseCase:  userusecases.NewCreateUserCase(repo),
		getByEmailCase: userusecases.NewGetUserByEmailCase(repo),
		getByIDCase:    userusecases.NewGetUserByIDCase(repo),
		updateUseCase:  userusecases.NewUpdateUserCase(repo),
		deleteUseCase:  userusecases.NewDeleteUserCase(repo),
	}
}

func (u *UserService) CreateUser(username, email, password string) (*entities.User, error) {
	input := userusecases.CreateUserInput{
		Username: username,
		Email:    email,
		Password: password,
	}
	return u.createUseCase.Execute(input)
}

func (u *UserService) GetUserByEmail(email string) (*entities.User, error) {
	return u.getByEmailCase.Execute(email)
}

func (u *UserService) GetUserByID(id string) (*entities.User, error) {
	return u.getByIDCase.Execute(id)
}

func (u *UserService) UpdateUser(id string, username, email string) (*entities.User, error) {
	input := userusecases.UpdateUserInput{
		ID:       id,
		Username: username,
		Email:    email,
	}
	return u.updateUseCase.Execute(input)
}

func (u *UserService) DeleteUser(id string) error {
	return u.deleteUseCase.Execute(id)
}
