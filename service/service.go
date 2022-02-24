package service

import "apiGO/db"

type Service struct {
	db      *db.Storage
	signKey []byte
}

func New(db *db.Storage, signKey []byte) *Service {
	return &Service{
		db:      db,
		signKey: signKey,
	}
}
