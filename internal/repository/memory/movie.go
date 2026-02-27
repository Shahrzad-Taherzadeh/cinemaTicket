package memory

import (
	"errors"
	"sort"
	"sync"

	"github.com/Shahrzad-Taherzadeh/cinemaTicket/internal/model"
)

var ErrMovieNotFound = errors.New("movie not found")

type MovieRepository interface {
	Create(movie model.Movie) (model.Movie, error)
	GetByID(id int) (model.Movie, error)
	GetAll() ([]model.Movie, error)
	Update(id int, movie model.Movie) (model.Movie, error)
	Delete(id int) error
}

type movieRepository struct {
	mu     sync.RWMutex
	data   map[int]model.Movie
	nextID int
}

func NewMovieRepository() MovieRepository {
	return &movieRepository{
		data:   make(map[int]model.Movie),
		nextID: 1,
	}
}

func (r *movieRepository) Create(movie model.Movie) (model.Movie, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	movie.ID = r.nextID
	r.nextID++

	r.data[movie.ID] = movie
	return movie, nil
}

func (r *movieRepository) GetByID(id int) (model.Movie, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	movie, ok := r.data[id]
	if !ok {
		return model.Movie{}, ErrMovieNotFound
	}
	return movie, nil
}

func (r *movieRepository) GetAll() ([]model.Movie, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	movies := make([]model.Movie, 0, len(r.data))
	for _, m := range r.data {
		movies = append(movies, m)
	}

	// sort by ID ascending
	sort.Slice(movies, func(i, j int) bool {
		return movies[i].ID < movies[j].ID
	})

	return movies, nil
}

func (r *movieRepository) Update(id int, movie model.Movie) (model.Movie, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.data[id]
	if !ok {
		return model.Movie{}, ErrMovieNotFound
	}

	movie.ID = id
	r.data[id] = movie
	return movie, nil
}

func (r *movieRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.data[id]; !ok {
		return ErrMovieNotFound
	}

	delete(r.data, id)
	return nil
}