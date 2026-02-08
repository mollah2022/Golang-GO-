package product

import "ecomerce/domain"

type Service interface {
	Creat(domain.Products) (*domain.Products, error)
	Get(id int) (*domain.Products, error)
	List(page, limit int64) ([]*domain.Products, error)
	Count () (int64,error)
	Update(domain.Products) (*domain.Products, error)
	Delete(id int) error
}