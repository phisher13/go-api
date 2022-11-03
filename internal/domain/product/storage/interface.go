package storage

import "github.com/phisher13/go-api/internal/domain/product/entity"

type ProductStorage interface {
	GetByUUID(uuid string) (entity.ProductModel, error)
	GetAll(user_uuid string) ([]entity.ProductModel, error)
	Create(dto entity.ProductDTO, user_uuid string) (string, error)
	Update(uuid string, dto entity.ProductDTO) error
	Delete(uuid string) error
}
