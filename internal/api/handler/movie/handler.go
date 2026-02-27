package movie

import (
	"github.com/Shahrzad-Taherzadeh/cinemaTicket/internal/service/movie"
)

type Handler struct {
	service *movie.Service
}

func NewHandler(service *movie.Service) *Handler {
	return &Handler{service: service}
}
