package db

import (
	"github.com/eliasyoung/fiber-flavor/internal/model"
	"github.com/google/uuid"
)

func (s *Store) CreateUser(username, email string) (*model.User, error) {
	user := model.User{
		Username: username,
		Email:    email,
	}

	result := s.DB.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (s *Store) GetAllUsers() ([]model.User, error) {
	var users []model.User

	result := s.DB.Find(&users)
	if result.Error != nil {
		return users, result.Error
	}

	return users, nil
}

func (s *Store) GetUserById(uid uuid.UUID) (model.User, error) {
	var user model.User

	result := s.DB.Where("id = ?", uid).First(&user)
	if result.Error != nil {
		return model.User{}, result.Error
	}

	// txErr := s.DB.Transaction(func(tx *gorm.DB) error {
	// 	var createdUser model.User
	// 	result := tx.Where("id = ?", uid).First(&createdUser)
	// 	if result.Error != nil {
	// 		return result.Error
	// 	}

	// 	user = createdUser
	// 	return nil
	// })
	// if txErr != nil {
	// 	return model.User{}, txErr
	// }

	return user, nil
}
