package repositories

import (
	"hermes/internal/models/entities"
	"hermes/internal/models/interfaces"
	"log"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.UserInterfaces {
	return &userRepository{db: db}
}

func (r *userRepository) GetALl() ([]entities.User, error) {
	var users []entities.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) CreateUser(user *entities.User) error {
	err := r.db.Create(user).Error
	if err != nil {
		log.Printf("Error creating user: %v", err)
	}
	return err
}

func (r *userRepository) GetUserById(id uint) (*entities.User, error) {
	var user entities.User
	err := r.db.First(&user, id).Error
	if err != nil {
		log.Printf("Error getting user: %v", err)
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetUserByEmail(email string) (*entities.User, error) {
	var user entities.User
	err := r.db.First(&user, "email = ?", email).Error
	if err != nil {
		log.Printf("Error getting user: %v", err)
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetUserByUsername(username string) (*entities.User, error) {
	var user entities.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("User with this username %s not found", username)
			return nil, err
		}
		log.Printf("Error getting user by username: %v", err)
	}
	return &user, nil
}

func (r *userRepository) UpdateUser(user *entities.User) error {
	err := r.db.Save(user).Error
	if err != nil {
		log.Printf("Error updating user: %v", err)
	}
	return err
}

func (r *userRepository) Delete(id uint) error {
	err := r.db.Unscoped().Delete(&entities.User{}, id).Error
	if err != nil {
		log.Printf("Error deleting user: %v", err)
	}
	return err
}
