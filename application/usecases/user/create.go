package userusecases

import irepositories "github.com/joaofilippe/pegtech/application/repositories"

type CreateUserCase struct {
	repo irepositories.IUserRepository
}

func NewCreateUserCase(repo irepositories.IUserRepository) *CreateUserCase {
	return &CreateUserCase{
		repo: repo,
	}
}

func (u *CreateUserCase) Execute(email string, password string) error {
	return u.repo.Save(email, password)
}
