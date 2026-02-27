package movie

import (
	"sort"

	"github.com/Shahrzad-Taherzadeh/cinemaTicket/internal/model"
)

// Get by ID
func (s *Service) GetByID(id int) (model.Movie, error) {
	return s.repo.GetByID(id)
}

// Get all movies (sorted by ID ascending)
func (s *Service) GetAll() ([]model.Movie, error) {
	movies, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	// extra sort in service layer (double safe)
	sort.Slice(movies, func(i, j int) bool {
		return movies[i].ID < movies[j].ID
	})
	return movies, nil
}