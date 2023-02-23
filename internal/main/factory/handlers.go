package factory

import (
	"github.com/bvaledev/go-expert-commerce-api/internal/application/handlers"
)

func MakeUserHandler() *handlers.UserHandlers {
	return handlers.NewUserHandler(MakeUserRepository())
}

func MakeProductHandler() *handlers.ProductHandler {
	return handlers.NewProductHandler(MakeProductRepository())
}
