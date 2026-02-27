package movie

import (
	"time"

	"github.com/Shahrzad-Taherzadeh/cinemaTicket/internal/model"
)

func (s *Service) Create(input model.Movie) (model.Movie, error) {
	input.CreatedAt = time.Now()
	input.UpdatedAt = time.Now()
	return s.repo.Create(input)
}