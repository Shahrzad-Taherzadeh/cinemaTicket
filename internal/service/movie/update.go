package movie

import (
	"time"

	"github.com/Shahrzad-Taherzadeh/cinemaTicket/internal/model"
)

func (s *Service) Update(id int, input model.Movie) (model.Movie, error) {
	input.UpdatedAt = time.Now()
	return s.repo.Update(id, input)
}