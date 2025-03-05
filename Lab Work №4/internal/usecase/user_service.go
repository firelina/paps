package usecase

import (
	"context"
	"errors"
	models "marine/internal/models"
)

var ErrUserNotFound = errors.New("user not found")

type UserService interface {
	RegisterUser(ctx context.Context, user *models.User) (int, error)
	GetUser(ctx context.Context, id int) (*models.User, error)
}

type userService struct {
	users map[int]*models.User
}

func NewUserService() UserService {
	return &userService{
		users: make(map[int]*models.User),
	}
}

func (s *userService) RegisterUser(ctx context.Context, user *models.User) (int, error) {
	for _, u := range s.users {
		if u.Username == user.Username {
			return 0, errors.New("user is already exist")
		}
	}

	user.ID = len(s.users) + 1
	s.users[user.ID] = user
	return user.ID, nil
}

func (s *userService) GetUser(ctx context.Context, id int) (*models.User, error) {
	user, exists := s.users[id]
	if !exists {
		return nil, ErrUserNotFound
	}
	return user, nil
}
