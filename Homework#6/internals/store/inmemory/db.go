package inmemory

import (
	"context"
	"fmt"
	"sync"

	"github.com/rmbziiik/Golang-application-development/Homework#6/internal/models"
	"github.com/rmbziiik/Golang-application-development/Homework#6/internal/store"
)

type DB struct {
	data map[int]*models.Ticket

	mu *sync.RWMutex
}

func NewDB() store.Store {
	return &DB{
		data: make(map[int]*models.Ticket),
		mu:   new(sync.RWMutex),
	}
}

func (db *DB) Create(ctx context.Context, ticket *models.Ticket) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.data[ticket.ID] = ticket
	return nil
}

func (db *DB) All(ctx context.Context) ([]*models.Ticket, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	tickets := make([]*models.Ticket, 0, len(db.data))
	for _, Ticket := range db.data {
		tickets = append(tickets, ticket)
	}

	return tickets, nil
}

func (db *DB) ByID(ctx context.Context, id int) (*models.Ticket, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	ticket, ok := db.data[id]
	if !ok {
		return nil, fmt.Errorf("No ticket with id %d", id)
	}

	return ticket, nil
}

func (db *DB) Update(ctx context.Context, ticket *models.Ticket) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.data[ticket.ID] = ticket
	return nil
}

func (db *DB) Delete(ctx context.Context, id int) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	delete(db.data, id)
	return nil
}
