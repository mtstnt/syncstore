package user_test

import (
	"encoding/base64"
	"syncstore/domains/user"
	"testing"

	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

type MockRepository struct {
	mock.Mock
}

func (m MockRepository) GetByUsername(username string) (*user.User, error) {
	m.Called(username)
	return &user.User{
		Username: username,
		Password: "$2a$12$qtnNIZz0CyoK3LmJlKJEYuRjFDfVJAqKrKzdYSAEKzCs27gqs2T16",
		Status:   true,
	}, nil
}

func (m MockRepository) CreateUser(username, password string) (*user.User, error) {
	m.Called(username, password)
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return nil, err
	}

	return &user.User{
		Username: username,
		Password: string(hash),
		Status:   true,
	}, nil
}

func TestUserLogin(t *testing.T) {
	m := MockRepository{}

	m.On("GetByUsername", "username").Return(&user.User{
		Username: "username",
		Password: "$2a$12$qtnNIZz0CyoK3LmJlKJEYuRjFDfVJAqKrKzdYSAEKzCs27gqs2T16",
		Status:   true,
	}, nil)

	s := user.NewService(m)
	b64Pass := base64.StdEncoding.EncodeToString([]byte("123456"))
	token, err := s.Login("username", b64Pass)

	if err != nil || token == "" {
		t.Fail()
	}
}

func TestCreateUser(t *testing.T) {
	m := MockRepository{}

	m.On("CreateUser", "username", "123456").Return(&user.User{
		Username: "username",
		Password: "$2a$12$qtnNIZz0CyoK3LmJlKJEYuRjFDfVJAqKrKzdYSAEKzCs27gqs2T16",
		Status:   true,
	}, nil)

	s := user.NewService(m)
	b64Pass := base64.StdEncoding.EncodeToString([]byte("123456"))
	user, token, err := s.Register("username", b64Pass)
	if user.Username != "username" || err != nil {
		t.Fail()
	}

	checkPwErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte("123456"))
	if checkPwErr != nil || token == "" {
		t.Fail()
	}
}
