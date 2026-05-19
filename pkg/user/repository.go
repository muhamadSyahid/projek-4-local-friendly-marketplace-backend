package user

import "pade-backend/pkg/entities"

type Repository interface {
	Create(user *entities.User) (*entities.User, error)
	GetByID(id string) (*entities.User, error)
	GetByEmail(email string) (*entities.User, error)
	Update(user *entities.User) (*entities.User, error)
	Delete(id string) error
	GetAll() ([]*entities.User, error)
}
