package orders

import "time"

type Repository interface {
	LoadOrderByID(id string) (*Order, error)
	DeleteOrderByID(id string) error
}

type Order struct {
	ID          int64
	Description string
	CreatedAt   time.Time
}

type PlainService struct {
	repo Repository
}

func (s *PlainService) GetOrderByID(id string) (*Order, error) {
	return s.repo.LoadOrderByID(id)
}

func (s *PlainService) DeleteOrderByID(id string) error {
	return s.repo.DeleteOrderByID(id)
}

func NewService(repo Repository) *PlainService {
	return &PlainService{repo}
}
