package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

// Ads struct describe ads object.
type Ads struct {
	ID        uuid.UUID `db:"id" json:"id" validate:"required,id"`
	Name      string    `db:"name" json:"name" validate:"required, name"`
	About     string    `db:"about" json:"about" validate:"required, about"`
	Photos    []string  `db:"photos" json:"photos" validate:"required, photos"`
	Price     uint      `db:"price" json:"price"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

// Value make the Ads struct implement the driver.Valuer interface.
// This method simply returns the JSON-encoded representation of the struct.
func (a Ads) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan make the Ads struct implement the sql.Scanner interface.
// This method simply decodes a JSON-encoded value into the struct fields.
func (a *Ads) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &a)
}
