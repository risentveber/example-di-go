package orders

import "time"

type Order struct {
	ID          int64
	Description string
	CreatedAt   time.Time
}

type Service interface {
	GetOrderByID(id string) (*Order, error)
	DeleteOrderByID(id string) error
}

type service struct {
	repo Repository
}

func (s *service) GetOrderByID(id string) (*Order, error) {
	return s.repo.LoadOrderByID(id)
}

func (s *service) DeleteOrderByID(id string) error {
	return s.repo.DeleteOrderByID(id)
}

func NewService(repo Repository) Service {
	return &service{repo}
}
