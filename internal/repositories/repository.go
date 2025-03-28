package repositories

import (
	"github.com/upper/db/v4"
)

// if needed, would make interface to be able to swap out the database, keeping same func
// would split this into each model
type Repository struct {
	db db.Session
}

func NewRepository(db db.Session) *Repository {
	return &Repository{
		db: db,
	}
}
