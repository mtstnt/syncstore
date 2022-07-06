package tests

import (
	"fmt"
	"reflect"
	"syncstore/domains/user"
	"syncstore/session"
	"testing"
	"time"
)

var (
	testUser = &user.User{
		Username: "testUser",
		Password: "$2a$12$qtnNIZz0CyoK3LmJlKJEYuRjFDfVJAqKrKzdYSAEKzCs27gqs2T16",
		Status:   true,
	}

	testUser2 = &user.User{
		Username: "testUser2",
		Password: "$2a$12$qtnNIZz0CyoK3LmJlKJEYuRjFDfVJAqKrKzdYSAEKzCs27gqs2T16",
		Status:   true,
	}
)

func TestAddSession(t *testing.T) {
	type args struct {
		token  string
		userID uint
	}

	tests := []struct {
		name    string
		sessMap session.SessionDict
		args    args
		want    session.SessionDict
	}{
		{
			name:    "it stores a session",
			sessMap: session.SessionDict{},
			args:    args{token: "12345", userID: testUser.ID},
			want: session.SessionDict{
				"12345": session.SessionItem{
					UserID:         testUser.ID,
					ExpireDuration: 1 * time.Hour,
					StartTime:      time.Now(),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			session.AddSession(tt.sessMap, tt.args.token, tt.args.userID)
			if !reflect.DeepEqual(tt.sessMap, tt.want) {
				t.Errorf("Failed to equal session %v with %v", tt.sessMap, tt.want)
			}
		})
	}
}

func TestGetSession(t *testing.T) {
	type args struct {
		sessions session.SessionDict
		token    string
	}

	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		t.Errorf("Wrong location")
	}

	var (
		startTime1 = time.Date(2022, 1, 1, 1, 0, 0, 0, loc)
		startTime2 = time.Date(2022, 1, 1, 1, 0, 0, 0, loc)
	)

	tests := []struct {
		name        string
		args        args
		want        uint
		shouldFail  bool
		expectedErr error
	}{
		{
			name: "it gets the correct token",
			args: args{
				sessions: session.SessionDict{
					"12345561212": session.SessionItem{
						UserID:         testUser.ID,
						ExpireDuration: 1 * time.Hour,
						StartTime:      time.Now(),
					},
					"12345561213": session.SessionItem{
						UserID:         testUser2.ID,
						ExpireDuration: 1 * time.Hour,
						StartTime:      time.Now(),
					},
				},
				token: "12345561212",
			},
			want:        testUser.ID,
			shouldFail:  false,
			expectedErr: nil,
		},
		{
			name: "it knows to not accept an expired token",
			args: args{
				sessions: session.SessionDict{
					"12345561212": session.SessionItem{
						UserID:         testUser.ID,
						ExpireDuration: 1 * time.Hour,
						StartTime:      startTime1,
					},
					"12345561213": session.SessionItem{
						UserID:         testUser2.ID,
						ExpireDuration: 1 * time.Minute,
						StartTime:      startTime2,
					},
				},
				token: "12345561213",
			},
			want:        0,
			shouldFail:  true,
			expectedErr: fmt.Errorf("session expired"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := session.GetSession(tt.args.sessions, tt.args.token)
			if err != nil && !tt.shouldFail {
				t.Errorf("Excepted not error, got: %v", got)
			} else if err == nil && tt.shouldFail {
				t.Errorf("Excepted error, got: %v", got)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSession() = %v, want %v", got, tt.want)
			}
		})
	}
}
