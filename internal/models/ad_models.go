package models

import (
	"time"

	"github.com/google/uuid"
)

// Ads struct describe ads object.
type Ads struct {
	ID        uuid.UUID `db:"id" json:"id"`
	Name      string    `db:"name" json:"name" validate:"required,lte=200"`
	About     string    `db:"about" json:"about" validate:"required,lte=1000"`
	Photos    string    `db:"photos" json:"photos" validate:"required,lte=300"`
	Price     uint      `db:"price" json:"price"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
