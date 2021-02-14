package orders

import "database/sql"

type SQLRepository struct {
	db *sql.DB
}

func (r *SQLRepository) LoadOrderByID(id string) (*Order, error) {
	row := r.db.QueryRow("SELECT id, description, created_at FROM orders WHERE id = $1", id)
	o := &Order{}
	return o, row.Scan(&o.ID, &o.Description, &o.CreatedAt)
}

func (r *SQLRepository) DeleteOrderByID(id string) error {
	_, err := r.db.Exec("DELETE FROM orders WHERE id = $1", id)
	return err
}

func NewRepository(db *sql.DB) *SQLRepository {
	return &SQLRepository{db}
}
