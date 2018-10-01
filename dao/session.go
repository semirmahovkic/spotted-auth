package dao

import (
	"time"

	"github.com/x64puzzle/spotted-common/storage"
)

// Session Data Access Object
type Session struct{}

// Create new session
func (s *Session) Create(key, val string) error {
	if err := storage.Redis.Set(key, val, 24*time.Hour).Err(); err != nil {
		return err
	}

	return nil
}

// Get value from session
func (s *Session) Get(key string) (string, error) {
	val, err := storage.Redis.Get(key).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}

// Destroy session
func (s *Session) Destroy(key string) error {
	if err := storage.Redis.Del(key).Err(); err != nil {
		return err
	}

	return nil
}
