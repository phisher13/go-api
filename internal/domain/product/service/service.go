package service

import (
	"github.com/phisher13/go-api/internal/domain/product/entity"
	"github.com/phisher13/go-api/internal/domain/product/storage"
)

type service struct {
	store storage.ProductStorage
}

func (s *service) Update(uuid string, dto entity.ProductDTO) error {
	return s.store.Update(uuid, dto)
}

func (s *service) Delete(uuid string) error {
	return s.store.Delete(uuid)
}

func (s *service) GetByUUID(uuid string) (entity.ProductModel, error) {
	return s.store.GetByUUID(uuid)
}

func (s *service) GetAll(user_uuid string) ([]entity.ProductModel, error) {
	return s.store.GetAll(user_uuid)
}

func (s *service) Create(dto entity.ProductDTO, user_uuid string) (string, error) {
	return s.store.Create(dto, user_uuid)
}

func NewProductService(store storage.ProductStorage) ProductService {
	return &service{store: store}
}
