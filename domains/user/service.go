package user

import (
	"encoding/base64"
	"fmt"
	"syncstore/helpers"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo userRepository
}

func NewService(r userRepository) Service {
	return Service{
		repo: r,
	}
}

func (s Service) Login(username, encodedPassword string) (string, *fiber.Error) {
	user, err := s.repo.GetByUsername(username)
	if user == nil {
		return "", helpers.MakeError(fiber.StatusUnauthorized, fmt.Errorf("invalid credentials"))
	}
	if err != nil {
		return "", helpers.MakeError(fiber.StatusInternalServerError, err)
	}

	decodedPassword, err := base64.StdEncoding.DecodeString(encodedPassword)
	if err != nil {
		return "", helpers.MakeError(fiber.StatusUnauthorized, fmt.Errorf("invalid credentials"))
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), decodedPassword); err != nil {
		return "", helpers.MakeError(fiber.StatusUnauthorized, fmt.Errorf("invalid credentials"))
	}

	token := uuid.New().String()

	return token, nil
}

func (s Service) Register(username, encodedPassword string) (*User, string, *fiber.Error) {
	password, err := base64.StdEncoding.DecodeString(encodedPassword)
	if err != nil {
		return nil, "", helpers.MakeError(fiber.StatusBadRequest, fmt.Errorf("password incorrectly formatted/encoded"))
	}

	user, err := s.repo.CreateUser(username, string(password))
	if err != nil {
		return nil, "", helpers.MakeError(fiber.StatusInternalServerError, err)
	}

	token := uuid.New().String()

	return user, token, nil
}
