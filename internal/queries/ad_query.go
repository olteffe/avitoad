package queries

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/olteffe/avitoad/internal/models"
)

// AdQueries struct for queries from Ads model.
type AdQueries struct {
	*sqlx.DB
}

// GetAds func for getting all ads.
func (q *AdQueries) GetAds() ([]models.Ads, error) {
	// Define ads variable.
	var ads []models.Ads

	// Send query to database.
	if err := q.Select(&ads, `SELECT * FROM ads`); err != nil {
		return []models.Ads{}, err
	}

	return ads, nil
}

// GetAd func for getting one ad by given ID.
func (q *AdQueries) GetAd(id uuid.UUID) (models.Ads, error) {
	// Define ad variable.
	var ad models.Ads

	// Send query to database.
	if err := q.Get(&ad, `SELECT * FROM ads WHERE id = $1`, id); err != nil {
		return models.Ads{}, err
	}
	return ad, nil
}

// CreateAd func for creating ad by given Ad object.
func (q *AdQueries) CreateAd(a *models.Ads) error {
	// Send query to database.
	if _, err := q.Exec(
		`INSERT INTO ads VALUES ($1, $2, $3, $4, $5, $6)`,
		a.ID,
		a.Name,
		a.About,
		a.Photos,
		a.Price,
		a.CreatedAt,
	); err != nil {
		return err
	}

	return nil
}
