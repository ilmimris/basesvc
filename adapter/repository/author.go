package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/iwanjunaid/basesvc/internal/logger"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/iwanjunaid/basesvc/domain/model"
	"github.com/iwanjunaid/basesvc/usecase/author/repository"
)

const authorsTable = "authors"

type AuthorRepositoryWriterImpl struct {
	db  *sql.DB
	mdb *mongo.Database
}

func (author *AuthorRepositoryWriterImpl) InsertDocument(ctx context.Context) error {
	panic("implement me")
}

func (author *AuthorRepositoryWriterImpl) Publish(ctx context.Context, topic string, message []byte) (err error) {
	panic("implement me")
}

type Option func(impl *AuthorRepositoryWriterImpl)

func NewAuthorRepositoryWriter(db *sql.DB, kp *kafka.Producer, mdb *mongo.Database) repository.AuthorRepository {
	repo := &AuthorRepositoryWriterImpl{
		db:  db,
		mdb: mdb,
	}

	return repo
}

func (author *AuthorRepositoryWriterImpl) FindAll(ctx context.Context) ([]*model.Author, error) {
	query := fmt.Sprintf("SELECT id, name, email FROM %s", authorsTable)

	rows, err := author.db.QueryContext(ctx, query)

	if err != nil {
		logger.WithFields(logger.Fields{"repository": "get authors"}).Errorf("%v", err)
		return nil, err
	}

	defer rows.Close()

	var authors []*model.Author

	for rows.Next() {
		var (
			id    sql.NullInt32
			name  sql.NullString
			email sql.NullString
		)

		err := rows.Scan(&id, &name, &email)

		if err != nil {
			logger.WithFields(logger.Fields{"repository": "get authors"}).Errorf("%v", err)
			return nil, err
		}

		authors = append(authors, &model.Author{ID: uint(id.Int32), Name: name.String, Email: email.String})
	}

	return authors, nil
}

func (author *AuthorRepositoryImpl) InsertDocument(ctx context.Context) error {
	panic("implement me")
}

func (author *AuthorRepositoryImpl) Publish(ctx context.Context, topic string, message []byte) (err error) {
	deliveryChan := make(chan kafka.Event)

	err = author.kp.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          message,
	}, deliveryChan)

	e := <-deliveryChan

	m := e.(*kafka.Message)
	if m.TopicPartition.Error != nil {
		return errors.New(fmt.Sprintf("Delivery failed: %v\n", m.TopicPartition.Error))
	} else {
		return errors.New(fmt.Sprintf("Delivered message to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset))
	}

	close(deliveryChan)
	return
}
