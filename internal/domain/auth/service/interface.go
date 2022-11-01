package service

import "github.com/phisher13/go-api/internal/domain/auth/entity"

type AuthService interface {
	GetUser(username, passwordHash string) (entity.UserModel, error)
	CreateUser(dto *entity.UserDTO) (string, error)
}
