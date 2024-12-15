package services

import (
	"chat_api/models"
	"chat_api/repositories"
)

type UserService struct {
	UserRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return UserService{UserRepository: userRepository}
}

func (us *UserService) CreateUser(user models.User) (models.User, error) {
	createdUser, err := us.UserRepository.Create(&user)
	if err != nil {
		return models.User{}, err
	}
	return *createdUser, nil
}
