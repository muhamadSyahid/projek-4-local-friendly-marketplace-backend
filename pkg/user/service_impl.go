package user

import (
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"pade-backend/pkg/entities"
)

type serviceImpl struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &serviceImpl{
		repo: repository,
	}
}

func (s *serviceImpl) Register(user *entities.User) (*entities.User, error) {
	if user == nil {
		return nil, errors.New("user is required")
	}
	if strings.TrimSpace(user.Email) == "" {
		return nil, errors.New("email is required")
	}
	if strings.TrimSpace(user.Password) == "" {
		return nil, errors.New("password is required")
	}

	existing, err := s.repo.GetByEmail(user.Email)
	if err == nil && existing != nil {
		return nil, errors.New("email already registered")
	}
	if err != nil && err.Error() != "user not found" {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(hashedPassword)
	if len(user.Roles) == 0 {
		user.Roles = []entities.Roles{entities.RoleBuyer}
	}

	createdUser, err := s.repo.Create(user)
	if err != nil {
		return nil, err
	}

	return sanitizeUser(createdUser), nil
}

func (s *serviceImpl) Login(email, password string) (*entities.User, error) {
	if strings.TrimSpace(email) == "" {
		return nil, errors.New("email is required")
	}
	if strings.TrimSpace(password) == "" {
		return nil, errors.New("password is required")
	}

	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	return sanitizeUser(user), nil
}

func (s *serviceImpl) GetUserByID(id string) (*entities.User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return sanitizeUser(user), nil
}

func (s *serviceImpl) UpdateUser(user *entities.User) (*entities.User, error) {
	if user != nil && strings.TrimSpace(user.Password) != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		user.Password = string(hashedPassword)
	}

	updatedUser, err := s.repo.Update(user)
	if err != nil {
		return nil, err
	}

	return sanitizeUser(updatedUser), nil
}

func (s *serviceImpl) DeleteUser(id string) error {
	return s.repo.Delete(id)
}

func (s *serviceImpl) GetAllUsers() ([]*entities.User, error) {
	users, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		sanitizeUser(user)
	}

	return users, nil
}

func sanitizeUser(user *entities.User) *entities.User {
	if user == nil {
		return nil
	}

	sanitized := *user
	sanitized.Password = ""
	return &sanitized
}
