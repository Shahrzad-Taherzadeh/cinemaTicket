package movie

import (
	"github.com/gofiber/fiber/v2"
	"github.com/Shahrzad-Taherzadeh/cinemaTicket/internal/model"
)

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