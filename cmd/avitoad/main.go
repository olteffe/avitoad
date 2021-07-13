package main

import (
	"github.com/gorilla/mux"
	"github.com/olteffe/avitoad/configs"
	"github.com/olteffe/avitoad/internal/routes"
	"github.com/olteffe/avitoad/internal/utils"

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
	//_ "github.com/olteffe/avitoad/docs"   // load Swagger docs
)

// main func
func main() {
	// Initialize a new router.
	router := mux.NewRouter()

	// List of app routes:
	routes.PublicRoutes(router)
	routes.SwaggerRoutes(router)

	// Initialize server.
	server := configs.ServerConfig(router)

	// Start API server.
	utils.StartServerWithGracefulShutdown(server)
}
