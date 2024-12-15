package repositories

import (
	"chat_api/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{DB: db}
}

func (u UserRepository) Create(user *models.User) (*models.User, error) {
	u.DB.AutoMigrate(&models.User{})
	result := u.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
