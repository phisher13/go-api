package service

import (
	"github.com/phisher13/go-api/internal/domain/auth/entity"
	"github.com/phisher13/go-api/internal/domain/auth/storage"
)

type service struct {
	store storage.Authorization
}

func NewAuthService(store storage.Authorization) AuthService {
	return &service{store: store}
}

func (s *service) GetUser(username, passwordHash string) (entity.UserModel, error) {
	return s.store.GetUser(username, passwordHash)
}

func (s *service) CreateUser(dto *entity.UserDTO) (string, error) {
	return s.store.CreateUser(dto)
}
