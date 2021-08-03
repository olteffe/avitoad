package controllers

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/olteffe/avitoad/internal/utils"
	"github.com/olteffe/avitoad/internal/validators"
	"net/http"
	"time"

	"github.com/olteffe/avitoad/internal/database/pg"
	"github.com/olteffe/avitoad/internal/models"
)

// GetAds func gets all exists ads.
// @Description Get all exists ads.
// @Summary get all exists ads
// @Tags Ads
// @Accept json
// @Produce json
// @Success 200 {array} models.Ads
// @Failure 404 {string} string "error"
// @Failure 500 {string} string "error"
// @Router /v1/ads [get]
func GetAds(w http.ResponseWriter, r *http.Request) {
	// Define content type.
	w.Header().Set("Content-Type", "application/json")

	// Create database connection.
	db, err := pg.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Get all ads.
	ads, err := db.GetAds()
	if err != nil {
		// Return status 404 and not found message.
		w.WriteHeader(http.StatusNotFound)
	} else {
		payload, _ := json.Marshal(ads)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(payload))
	}
}

// GetAd func gets one ad by given ID or 404 error.
// @Description Get ad by given ID.
// @Summary get ad by given ID
// @Tags Ad
// @Accept json
// @Produce json
// @Param id query UUID true "Ad ID"
// @Success 200 {object} models.Ads
// @Failure 400 {string} string "error"
// @Failure 404 {string} string "error"
// @Failure 500 {string} string "error"
// @Router /v1/ad [get]
func GetAd(w http.ResponseWriter, r *http.Request) {
	// Define content type and CORS.
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Catch ad ID from URL.
	id, err := uuid.Parse(r.URL.Query().Get("id"))
	if err != nil {
		// Return status 400.
		w.WriteHeader(http.StatusBadRequest)
	}

	// Create database connection.
	db, err := pg.OpenDBConnection()
	if err != nil {
		// Return status 500.
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Get ad by ID.
	ad, err := db.GetAd(id)
	if err != nil {
		// Return status 404.
		w.WriteHeader(http.StatusNotFound)
	} else {
		payload, _ := json.Marshal(ad)
		// Return status 200.
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(payload))
	}
}

// CreateAd func for creates a new ad.
// @Description Create a new ad.
// @Summary create a new ad
// @Tags Ad
// @Accept json
// @Produce json
// @Success 201 {object} string "ID"
// @Failure 400 {string} string "error"
// @Failure 500 {string} string "error"
// @Router /v1/ad [post]
func CreateAd(w http.ResponseWriter, r *http.Request) {
	// Define content type and CORS.
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Create a new ad struct.
	ad := &models.Ads{}

	// Checking received data from JSON body.
	if err := r.ParseForm(); err != nil {
		// Return status 400 error.
		w.WriteHeader(http.StatusBadRequest)
	}
	if err := json.NewDecoder(r.Body).Decode(&ad); err != nil {
		// Return status 400 error.
		w.WriteHeader(http.StatusBadRequest)
	}
	// Validate ad fields.
	validate := validators.AdValidator()
	if err := validate.Struct(ad); err != nil {
		// Return status 500 and database connection error.
		payload, _ := json.Marshal(map[string]interface{}{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(payload))
	}

	// Create database connection.
	db, err := pg.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		payload, _ := json.Marshal(map[string]interface{}{
			"error": true,
			"msg":   err.Error(),
		})
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(payload))
	}

	// Set initialized default data for ad:
	ad.ID = uuid.New()
	ad.CreatedAt = time.Now()

	// Create a new Ad with validated data.
	if err := db.CreateAd(ad); err != nil {
		// Return status 500 and database connection error.
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		payload, _ := json.Marshal(map[string]interface{}{
			"id": ad.ID,
		})
		// Return status 201 and ID.
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(payload))
	}
}
