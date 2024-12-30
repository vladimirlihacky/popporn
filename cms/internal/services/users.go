package services

import (
	"cms/internal/models"
	"cms/internal/repositories"
)

type UserWithCredentials struct {
	*models.User
	Password string
}

type Auth interface {
	Register(*UserWithCredentials) error
	Login(*UserWithCredentials) error
}

type Users struct {
	auth Auth
	repo *repositories.Users
}

func (users *Users) Create(data *UserWithCredentials) error {

	if err := users.auth.Register(data); err != nil {
		return err
	}

	if err := users.repo.Create(data.User); err != nil {
		return err
	}

	return nil
}

func NewUsersService(auth Auth, repo *repositories.Users) *Users {
	service := &Users{
		auth: auth,
		repo: repo,
	}

	return service
}
