package storage

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/phisher13/go-api/internal/domain/product/entity"
)

const (
	productsTable = "product"
)

type storage struct {
	db *sqlx.DB
}

func NewProductStorage(db *sqlx.DB) ProductStorage {
	return &storage{db: db}
}

func (s *storage) GetByUUID(uuid string) (entity.ProductModel, error) {
	var product entity.ProductModel
	query := fmt.Sprintf("SELECT uuid, title, description, price FROM %s WHERE product.uuid=$1 LIMIT 1", productsTable)
	err := s.db.Get(&product, query, uuid)

	return product, err
}

func (s *storage) GetAll(user_uuid string) ([]entity.ProductModel, error) {
	var products []entity.ProductModel
	query := fmt.Sprintf("SELECT uuid, title, description, price FROM %s WHERE product.user_uuid=$1", productsTable)
	err := s.db.Select(&products, query, user_uuid)

	return products, err
}

func (s *storage) Create(dto entity.ProductDTO, user_uuid string) (string, error) {
	var uuid string
	query := fmt.Sprintf("INSERT INTO %s(title, description, price, user_uuid) VALUES ($1, $2, $3, $4) RETURNING uuid", productsTable)
	err := s.db.QueryRow(query, dto.Title, dto.Description, dto.Price, user_uuid).Scan(&uuid)

	return uuid, err
}

func (s *storage) Update(uuid string, dto entity.ProductDTO) error {
	query := fmt.Sprintf("UPDATE %s SET title=$1, description=$2, price=$3 WHERE uuid=$4", productsTable)
	_, err := s.db.Exec(query, dto.Title, dto.Description, dto.Price, uuid)
	return err
}

func (s *storage) Delete(uuid string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE uuid=$1", productsTable)
	_, err := s.db.Exec(query, uuid)
	return err
}
