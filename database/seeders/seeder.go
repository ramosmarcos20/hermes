package seeders

import (
	"hermes/config"
	"hermes/internal/models/entities"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RunSeed() {
	if err := UserData(); err != nil {
		log.Fatalf("Error al cargar datos de usuario: %v", err)
	}
}

func UserData() error {
	password, err := generatePassword("admin123")
	if err != nil {
		return err
	}

	user := entities.User{
		Email:    "admin@admin.com",
		UserName: "admin",
		Password: string(password),
		IsActive: true,
	}

	env := config.LoadEnv()
	db, err := config.Connection(env)
	if err != nil {
		return err
	}

	var existingUser entities.User
	if err := db.Where("email = ?", user.Email).First(&existingUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			if err := db.Create(&user).Error; err != nil {
				return err
			}
			log.Println("User Created:", user.UserName)
		} else {
			return err
		}
	}

	return nil
}
func generatePassword(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hash, nil
}
