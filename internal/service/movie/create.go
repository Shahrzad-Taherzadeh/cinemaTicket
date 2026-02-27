package movie

import (
	"errors"
	"github.com/Shahrzad-Taherzadeh/cinemaTicket/internal/model"
)

type Repository interface {
	Create(m *model.Movie) error
}

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) Create(m *model.Movie) error {
	if m.Name == "" {
		return errors.New("movie name cannot be empty")
	}
	if m.Year <= 1800 || m.Year > 2100 {
		return errors.New("movie year is invalid")
	}
	if m.Genre == "" {
		return errors.New("movie genre cannot be empty")
	}
	if m.Director == "" {
		return errors.New("movie director cannot be empty")
	}
	if m.DurationMinutes <= 0 {
		return errors.New("movie duration must be greater than zero")
	}

	return s.repo.Create(m)
}