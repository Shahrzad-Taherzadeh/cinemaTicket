package movie

import (
	"time"

	"github.com/Shahrzad-Taherzadeh/cinemaTicket/internal/model"
	"github.com/Shahrzad-Taherzadeh/cinemaTicket/internal/service/movie"
	"github.com/gofiber/fiber/v2"
)

type MovieHandler struct {
	service *movie.Service
}

func NewMovieHandler(s *movie.Service) *MovieHandler {
	return &MovieHandler{service: s}
}

func (h *MovieHandler) Create(c *fiber.Ctx) error {
	var req model.Movie

	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	req.CreatedAt = time.Now()
	req.UpdatedAt = time.Now()

	if err := h.service.Create(&req); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(req)
}
