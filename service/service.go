package service

import (
	"apiGO/cache"
	"apiGO/db"
)

type Service struct {
	db      *db.Storage
	signKey []byte
	cache   *cache.Cache
}

func New(db *db.Storage, cache *cache.Cache, signKey []byte) *Service {
	return &Service{
		db:      db,
		signKey: signKey,
		cache:   cache,
	}
}
