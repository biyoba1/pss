package service

import (
	"github.com/biyoba1/messagio"
	"github.com/biyoba1/messagio/pkg/repository"
)

type MessageService struct {
	repo repository.MessageServices
}

func NewMessageService(repo repository.MessageServices) *MessageService {
	return &MessageService{repo: repo}
}

func (s *MessageService) CreateMessage(message messagio.Message) (int, error) {
	return s.repo.CreateMessage(message)
}

func (s *MessageService) GetMessageById(id int) ([]messagio.Message, error) {
	return s.repo.GetMessageById(id)
}

func (s *MessageService) GetAllMessages() ([]messagio.Message, error) {
	return s.repo.GetAllMessages()
}

func (s *MessageService) GetProcessedMessages() ([]messagio.Message, error) {
	return s.repo.GetProcessedMessages()
}

func (s *MessageService) GetStatistics() ([]messagio.Message, error) {
	return s.repo.GetStatistics()
}

func (s *MessageService) UpdateMessageProcessed(id int, processed bool) error {
	return nil
}
