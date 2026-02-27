package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/Shahrzad-Taherzadeh/cinemaTicket/internal/api/handler/movie"
	"github.com/Shahrzad-Taherzadeh/cinemaTicket/internal/repository/memory"
	movieService "github.com/Shahrzad-Taherzadeh/cinemaTicket/internal/service/movie"
)

func NewServer() *fiber.App {
	app := fiber.New()

	repo := memory.NewMovieRepository()
	service := movieService.NewService(repo)
	handler := movie.NewHandler(service)

	movies := app.Group("/movies")
	movies.Post("/", handler.Create)
	movies.Get("/", handler.GetAll)
	movies.Get("/:id", handler.GetByID)
	movies.Put("/:id", handler.Update)
	movies.Delete("/:id", handler.Delete)

	return app
}