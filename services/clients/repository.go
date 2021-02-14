package clients

import "database/sql"

type SQLRepository struct {
	db *sql.DB
}

func (r *SQLRepository) LoadClientByID(id string) (*Client, error) {
	row := r.db.QueryRow("SELECT id, name, created_at FROM clients WHERE id = $1", id)
	o := &Client{}
	return o, row.Scan(&o.ID, &o.Name, &o.CreatedAt)
}

func NewRepository(db *sql.DB) *SQLRepository {
	return &SQLRepository{db}
}
