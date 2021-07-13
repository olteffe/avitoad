package controllers

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/olteffe/avitoad/internal/database/pg"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/olteffe/avitoad/internal/models"
	"github.com/olteffe/avitoad/internal/utils"
	"github.com/olteffe/avitoad/internal/validators"
)

// GetUsers func gets all exists users.
// @Description Get all exists users.
// @Summary get all exists users
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {array} models.User
// @Router /v1/users [get]
func GetUsers(w http.ResponseWriter, r *http.Request) {
	// Define content type.
	w.Header().Set("Content-Type", "application/json")

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

	// Get all users.
	users, err := db.GetUsers()
	if err != nil {
		// Return status 404 and not found message.
		payload, _ := json.Marshal(map[string]interface{}{
			"error": true,
			"msg":   "users were not found",
			"count": 0,
			"users": nil,
		})
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(payload))
	}

	payload, _ := json.Marshal(map[string]interface{}{
		"error": false,
		"msg":   nil,
		"count": len(users),
		"users": users,
	})
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(payload))
}

// GetUser func gets one user by given ID or 404 error.
// @Description Get user by given ID.
// @Summary get user by given ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Router /v1/user/{id} [get]
func GetUser(w http.ResponseWriter, r *http.Request) {
	// Define content type and CORS.
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Catch user ID from URL.
	id, err := uuid.Parse(r.URL.Query().Get("id"))
	if err != nil {
		// Return status 500 and database connection error.
		payload, _ := json.Marshal(map[string]interface{}{
			"error": true,
			"msg":   err.Error(),
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

	// Get user by ID.
	user, err := db.GetUser(id)
	if err != nil {
		// Return status 404 and not found message.
		payload, _ := json.Marshal(map[string]interface{}{
			"error": true,
			"msg":   "user with the given ID is not found",
			"user":  nil,
		})
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(payload))
	}

	payload, _ := json.Marshal(map[string]interface{}{
		"error": false,
		"msg":   nil,
		"user":  user,
	})
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(payload))
}

// CreateUser func for creates a new user.
// @Description Create a new user.
// @Summary create a new user
// @Tags User
// @Accept json
// @Produce json
// @Param email body string true "E-mail"
// @Success 200 {object} models.User
// @Security ApiKeyAuth
// @Router /v1/user [post]
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Define content type and CORS.
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Get now time.
	now := time.Now().Unix()

	// Get data from JWT.
	token := r.Context().Value("jwt").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	// Set expiration time from JWT data of current user.
	expires := claims["expires"].(int64)

	// Set credential `user:create` from JWT data of current user.
	credential := claims["user:create"].(bool)

	// Create a new user struct.
	user := &models.User{}

	// Checking received data from JSON body.
	if err := r.ParseForm(); err != nil {
		// Return status 500 and database connection error.
		payload, _ := json.Marshal(map[string]interface{}{
			"error": true,
			"msg":   err.Error(),
		})
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(payload))
	}

	// Only user with `user:create` credential can create a new user profile.
	if credential && now < expires {
		// Create a new validator for user model.
		validate := validators.UserValidator()

		// Validate user fields.
		if err := validate.Struct(user); err != nil {
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

		// Set initialized default data for user:
		user.ID = uuid.New()
		user.CreatedAt = time.Now()
		user.UpdatedAt = time.Time{}
		user.UserStatus = 1 // 0 == blocked, 1 == active
		user.UserAttrs = models.UserAttrs{}

		// Create a new user with validated data.
		if err := db.CreateUser(user); err != nil {
			// Return status 500 and database connection error.
			payload, _ := json.Marshal(map[string]interface{}{
				"error": true,
				"msg":   err.Error(),
			})
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(payload))
		}
	} else {
		// Return status 500 and database connection error.
		payload, _ := json.Marshal(map[string]interface{}{
			"error": true,
			"msg":   "permission denied, check credentials or expiration time of your token",
			"user":  nil,
		})
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte(payload))
	}

	payload, _ := json.Marshal(map[string]interface{}{
		"error": false,
		"msg":   nil,
		"user":  user,
	})
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(payload))
}

// UpdateUser func for updates user by given ID.
// @Description Update user.
// @Summary update user
// @Tags User
// @Accept json
// @Produce json
// @Param id body string true "User ID"
// @Param user_status body integer true "User status"
// @Param user_attrs body models.UserAttrs true "User attributes"
// @Success 200 {object} models.User
// @Security ApiKeyAuth
// @Router /v1/user [put]
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Define content type and CORS.
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Get now time.
	now := time.Now().Unix()

	// Get data from JWT.
	token := r.Context().Value("jwt").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	// Set expiration time from JWT data of current user.
	expires := claims["expires"].(int64)

	// Set credential `user:update` from JWT data of current user.
	credential := claims["user:update"].(bool)

	// Create a new user struct.
	user := &models.User{}

	// Checking received data from JSON body.
	if err := r.ParseForm(); err != nil {
		// Return status 500 and database connection error.
		payload, _ := json.Marshal(map[string]interface{}{
			"error": true,
			"msg":   err.Error(),
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

	// Checking, if user with given ID is exists.
	if _, err := db.GetUser(user.ID); err != nil {
		// Return status 404 and user not found error.
		payload, _ := json.Marshal(map[string]interface{}{
			"error": true,
			"msg":   "user not found",
		})
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(payload))
	}

	// Only user with `user:update` credential can update user profile.
	if credential && now < expires {
		// Create a new validator for user model.
		validate := validators.UserValidator()

		// Validate user fields.
		if err := validate.Struct(user); err != nil {
			// Return status 500 and fields are not valid.
			payload, _ := json.Marshal(map[string]interface{}{
				"error": true,
				"msg":   utils.ValidatorErrors(err),
			})
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(payload))
		}

		// Set user data to update:
		user.UpdatedAt = time.Now()

		// Update user.
		if err := db.UpdateUser(user); err != nil {
			// Return status 500 and database connection error.
			payload, _ := json.Marshal(map[string]interface{}{
				"error": true,
				"msg":   err.Error(),
			})
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(payload))
		}
	} else {
		// Return status 403 and permission denied error.
		payload, _ := json.Marshal(map[string]interface{}{
			"error": true,
			"msg":   "permission denied, check credentials or expiration time of your token",
			"user":  nil,
		})
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte(payload))
	}

	payload, _ := json.Marshal(map[string]interface{}{
		"error": false,
		"msg":   nil,
		"user":  user,
	})
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(payload))
}

// DeleteUser func for deletes user by given ID.
// @Description Delete user by given ID.
// @Summary delete user by given ID
// @Tags User
// @Accept json
// @Produce json
// @Param id body string true "User ID"
// @Success 200 {string} string "ok"
// @Security ApiKeyAuth
// @Router /v1/user [delete]
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Define content type and CORS.
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Get now time.
	now := time.Now().Unix()

	// Get data from JWT.
	token := r.Context().Value("jwt").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	// Set expiration time from JWT data of current user.
	expires := claims["expires"].(int64)

	// Set credential `user:delete` from JWT data of current user.
	credential := claims["user:delete"].(bool)

	// Create new User struct
	user := &models.User{}

	// Check, if received JSON data is valid.
	if err := r.ParseForm(); err != nil {
		// Return status 500 and JSON parse error.
		payload, _ := json.Marshal(map[string]interface{}{
			"error": true,
			"msg":   err.Error(),
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

	// Checking, if user with given ID is exists.
	if _, err := db.GetUser(user.ID); err != nil {
		// Return status 404 and user not found error.
		payload, _ := json.Marshal(map[string]interface{}{
			"error": true,
			"msg":   "user not found",
		})
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(payload))
	}

	// Only user with `user:delete` credential can delete user profile.
	if credential && now < expires {
		// Delete user by given ID.
		if err := db.DeleteUser(user.ID); err != nil {
			// Return status 500 and delete user process error.
			payload, _ := json.Marshal(map[string]interface{}{
				"error": true,
				"msg":   err.Error(),
			})
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(payload))
		}
	} else {
		// Return status 403 and permission denied error.
		payload, _ := json.Marshal(map[string]interface{}{
			"error": true,
			"msg":   "permission denied, check credentials or expiration time of your token",
		})
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte(payload))
	}

	payload, _ := json.Marshal(map[string]interface{}{
		"error": false,
		"msg":   nil,
	})
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(payload))
}
