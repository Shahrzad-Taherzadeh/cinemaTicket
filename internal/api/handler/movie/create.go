package movie

import (
	"github.com/gofiber/fiber/v2"
	"github.com/Shahrzad-Taherzadeh/cinemaTicket/internal/model"
)

// @Summary Create a new movie
// @Description Create a new movie with the provided details
// @Tags Movies
// @Accept json
// @Produce json
// @Param movie body model.Movie true "Movie payload"
// @Success 201 {object} model.Movie
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /movies [post]
func (h *Handler) Create(c *fiber.Ctx) error {
	var input model.Movie

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	movie, err := h.service.Create(input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(movie)
}