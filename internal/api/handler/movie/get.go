package movie

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// @Summary Get all movies
// @Description Retrieve all movies
// @Tags Movies
// @Produce json
// @Success 200 {array} model.Movie
// @Failure 500 {object} map[string]string
// @Router /movies [get]
func (h *Handler) GetAll(c *fiber.Ctx) error {
	movies, err := h.service.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(movies)
}

// @Summary Get movie by ID
// @Description Retrieve a movie by its ID
// @Tags Movies
// @Produce json
// @Param id path int true "Movie ID"
// @Success 200 {object} model.Movie
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /movies/{id} [get]
func (h *Handler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	movie, err := h.service.GetByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(movie)
}