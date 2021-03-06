package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/olteffe/avitoad/internal/utils"

	httpSwagger "github.com/swaggo/http-swagger"
)

// SwaggerRoutes func for describe group of Swagger routes.
func SwaggerRoutes(router *mux.Router) {
	// Define server settings:
	serverConnURL, _ := utils.ConnectionURLBuilder("server")

	// Build Swagger route.
	getSwagger := httpSwagger.Handler(
		httpSwagger.URL("http://"+serverConnURL+"/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	)

	// Routes for GET method:
	router.PathPrefix("/swagger/").Handler(getSwagger).Methods(http.MethodGet) // get one ad by ID
}
