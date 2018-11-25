package dao

import (
	"github.com/x64integer/go-common/storage"
	"github.com/x64integer/go-common/storage/cache"
	"github.com/x64integer/go-common/util"
)

// Session Data Access Object
type Session struct{}

// Create new session
func (s *Session) Create(key, val string) error {
	if err := storage.Cache.Store(&cache.Item{
		Key:        key,
		Value:      val,
		Expiration: util.LoginExp,
	}); err != nil {
		return err
	}

	return nil
}

// Get value from session
func (s *Session) Get(key string) (string, error) {
	val, err := storage.Cache.Get(&cache.Item{Key: key})
	if err != nil {
		return "", err
	}

	return string(val), nil
}

// Destroy session
func (s *Session) Destroy(key string) error {
	if err := storage.Cache.Delete(&cache.Item{Key: key}); err != nil {
		return err
	}

	return nil
}
