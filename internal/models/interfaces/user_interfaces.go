package interfaces

import "hermes/internal/models/entities"

type UserInterfaces interface {
	GetALl() ([]entities.User, error)
	CreateUser(user *entities.User) error
	GetUserById(id uint) (*entities.User, error)
	GetUserByEmail(email string) (*entities.User, error)
	GetUserByUsername(username string) (*entities.User, error)
	UpdateUser(user *entities.User) error
	Delete(id uint) error
}
