package pg

import "github.com/olteffe/avitoad/internal/queries"

// Queries struct for collect all app queries.
type Queries struct {
	*queries.AdQueries // load queries from User model
}

// OpenDBConnection func for opening database connection.
func OpenDBConnection() (*Queries, error) {
	// Define a new PostgreSQL connection.
	db, err := PostgreSQLConnection()
	if err != nil {
		return nil, err
	}

	return &Queries{
		// Set queries from models:
		AdQueries: &queries.AdQueries{DB: db}, // from user model
	}, nil
}
