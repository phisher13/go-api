package service

import (
	"crypto/sha1"
	"errors"
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

type Service struct {
	store storage.Authorization
}

func NewAuthService(store storage.Authorization) AuthService {
	return &Service{store: store}
}

func (s *Service) GetUser(username, passwordHash string) (entity.UserModel, error) {
	return s.store.GetUser(username, generatePasswordHash(passwordHash))
}

func (s *Service) CreateUser(dto entity.UserDTO) (string, error) {
	dto.Password = generatePasswordHash(dto.Password)
	return s.store.CreateUser(dto)
}

func (s *Service) GenerateToken(username, password string) (string, error) {
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

func (s *Service) ParseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserUUID, nil
}

func generatePasswordHash(password string) string {
	salt := os.Getenv("SALT")
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
