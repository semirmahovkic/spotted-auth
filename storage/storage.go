package storage

import _storage "github.com/x64integer/go-common/storage"

var (
	// Cache instance exposed
	Cache = _storage.Cache
	// DB client exposed
	DB = _storage.PG
)
