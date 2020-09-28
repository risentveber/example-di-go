package clients

import "database/sql"

type Repository interface {
	LoadClientByID(id string) (*Client, error)
}

type repository struct {
	db *sql.DB
}

func (r *repository) LoadClientByID(id string) (*Client, error) {
	row := r.db.QueryRow("SELECT id, name, created_at FROM clients WHERE id = $1", id)
	o := &Client{}
	return o, row.Scan(&o.ID, &o.Name, &o.CreatedAt)
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}
