package clients

import "time"

type Client struct {
	ID        string
	Name      string
	CreatedAt time.Time
}

type Repository interface {
	LoadClientByID(id string) (*Client, error)
}

type PlainService struct {
	repo Repository
}

func (s *PlainService) GetClientByID(id string) (*Client, error) {
	return s.repo.LoadClientByID(id)
}

func NewService(repo Repository) *PlainService {
	return &PlainService{repo}
}
