package session

import (
	"fmt"
	"time"
)

type SessionItem struct {
	UserID         uint
	ExpireDuration time.Duration
	StartTime      time.Time
}

type SessionDict map[string]SessionItem

var defaultSessMap = SessionDict{}

func GetSessionMap() SessionDict {
	return defaultSessMap
}

func AddSession(sessions SessionDict, token string, userID uint) {
	sessions[token] = SessionItem{
		UserID:         userID,
		ExpireDuration: 1 * time.Hour,
		StartTime:      time.Now(),
	}
}

func GetSession(sessions SessionDict, token string) (uint, error) {
	sessItem, exists := sessions[token]
	if !exists {
		return 0, fmt.Errorf("session not found")
	}

	if sessItem.StartTime.Add(sessItem.ExpireDuration).Unix() < time.Now().Unix() {
		RevokeSession(sessions, token)
		return 0, fmt.Errorf("session expired")
	}

	return sessItem.UserID, nil
}

func RevokeSession(sessions SessionDict, token string) (uint, error) {
	sessItem, exists := sessions[token]
	if !exists {
		return 0, fmt.Errorf("session not exist")
	}
	delete(sessions, token)
	return sessItem.UserID, nil
}
