package repository

import (
	"fmt"
	"github.com/biyoba1/messagio"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jmoiron/sqlx"
)

type MessagePostgres struct {
	db            *sqlx.DB
	kafkaProducer *kafka.Producer
}

func NewMessagePostgres(db *sqlx.DB, kafkaProducer *kafka.Producer) *MessagePostgres {
	return &MessagePostgres{
		db:            db,
		kafkaProducer: kafkaProducer,
	}
}

func (m *MessagePostgres) CreateMessage(message messagio.Message) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (text, processed) VALUES ($1, $2) RETURNING id", messagesTable)
	row := m.db.QueryRow(query, message.Text, message.Processed)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	topic := "my_topic" // replace with your Kafka topic
	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message.Text),
	}

	err := m.kafkaProducer.Produce(msg, nil)
	if err != nil {
		return 0, err
	}

	// Update the message as processed in the database
	err = m.UpdateMessageProcessed(id, true)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (m *MessagePostgres) GetMessageById(id int) ([]messagio.Message, error) {
	return nil, nil
}

func (m *MessagePostgres) GetAllMessages() ([]messagio.Message, error) {
	return nil, nil
}

func (m *MessagePostgres) GetProcessedMessages() ([]messagio.Message, error) {
	return nil, nil
}

func (m *MessagePostgres) GetStatistics() ([]messagio.Message, error) {
	return nil, nil
}
