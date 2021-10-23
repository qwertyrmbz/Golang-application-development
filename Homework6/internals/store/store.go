package store

import (
	"context"

	"github.com/rmbziiik/Golang-application-development/Homework6/internal/models"
)

type Store interface {
	Create(ctx context.Context, ticket *models.Ticket) error
	All(ctx context.Context) ([]*models.Ticket, error)
	ByID(ctx context.Context, id int) (*models.Ticket, error)
	Update(ctx context.Context, ticket *models.Ticket) error
	Delete(ctx context.Context, id int) error
}
