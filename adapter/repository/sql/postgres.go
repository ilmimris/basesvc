package sql

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/iwanjunaid/basesvc/internal/telemetry"

	"github.com/RoseRocket/xerrs"

	"github.com/jmoiron/sqlx"

	"github.com/iwanjunaid/basesvc/domain/model"
	"github.com/iwanjunaid/basesvc/usecase/author/repository"
	uuid "github.com/satori/go.uuid"
)

const (
	authorsTable = "authors"
)

type AuthorSQLRepositoryImpl struct {
	db *sqlx.DB
}

func (as *AuthorSQLRepositoryImpl) fetch(ctx context.Context, query string, args ...interface{}) (result []*model.Author, err error) {
	rows, err := as.db.QueryContext(ctx, query, args...)

	if err != nil {
		err = xerrs.Mask(err, errors.New("error query select"))
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var (
			ID                   uuid.UUID
			name, email          string
			createdAt, updatedAt time.Time
		)
		err := rows.Scan(&ID, &name, &email, &createdAt, &updatedAt)

		if err != nil {
			err = xerrs.Mask(err, errors.New("error query select"))
			return nil, err
		}

		result = append(result, &model.Author{
			ID:        ID,
			Name:      name,
			Email:     email,
			CreatedAt: createdAt.Unix(),
			UpdatedAt: updatedAt.Unix(),
		})
	}

	return result, nil
}

func (as *AuthorSQLRepositoryImpl) Find(ctx context.Context, id string) (author *model.Author, err error) {
	query := fmt.Sprintf(`SELECT id, name, email, created_at, updated_at FROM %s WHERE id = $1`, authorsTable)

	ds := telemetry.StartDataSegment(ctx, map[string]interface{}{
		"collection":   authorsTable,
		"operation":    "READ",
		"query":        query,
		"query_params": map[string]interface{}{},
	})

	list, err := as.fetch(ctx, query, id)
	ds.End()
	if err != nil {
		err = xerrs.Mask(err, errors.New("error query select"))
		return author, err
	}

	if len(list) > 0 {
		author = list[0]
	} else {
		err = xerrs.Mask(err, errors.New("not found"))
		return author, err
	}

	return author, nil
}

func (as *AuthorSQLRepositoryImpl) FindAll(ctx context.Context) ([]*model.Author, error) {

	var authors []*model.Author
	query := fmt.Sprintf(`SELECT id, name, email, created_at, updated_at FROM %s`, authorsTable)
	ds := telemetry.StartDataSegment(ctx, map[string]interface{}{
		"collection":   authorsTable,
		"operation":    "READ",
		"query":        query,
		"query_params": map[string]interface{}{},
	})

	// return authors, nil
	authors, err := as.fetch(ctx, query)
	ds.End()
	if err != nil {
		err = xerrs.Mask(err, errors.New("error query select"))
		return authors, err
	}

	return authors, nil
}

func (as *AuthorSQLRepositoryImpl) Create(ctx context.Context, author *model.Author) (*model.Author, error) {
	var (
		id        = uuid.NewV4()
		createdAt = time.Now()
		updatedAt = time.Now()
	)

	query := fmt.Sprintf(`INSERT INTO %s 
	(id, name, email, created_at, updated_at) 
	VALUES 
	($1, $2, $3, $4, $5)`, authorsTable)

	ds := telemetry.StartDataSegment(ctx, map[string]interface{}{
		"collection":   authorsTable,
		"operation":    "INSERT",
		"query":        query,
		"query_params": map[string]interface{}{},
	})

	_, err := as.db.ExecContext(ctx, query, id, author.Name, author.Email, createdAt, updatedAt)
	ds.End()
	if err != nil {
		err = xerrs.Mask(err, errors.New("error query insert"))
		return author, err
	}
	author.ID = id
	author.CreatedAt = createdAt.Unix()
	author.UpdatedAt = updatedAt.Unix()
	return author, nil
}

func NewAuthorRepository(db *sqlx.DB) repository.AuthorSQLRepository {
	repo := &AuthorSQLRepositoryImpl{
		db: db,
	}

	return repo
}
