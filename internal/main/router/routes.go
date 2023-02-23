package router

import (
	"github.com/bvaledev/go-expert-commerce-api/configs"
	"github.com/bvaledev/go-expert-commerce-api/internal/main/factory"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
)

func SetupRoutes(config *configs.Config) *chi.Mux {
	router := chi.NewRouter()

	defaultMiddlewares(router)

	userRoutes(router, config)
	productRoutes(router, config)
	docsRoutes(router, config)

	return router
}

func defaultMiddlewares(router *chi.Mux) {
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
}

func userRoutes(router *chi.Mux, config *configs.Config) {
	userHandler := factory.MakeUserHandler()

	router.Route("/users", func(r chi.Router) {
		router.Use(middleware.WithValue("jwt", config.TokenAuth))
		router.Use(middleware.WithValue("jwtExpiresIn", config.JWTExpiresIn))
		r.Post("/", userHandler.CreateUser)
		r.Post("/login", userHandler.AuthenticateUser)
	})
}

func productRoutes(router *chi.Mux, config *configs.Config) {
	productHandler := factory.MakeProductHandler()

	router.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(config.TokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Post("/", productHandler.CreateProduct)
		r.Get("/{id}", productHandler.GetProduct)
		r.Get("/", productHandler.GetProducts)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})
}

func docsRoutes(router *chi.Mux, config *configs.Config) {
	router.Get("/docs/*", httpSwagger.Handler(
		httpSwagger.URL("./docs/doc.json"), //The url pointing to API definition
	))
}
