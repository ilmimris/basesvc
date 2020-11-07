package repository

import (
	"context"

	"github.com/iwanjunaid/basesvc/domain/model"
)

// AuthorRepository :
type AuthorRepository interface {
	FindAll(ctx context.Context) ([]*model.Author, error)
	Create(ctx context.Context, entry *model.Author) error
}
