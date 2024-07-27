package service

import (
	"github.com/biyoba1/messagio"
	"github.com/biyoba1/messagio/pkg/repository"
)

type MessageServices interface {
	CreateMessage(message messagio.Message) (int, error)
	GetMessageById(id int) ([]messagio.Message, error)
	GetAllMessages() ([]messagio.Message, error)
	GetProcessedMessages() ([]messagio.Message, error)
	GetStatistics() ([]messagio.Message, error)
	UpdateMessageProcessed(id int, processed bool) error
}

type Service struct {
	MessageServices
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		MessageServices: NewMessageService(repos.MessageServices),
	}
}
