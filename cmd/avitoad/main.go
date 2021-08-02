package main

import (
	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload" // load .env file automatically

	"github.com/olteffe/avitoad/configs"
	_ "github.com/olteffe/avitoad/docs" // load Swagger docs
	"github.com/olteffe/avitoad/internal/routes"
	"github.com/olteffe/avitoad/internal/utils"
)

// @title API
// @version 1.0
// @description This is an auto-generated API Docs for the test task.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your@mail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api
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
