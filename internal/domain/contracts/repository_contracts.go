package contracts

import "github.com/bvaledev/go-expert-commerce-api/internal/domain/entity"

type IUserRepository interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
	FindById(id string) (*entity.User, error)
}

type IProductRepository interface {
	Create(product *entity.Product) error
	FindAll(page, limit int, sort string) ([]entity.Product, error)
	FindByID(id string) (*entity.Product, error)
	Update(product *entity.Product) error
	Delete(id string) error
}
