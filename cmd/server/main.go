package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bvaledev/go-expert-commerce-api/configs"
	_ "github.com/bvaledev/go-expert-commerce-api/docs"
	_ "github.com/bvaledev/go-expert-commerce-api/internal/main/factory"
	"github.com/bvaledev/go-expert-commerce-api/internal/main/router"
	_ "github.com/swaggo/files"
)

//	@title			GO Expert E-commerce
//	@version		1.0
//	@description	Uma aplicação de ecommerce feita em go.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.brendo.dev.br/support
//	@contact.email	brendo@brendo.dev.br

// @host						localhost:8000
// @BasePath					/
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
func main() {
	config := configs.LoadConfig(".")

	appRouter := router.SetupRoutes(config)

	port := fmt.Sprintf(":%v", config.WebServerPort)
	StartMessage(port)
	log.Fatal(http.ListenAndServe(port, appRouter))
}

func StartMessage(port string) {
	fmt.Printf("\n\n")
	fmt.Printf("App Running: http://localhost%s\n", port)
	fmt.Printf("Docs: http://localhost%s/docs/index.html\n", port)
}
