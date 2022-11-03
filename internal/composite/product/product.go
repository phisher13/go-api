package product

import (
	"github.com/jmoiron/sqlx"
	handle "github.com/phisher13/go-api/internal/adapter/api/product"
	"github.com/phisher13/go-api/internal/domain/product/service"
	"github.com/phisher13/go-api/internal/domain/product/storage"
)

type ProductComposite struct {
	Storage storage.ProductStorage
	Service service.ProductService
	Handler handle.ProductHandler
}

func NewProductComposite(db *sqlx.DB) (*ProductComposite, error) {
	store := storage.NewProductStorage(db)
	srv := service.NewProductService(store)
	handler := handle.NewProductHandler(srv)
	return &ProductComposite{
		Storage: store,
		Service: srv,
		Handler: handler,
	}, nil
}
