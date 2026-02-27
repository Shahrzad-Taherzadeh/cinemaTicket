package movie

import "github.com/Shahrzad-Taherzadeh/cinemaTicket/internal/repository/memory"

type Service struct {
	repo memory.MovieRepository
}

func NewService(repo memory.MovieRepository) *Service {
	return &Service{repo: repo}
}
