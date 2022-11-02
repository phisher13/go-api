package storage

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/phisher13/go-api/internal/domain/auth/entity"
)

const (
	usersTable = "users"
)

type storage struct {
	db *sqlx.DB
}

func NewAuthorizationStorage(db *sqlx.DB) Authorization {
	return &storage{db: db}
}

func (s *storage) GetUser(username, passwordHash string) (entity.UserModel, error) {
	var user entity.UserModel
	query := fmt.Sprintf("SELECT uuid FROM %s WHERE username=$1 and password_hash=$2", usersTable)
	err := s.db.Get(&user, query, username, passwordHash)

	return user, err
}

func (s *storage) CreateUser(dto entity.UserDTO) (string, error) {
	var uuid string
	query := fmt.Sprintf("INSERT INTO %s(username, email, password_hash) VALUES ($1, $2, $3) RETURNING uuid", usersTable)
	err := s.db.QueryRow(query, dto.Username, dto.Email, dto.Password).Scan(&uuid)

	return uuid, err
}
