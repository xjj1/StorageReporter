package db

import "database/sql"

type Repository struct {
	db *sql.DB
}

func (r *Repository) Close() {
	r.db.Close()
}
