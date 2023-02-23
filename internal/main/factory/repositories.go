package factory

import (
	"github.com/bvaledev/go-expert-commerce-api/internal/domain/contracts"
	"github.com/bvaledev/go-expert-commerce-api/internal/infra/database"
)

func MakeUserRepository() contracts.IUserRepository {
	return database.NewUserDB(DatabaseConection)
}

func MakeProductRepository() contracts.IProductRepository {
	return database.NewProductDB(DatabaseConection)
}
