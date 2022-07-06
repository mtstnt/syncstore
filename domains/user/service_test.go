package user

import (
	"errors"
	"fmt"
	"reflect"
	"syncstore/helpers"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockRepository struct {
	mock.Mock
}

func (m MockRepository) GetByUsername(username string) (*User, error) {
	args := m.Called(username)
	user, ok := args.Get(0).(*User)
	if !ok {
		user = nil
	}
	return user, args.Error(1)
}

func (m MockRepository) CreateUser(username, password string) (*User, error) {
	args := m.Called(username, password)
	user, ok := args.Get(0).(*User)
	if !ok {
		user = nil
	}
	return user, args.Error(1)
}

var testUser = &User{
	Username: "testUser",
	Password: "$2a$12$qtnNIZz0CyoK3LmJlKJEYuRjFDfVJAqKrKzdYSAEKzCs27gqs2T16",
	Status:   true,
}

func TestService_Login(t *testing.T) {
	m := MockRepository{}

	m.On("GetByUsername", "testUser").Return(testUser, nil)
	m.On("GetByUsername", "tes").Return(nil, gorm.ErrRecordNotFound)

	type args struct {
		username        string
		encodedPassword string
	}

	tests := []struct {
		name       string
		s          Service
		args       args
		want1      *fiber.Error
		shouldFail bool
	}{
		{
			name:       "it logs in the correct user",
			s:          NewService(m),
			args:       args{username: "testUser", encodedPassword: "MTIzNDU2"},
			shouldFail: false,
			want1:      nil,
		},
		{
			name:       "it fails to log in with incorrect password",
			s:          NewService(m),
			args:       args{username: "testUser", encodedPassword: "MTIzNDU2123123213"},
			shouldFail: true,
			want1:      helpers.MakeError(fiber.StatusUnauthorized, fmt.Errorf("invalid credentials")),
		},
		{
			name:       "it fails to log in with nonexistent username",
			s:          NewService(m),
			args:       args{username: "tes", encodedPassword: "MTIzNDU2"},
			shouldFail: true,
			want1:      helpers.MakeError(fiber.StatusUnauthorized, fmt.Errorf("invalid credentials")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Login(tt.args.username, tt.args.encodedPassword)
			if tt.shouldFail && got != "" {
				t.Errorf("Service.Login() got = %v, want empty string", got)
			} else if !tt.shouldFail && got == "" {
				t.Errorf("Service.Login() got = %v, want non-empty string", got)
			}

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Service.Login() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestService_Register(t *testing.T) {
	m := MockRepository{}

	m.On("CreateUser", "testUser", mock.Anything).Return(testUser, nil)
	m.On("CreateUser", "duplicateUser", mock.Anything).Return(nil, errors.New("duplicate key"))

	type args struct {
		username        string
		encodedPassword string
	}

	tests := []struct {
		name       string
		s          Service
		args       args
		user       *User
		fErr       *fiber.Error
		shouldFail bool
	}{
		{
			name:       "it registers a user",
			s:          NewService(m),
			args:       args{username: "testUser", encodedPassword: "MTIzNDU2"},
			user:       testUser,
			fErr:       nil,
			shouldFail: false,
		},
		{
			name:       "it does not register a duplicate username",
			s:          NewService(m),
			args:       args{username: "duplicateUser", encodedPassword: "MTIzNDU2"},
			user:       nil,
			fErr:       helpers.MakeError(fiber.StatusInternalServerError, fmt.Errorf("duplicate key")),
			shouldFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, token, err := tt.s.Register(tt.args.username, tt.args.encodedPassword)
			if tt.shouldFail {
				if len(token) > 0 {
					t.Errorf("Service.Register() token = %v, want empty token", token)
				}
				if err == nil {
					t.Errorf("Service.Register() err = %v, want %v", err, tt.fErr)
				}
				if user != nil {
					t.Errorf("Service.Register() user = %v, want nil user", user)
				}
			} else {
				if !reflect.DeepEqual(user, tt.user) {
					t.Errorf("Service.Register() user = %v, want %v", user, tt.user)
				}
				if len(token) == 0 {
					t.Errorf("Service.Register() token = %v, want nonempty token", token)
				}
				if err != nil {
					t.Errorf("Service.Register() err = %v, want nil err", err)
				}
			}

		})
	}
}
