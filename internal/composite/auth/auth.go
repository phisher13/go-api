package auth

import (
	"github.com/jmoiron/sqlx"
	handle "github.com/phisher13/go-api/internal/adapter/api/auth"
	"github.com/phisher13/go-api/internal/domain/auth/service"
	"github.com/phisher13/go-api/internal/domain/auth/storage"
)

type AuthorizationComposite struct {
	Storage storage.Authorization
	Service service.AuthService
	Handler handle.AuthorizationHandler
}

func NewAuthorizationComposite(db *sqlx.DB) (*AuthorizationComposite, error) {
	store := storage.NewAuthorizationStorage(db)
	srv := service.NewAuthService(store)
	handler := handle.NewAuthorizationHandler(srv)

	return &AuthorizationComposite{
		Storage: store,
		Service: srv,
		Handler: handler,
	}, nil
}
