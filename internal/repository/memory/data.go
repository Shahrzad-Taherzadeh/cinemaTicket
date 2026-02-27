package memory

import (
	"sync"
	"time"

	"github.com/Shahrzad-Taherzadeh/cinemaTicket/internal/model"
)

type Store struct {
	mu sync.RWMutex

	Movies     map[int]*model.Movie
	Halls      map[int]*model.Hall
	ShowTimes  map[int]*model.ShowTime
	Screenings map[int]*model.Screening

	nextID map[string]int
}

func NewStore() *Store {
	return &Store{
		Movies:     make(map[int]*model.Movie),
		Halls:      make(map[int]*model.Hall),
		ShowTimes:  make(map[int]*model.ShowTime),
		Screenings: make(map[int]*model.Screening),
		nextID:     make(map[string]int),
	}
}

func (s *Store) GenerateID(entity string) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.nextID[entity]++
	return s.nextID[entity]
}

func (s *Store) SeedData() {
	now := time.Now()
	s.Movies[1] = &model.Movie{
		ID:              1,
		Name:            "Oppenheimer",
		Year:            2026,
		Genre:           "Drama",
		Director:        "Christopher Nolan",
		DurationMinutes: 180,
		CreatedAt:       now,
		UpdatedAt:       now,
	}
	s.nextID["movie"] = 1
}
