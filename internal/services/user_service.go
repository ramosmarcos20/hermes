package services

import (
	"errors"
	"hermes/internal/models/entities"
	"hermes/internal/models/interfaces"
	"log"

	"gorm.io/gorm"
)

type UserService struct {
	repo interfaces.UserInterfaces
}

// Constructor para UserService
func NewUserService(repo interfaces.UserInterfaces) *UserService {
	return &UserService{repo: repo}
}

// Obtener todos los usuarios
func (s *UserService) GetAllUsers() ([]entities.User, error) {
	users, err := s.repo.GetALl()
	if err != nil {
		log.Printf("Error retrieving all users: %v", err)
		return nil, err
	}
	return users, nil
}

// Crear un usuario
func (s *UserService) CreateUser(user *entities.User) []error {
	var errs []error

	// Validar que el email no esté en uso
	existingUserByEmail, err := s.repo.GetUserByEmail(user.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("Error checking email: %v", err)
		errs = append(errs, err)
	}
	if existingUserByEmail != nil {
		errs = append(errs, errors.New("email already in use"))
	}

	// Validar que el nombre de usuario no esté en uso
	existingUserByUsername, err := s.repo.GetUserByUsername(user.UserName)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("Error checking username: %v", err)
		errs = append(errs, err)
	}
	if existingUserByUsername != nil {
		errs = append(errs, errors.New("username already in use"))
	}

	if len(errs) > 0 {
		return errs
	}

	// Crear usuario en el repositorio
	err = s.repo.CreateUser(user)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		errs = append(errs, err)
	}
	return errs
}

// Obtener un usuario por ID
func (s *UserService) GetUserById(id uint) (*entities.User, error) {
	user, err := s.repo.GetUserById(id)
	if err != nil {
		log.Printf("Error retrieving user by ID: %v", err)
		return nil, err
	}
	return user, nil
}

// Obtener un usuario por email
func (s *UserService) GetUserByEmail(email string) (*entities.User, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		log.Printf("Error retrieving user by email: %v", err)
		return nil, err
	}
	return user, nil
}

// Obtener un usuario por username
func (s *UserService) GetUserByUsername(username string) (*entities.User, error) {
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		log.Printf("Error retrieving user by username: %v", err)
		return nil, err
	}
	return user, nil
}

// Actualizar un usuario
func (s *UserService) UpdateUser(user *entities.User) error {
	err := s.repo.UpdateUser(user)
	if err != nil {
		log.Printf("Error updating user: %v", err)
		return err
	}
	return nil
}

// Eliminar un usuario
func (s *UserService) DeleteUser(id uint) error {
	err := s.repo.Delete(id)
	if err != nil {
		log.Printf("Error deleting user: %v", err)
		return err
	}
	return nil
}
