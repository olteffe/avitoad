package models

import (
	"github.com/lib/pq"
	"time"

	"github.com/google/uuid"
)

// Ads struct describe ads object.
type Ads struct {
	ID         uuid.UUID      `db:"id" json:"id,omitempty"`
	Name       string         `db:"name" json:"name" validate:"required,lte=200"`
	About      string         `db:"about" json:"about,omitempty" validate:"required,lte=1000"`
	Photos     pq.StringArray `db:"photos" json:"photos,omitempty" validate:"required,lte=3"`
	Price      uint           `db:"price" json:"price"`
	CreatedAt  time.Time      `db:"created_at" json:"-"`
	FirstPhoto string         `db:"first_photo" json:"first_photo"`
}

type FetchParam struct {
	Limit  uint64 `validate:"required,lte=50"`
	Cursor string
	Sort   string `validate:"required,oneof=price date"`
	Asc    string `validate:"required,oneof=ASC DESC"`
}
