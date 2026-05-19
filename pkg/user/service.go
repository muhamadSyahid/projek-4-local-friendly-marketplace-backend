package user

import "pade-backend/pkg/entities"

type Service interface {
	Register(user *entities.User) (*entities.User, error)
	Login(email, password string) (*entities.User, error)
	GetUserByID(id string) (*entities.User, error)
	UpdateUser(user *entities.User) (*entities.User, error)
	DeleteUser(id string) error
	GetAllUsers() ([]*entities.User, error)
}
