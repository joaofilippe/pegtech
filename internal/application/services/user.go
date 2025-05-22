package services

import (
	"github.com/joaofilippe/pegtech/internal/domain/entities"
	irepositories "github.com/joaofilippe/pegtech/internal/domain/irepositories"
	"github.com/joaofilippe/pegtech/internal/domain/iservices"
	userusecases "github.com/joaofilippe/pegtech/internal/domain/usecases/user"
)

type UserService struct {
	createUseCase  *userusecases.CreateUserCase
	getByEmailCase *userusecases.GetUserByEmailCase
	getByIDCase    *userusecases.GetUserByIDCase
	updateUseCase  *userusecases.UpdateUserCase
	deleteUseCase  *userusecases.DeleteUserCase
	loginUseCase   *userusecases.LoginUserCase
}

func NewUserService(repo irepositories.UserRepository,) iservices.UserService {
	return &UserService{
		createUseCase:  userusecases.NewCreateUserCase(repo),
		getByEmailCase: userusecases.NewGetUserByEmailCase(repo),
		getByIDCase:    userusecases.NewGetUserByIDCase(repo),
		updateUseCase:  userusecases.NewUpdateUserCase(repo),
		deleteUseCase:  userusecases.NewDeleteUserCase(repo),
		loginUseCase:   userusecases.NewLoginUserCase(repo),
	}
}

func (u *UserService) CreateUser(username, name, email, password, phone string, userType entities.UserType) (*entities.User, error) {
	input := userusecases.CreateUserInput{
		Username: username,
		Name:     name,
		Email:    email,
		Password: password,
		Phone:    phone,
		Type:     userType,
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

// Login authenticates a user and returns a JWT token
func (u *UserService) Login(email, password string) (*userusecases.LoginResponse, error) {
	input := userusecases.LoginUserInput{
		Email:    email,
		Password: password,
	}
	return u.loginUseCase.Execute(input)
}
