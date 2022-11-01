package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/phisher13/go-api/internal/domain/auth/entity"
	"github.com/phisher13/go-api/internal/domain/auth/storage"
	"os"
	"time"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserUUID string `json:"user_uuid"`
}

const (
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
)

type service struct {
	store storage.Authorization
}

func NewAuthService(store storage.Authorization) AuthService {
	return &service{store: store}
}

func (s *service) GetUser(username, passwordHash string) (entity.UserModel, error) {
	return s.store.GetUser(username, generatePasswordHash(passwordHash))
}

func (s *service) CreateUser(dto entity.UserDTO) (string, error) {
	dto.Password = generatePasswordHash(dto.Password)
	return s.store.CreateUser(dto)
}

func (s *service) GenerateToken(username, password string) (string, error) {
	user, err := s.store.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.UUID,
	})

	return token.SignedString([]byte(signingKey))
}

func generatePasswordHash(password string) string {
	salt := os.Getenv("SALT")
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
