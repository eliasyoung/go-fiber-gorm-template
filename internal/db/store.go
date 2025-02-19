package db

import "gorm.io/gorm"

type Store struct {
	DB *gorm.DB
}

func InitStore(conn *gorm.DB) *Store {
	return &Store{
		DB: conn,
	}
}
