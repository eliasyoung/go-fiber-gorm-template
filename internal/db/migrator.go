package db

import (
	"fmt"

	"github.com/eliasyoung/fiber-flavor/internal/model"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {

	// Add postgres extension
	err := db.Exec("CREATE EXTENSION IF NOT EXISTS citext;").Error
	if err != nil {
		return fmt.Errorf("failed to create citext extension: %w", err)
	}

	err = db.AutoMigrate(
		&model.User{},
	)
	if err != nil {
		return err
	}
	return nil
}
