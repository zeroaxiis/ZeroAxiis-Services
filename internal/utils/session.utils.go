package utils

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/zeroaxiis/ZeroAxiis-Services/internal/database"
)

const (
	SessionPrefix = "session:"
	SessionTTL    = 15 * time.Minute
)

func CreateSession(adminID string) (string, error) {
	sessionID := uuid.NewString()

	key := SessionPrefix + sessionID

	err := database.RedisClient.Set(
		context.Background(),
		key,
		adminID,
		SessionTTL,
	).Err()

	if err != nil {
		return "", err
	}

	return sessionID, nil
}

func DeleteSession(sessionToken string) (string, error) {
	return "", nil
}

func RefreshSession(sessionToken string) (string, error) {
	return "", nil
}

func GetSession(sessionToken string) (string, error) {
	return "", nil
}
