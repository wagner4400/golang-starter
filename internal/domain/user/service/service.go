// internal/domain/user/service.go
package service

import (
	"errors"
	"lawise-go/internal/domain/user"
	"lawise-go/internal/domain/user/repository"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	userRepository repository.UserRepository
}

func NewService(userRepository repository.UserRepository) *Service {
	return &Service{
		userRepository: userRepository,
	}
}

func (s *Service) CreateUser(input *user.CreateUserInput) (*user.User, error) {
	// Check if user already exists
	existing, _ := s.userRepository.FindByEmail(input.Email)
	if existing != nil {
		return nil, errors.New("user with this email already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create new user
	user := &user.User{
		ID:        uuid.New().String(),
		Name:      input.Name,
		Email:     input.Email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Save to repository
	if err := s.userRepository.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) GetUser(id string) (*user.User, error) {
	user, err := s.userRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (s *Service) GetUsers() ([]*user.User, error) {
	return s.userRepository.FindAll()
}

func (s *Service) UpdateUser(id string, input *user.UpdateUserInput) (*user.User, error) {
	user, err := s.userRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	if input.Name != "" {
		user.Name = input.Name
	}

	if input.Email != "" && input.Email != user.Email {
		// Check if new email is already taken
		existing, _ := s.userRepository.FindByEmail(input.Email)
		if existing != nil {
			return nil, errors.New("email already in use")
		}
		user.Email = input.Email
	}

	if input.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		user.Password = string(hashedPassword)
	}

	user.UpdatedAt = time.Now()

	if err := s.userRepository.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) DeleteUser(id string) error {
	user, err := s.userRepository.FindByID(id)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}

	return s.userRepository.Delete(id)
}

func (s *Service) ValidateCredentials(email, password string) (*user.User, error) {
	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}
