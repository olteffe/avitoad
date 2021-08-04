package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/olteffe/avitoad/internal/controllers"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(router *mux.Router) {
	// Routes for GET method:
	router.HandleFunc("/api/v1/ad/{id}", controllers.GetAd).Methods(http.MethodGet) // get one ad by ID
	router.HandleFunc("/api/v1/ads", controllers.GetAds).Methods(http.MethodGet)    // Get list of all ads TODO pagination and sort

	// Routes for POST method:
	router.HandleFunc("/api/v1/ad", controllers.CreateAd).Methods(http.MethodPost) // create new ad
}
