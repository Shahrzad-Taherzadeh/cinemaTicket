package memory

import (
	"errors"
	"time"

	"github.com/Shahrzad-Taherzadeh/cinemaTicket/internal/model"
)

type MovieRepository struct {
	store *Store
}

func NewMovieRepository(store *Store) *MovieRepository {
	return &MovieRepository{store: store}
}

func (r *MovieRepository) Create(m *model.Movie) error {
	r.store.mu.Lock()
	defer r.store.mu.Unlock()

	id := r.store.GenerateID("movie")
	m.ID = id
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()

	r.store.Movies[id] = m
	return nil
}

func (r *MovieRepository) GetByID(id int) (*model.Movie, error) {
	r.store.mu.Lock()
	defer r.store.mu.Unlock()

	movie, ok := r.store.Movies[id]
	if !ok {
		return nil, errors.New("Movie not found")
	}
	return movie, nil
}

func (r *MovieRepository) GetList() ([]*model.Movie, error) {
	r.store.mu.Lock()
	defer r.store.mu.Unlock()

	var moviesList []*model.Movie
	for _, m := range r.store.Movies {
		moviesList = append(moviesList, m)
	}
	return moviesList, nil

}

func (r *MovieRepository) UpdateMovie(m *model.Movie) error {
	r.store.mu.Lock()
	defer r.store.mu.Unlock()

	_, ok := r.store.Movies[m.ID]
	if !ok {
		return errors.New("Movie not found")
	}

	m.UpdatedAt = time.Now()
	r.store.Movies[m.ID] = m
	return nil
}

func (r *MovieRepository) DeleteMovie(id int) error {
	r.store.mu.Lock()
	defer r.store.mu.Unlock()

	_, ok := r.store.Movies[id]
	if !ok {
		return errors.New("Movie not found")
	}

	delete(r.store.Movies, id)
	return nil
}