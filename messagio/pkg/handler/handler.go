package handler

import (
	"github.com/biyoba1/messagio/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

type Handler struct {
	services      *service.Service
	kafkaProducer *kafka.Writer
	kafkaConsumer *kafka.Reader
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

//CreateMessage(message messagio.Message) (int, error)
//GetMessageById(id int) ([]messagio.Message, error)
//GetAllMessages() ([]messagio.Message, error)
//GetProcessedMessages() ([]messagio.Message, error)
//GetStatistics() ([]messagio.Message, error)

func (h Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	message := router.Group("/message")
	{
		message.POST("/", h.CreateMessage)
		message.GET("/", h.GetAllMessages)
		message.GET("/:id", h.GetMessageById)
		message.GET("/processed", h.GetProcessedMessages)
		message.GET("/statistics", h.GetStatistics)
	}
	return router
}
