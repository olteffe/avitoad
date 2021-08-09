package queries

import (
	"errors"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/olteffe/avitoad/internal/models"
	"github.com/olteffe/avitoad/internal/utils"
)

// AdQueries struct for queries from Ads model.
type AdQueries struct {
	*sqlx.DB
}

// GetAds func for getting ads by given sorting type and limit per page.
func (q *AdQueries) GetAds(params models.FetchParam) (res []models.Ads, nextCursor string, err error) {
	queryBuilder := sq.Select("id", "name", "created_at").From("ads").PlaceholderFormat(sq.Dollar).OrderBy("created_at DESC, id DESC")

	if params.Limit > 0 {
		queryBuilder = queryBuilder.Limit(params.Limit)
	}

	if params.Cursor != "" {
		createdCursor, adID, errCsr := utils.DecodeCursor(params.Cursor)
		if errCsr != nil {
			err = errors.New("invalid-cursor")
			return
		}
		queryBuilder = queryBuilder.Where(sq.LtOrEq{
			"created_at": createdCursor,
		})
		queryBuilder = queryBuilder.Where(sq.Lt{
			"id": adID,
		})
	}

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return
	}

	rows, err := q.Query(query, args...)
	if err != nil {
		return
	}

	// Create ads
	res = []models.Ads{}
	var createdTime time.Time // only using one for all loops, we only need the latest one in the end
	for rows.Next() {
		var item models.Ads
		err = rows.Scan(
			&item.ID,
			&item.Name,
			&item.Price,
			&item.FirstPhoto,
			&item.CreatedAt,
		)
		if err != nil {
			return
		}
		createdTime = item.CreatedAt
		res = append(res, item)
	}

	// Get encode cursor
	if len(res) > 0 {
		nextCursor = utils.EncodeCursor(createdTime, res[len(res)-1].ID.String())
	}
	return
}

// GetAd func for getting one ad by given ID.
func (q *AdQueries) GetAd(id uuid.UUID, fields bool) (models.Ads, error) {
	// Define ad variable.
	var ad models.Ads

	if fields {
		// Send full-field query to database.
		if err := q.Get(&ad, `SELECT id, name, about, photos, price, first_photo FROM ads WHERE id = $1`, id); err != nil {
			return models.Ads{}, err
		}
		return ad, nil
	}
	// Send query to database without additional fields.
	if err := q.Get(&ad, `SELECT id, name, price, first_photo FROM ads WHERE id = $1`, id); err != nil {
		return models.Ads{}, err
	}
	return ad, nil
}

// CreateAd func for creating ad by given Ad object.
func (q *AdQueries) CreateAd(a *models.Ads) error {
	// Send query to database.
	if _, err := q.Exec(
		`INSERT INTO ads VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		a.ID,
		a.Name,
		a.About,
		a.Photos,
		a.Price,
		a.CreatedAt,
		a.Photos[0],
	); err != nil {
		return err
	}

	return nil
}
