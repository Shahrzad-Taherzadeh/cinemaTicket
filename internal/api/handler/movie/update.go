package movie

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/Shahrzad-Taherzadeh/cinemaTicket/internal/model"
)

func (h *Handler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	var input model.Movie
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid body",
		})
	}

	updated, err := h.service.Update(id, input)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(updated)
}