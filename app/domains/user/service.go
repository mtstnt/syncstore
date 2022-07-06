package user

import (
	"encoding/base64"
	"errors"
	"fmt"
	"syncstore/helpers"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userRepository interface {
	GetByUsername(string) (*User, error)
	CreateUser(string, string) (*User, error)
}

type Service struct {
	repo userRepository
}

func NewService(r userRepository) Service {
	return Service{
		repo: r,
	}
}

func (s Service) Login(username, encodedPassword string) (*User, string, *fiber.Error) {
	user, err := s.repo.GetByUsername(username)
	if user == nil {
		return nil, "", helpers.MakeError(fiber.StatusUnauthorized, fmt.Errorf("invalid credentials"))
	}
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "", helpers.MakeError(fiber.StatusUnauthorized, fmt.Errorf("invalid credentials"))
		}
		return nil, "", helpers.MakeError(fiber.StatusInternalServerError, err)
	}

	decodedPassword, err := base64.StdEncoding.DecodeString(encodedPassword)
	if err != nil {
		return nil, "", helpers.MakeError(fiber.StatusUnauthorized, fmt.Errorf("invalid credentials"))
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), decodedPassword); err != nil {
		return nil, "", helpers.MakeError(fiber.StatusUnauthorized, fmt.Errorf("invalid credentials"))
	}

	token := uuid.New().String()

	return user, token, nil
}

func (s Service) Register(username, encodedPassword string) (*User, string, *fiber.Error) {
	password, err := base64.StdEncoding.DecodeString(encodedPassword)
	if err != nil {
		return nil, "", helpers.MakeError(fiber.StatusBadRequest, fmt.Errorf("password incorrectly formatted/encoded"))
	}

	hashedPw, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return nil, "", helpers.MakeError(fiber.StatusInternalServerError, err)
	}

	user, err := s.repo.CreateUser(username, string(hashedPw))
	if err != nil {
		return nil, "", helpers.MakeError(fiber.StatusInternalServerError, err)
	}

	token := uuid.New().String()

	return user, token, nil
}
