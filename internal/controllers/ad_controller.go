package controllers

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/olteffe/avitoad/internal/database/pg"
	"github.com/olteffe/avitoad/internal/models"
	"github.com/olteffe/avitoad/internal/utils"
	"github.com/olteffe/avitoad/internal/validators"
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
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	fetch := models.FetchParam{}

	// Get query param: limit, cursor, sort, asc if exist
	strLimit := r.FormValue("limit")
	limit, err := strconv.ParseUint(strLimit, 10, 64)
	if err != nil {
		fetch.Limit = 10
	}
	cursor := r.FormValue("cursor")
	sort := r.FormValue("sort")
	asc := r.FormValue("asc")

	// Set default and write query values
	if limit == 0 && cursor == "" && sort == "" && asc == "" {
		fetch.Limit = 10
		fetch.Cursor = utils.EncodeCursor(time.Now(), "")
		fetch.Sort = "date"
		fetch.Asc = "DESC"
	} else {
		fetch.Limit = limit
		fetch.Cursor = cursor
		fetch.Sort = sort
		fetch.Asc = strings.ToUpper(asc)
	}

	// Structure field validation
	validate := validator.New()
	if err := validate.Struct(fetch); err != nil {
		// Return status 400 and  error.
		payload, _ := json.Marshal(map[string]interface{}{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(payload))
		return
	}

	// Create database connection.
	db, err := pg.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Get all ads.
	ads, nextCursor, err := db.GetAds(fetch)
	if err != nil {
		// Return status 404 and not found message.
		w.WriteHeader(http.StatusNotFound)
		return
	}

	payload, _ := json.Marshal(ads)
	w.Header().Set("X-NextCursor", nextCursor)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(payload))

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
// @Router /v1/ad/{id} [get]
func GetAd(w http.ResponseWriter, r *http.Request) {
	// Define content type and CORS.
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Get Query from URL if existed.
	fields, err := strconv.ParseBool(r.URL.Query().Get("fields"))
	if err != nil {
		fields = false
	}

	// Catch ad ID from URL.
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		// Return status 400.
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Create database connection.
	db, err := pg.OpenDBConnection()
	if err != nil {
		// Return status 500.
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Get ad by ID.
	ad, err := db.GetAd(id, fields)
	if err != nil {
		// Return status 404.
		w.WriteHeader(http.StatusNotFound)
		return
	}
	payload, _ := json.Marshal(ad)
	// Return status 200.
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(payload))
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
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Create a new ad struct.
	ad := &models.Ads{}

	// Checking received data from JSON body.
	if err := r.ParseForm(); err != nil {
		// Return status 400 error.
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&ad); err != nil {
		// Return status 400 error.
		w.WriteHeader(http.StatusBadRequest)
		return
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
		return
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
		return
	}

	// Set initialized default data for ad:
	ad.ID = uuid.New()
	ad.CreatedAt = time.Now()

	// Create a new Ad with validated data.
	if err := db.CreateAd(ad); err != nil {
		// Return status 500 and database connection error.
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	payload, _ := json.Marshal(map[string]interface{}{
		"id": ad.ID,
	})
	// Return status 201 and ID.
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte(payload))

}
