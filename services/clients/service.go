package clients

import "time"

type Client struct {
	ID        string
	Name      string
	CreatedAt time.Time
}

type Service interface {
	GetClientByID(id string) (*Client, error)
}

type service struct {
	repo Repository
}

func (s *service) GetClientByID(id string) (*Client, error) {
	return s.repo.LoadClientByID(id)
}

func NewService(repo Repository) Service {
	return &service{repo}
}
