package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type Handler struct {
	logger *zap.Logger
}

func NewHandler(logger *zap.Logger) *Handler {
	return &Handler{
		logger: logger,
	}
}

// placeholder for the sake of prototype, a real cas would be a lot better than this
func (h *Handler) Authenticate(c *gin.Context) {
	key := c.Query("key")

	context, err := validateKey(key)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"context": context})
}

func validateKey(key string) (string, error) {
	keys := map[string]string{
		"operatorKey123":  fmt.Sprintf("operator.%d", 0),
		"hospitalKey456":  fmt.Sprintf("hospital.%d", 1),
		"ambulanceKey789": fmt.Sprintf("ambulance.%d", 1),
		"ambulanceKey123": fmt.Sprintf("ambulance.%d", 2),
	}

	if _, ok := keys[key]; ok {
		return keys[key], nil
	}

	return "", errors.New("invalid key")
}
