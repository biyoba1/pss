package repository

import (
	"github.com/biyoba1/messagio"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jmoiron/sqlx"
)

type MessageServices interface {
	CreateMessage(message messagio.Message) (int, error)
	GetMessageById(id int) ([]messagio.Message, error)
	GetAllMessages() ([]messagio.Message, error)
	GetProcessedMessages() ([]messagio.Message, error)
	GetStatistics() ([]messagio.Message, error)
	UpdateMessageProcessed(id int, processed bool) error
}

type Repository struct {
	MessageServices
}

func NewRepository(db *sqlx.DB, kafkaProducer *kafka.Producer) *Repository {
	return &Repository{
		MessageServices: NewMessagePostgres(db, kafkaProducer),
	}
}
