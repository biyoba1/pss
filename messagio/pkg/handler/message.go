package handler

import (
	"github.com/biyoba1/messagio"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CreateMessage(c *gin.Context) {
	var input messagio.Message
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid type body")
		return
	}
	id, err := h.services.CreateMessage(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (h *Handler) GetAllMessages(c *gin.Context) {
	return
}

func (h *Handler) GetMessageById(c *gin.Context) {
	return
}

func (h *Handler) GetProcessedMessages(c *gin.Context) {
	return
}
func (h *Handler) GetStatistics(c *gin.Context) {
	return
}
