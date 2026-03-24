package domain

import (
	"github.com/SemgaTeam/blog/internal/config"
	"github.com/SemgaTeam/blog/internal/entities"
	"github.com/SemgaTeam/blog/mock"
	"github.com/golang/mock/gomock"

	"context"
	"errors"
	"testing"
)

func TestLogIn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTokenRepo := mock.NewMockTokenRepository(ctrl)
	mockUserRepo := mock.NewMockUserRepository(ctrl)
	mockHashRepo := mock.NewMockHashRepository(ctrl)

	conf := config.GetConfig()

	authService, err := NewAuthService(conf.Auth, mockTokenRepo, mockUserRepo, mockHashRepo)
	if err != nil {
		t.Errorf("error initialization auth service: conf = %v", conf)
	}

	tests := []struct {
		testName  string
		name      string
		password  string
		wantError bool
		setupMock func()
	}{
		{
			testName:  "success case",
			name:      "user",
			password:  "password",
			wantError: false,
			setupMock: func() {
				mockUserRepo.
					EXPECT().
					GetUserByName(gomock.Any()).
					Return(&entities.User{}, nil)

				mockHashRepo.
					EXPECT().
					IsPasswordValid(gomock.Any(), gomock.Any()).
					Return(true)

				mockTokenRepo.
					EXPECT().
					GenerateAndSignToken(gomock.Any()).
					Return(&entities.AuthToken{}, nil).
					Times(2)
			},
		},
		{
			testName:  "invalid password",
			name:      "user",
			password:  "password",
			wantError: true,
			setupMock: func() {
				mockUserRepo.
					EXPECT().
					GetUserByName(gomock.Any()).
					Return(&entities.User{}, nil)

				mockHashRepo.
					EXPECT().
					IsPasswordValid(gomock.Any(), gomock.Any()).
					Return(false)
			},
		},
		{
			testName:  "empty name",
			name:      "",
			password:  "password",
			wantError: true,
			setupMock: func() {
				mockUserRepo.
					EXPECT().
					GetUserByName(gomock.Any()).
					Return(nil, errors.New("empty name"))
			},
		},
		{
			testName:  "user not found",
			name:      "user",
			password:  "password",
			wantError: true,
			setupMock: func() {
				mockUserRepo.
					EXPECT().
					GetUserByName(gomock.Any()).
					Return(nil, errors.New("user not found"))
			},
		},
		{
			testName:  "user repository error",
			name:      "user",
			password:  "password",
			wantError: true,
			setupMock: func() {
				mockUserRepo.
					EXPECT().
					GetUserByName(gomock.Any()).
					Return(nil, errors.New("repository error"))
			},
		},
		{
			testName:  "hash repository error",
			name:      "user",
			password:  "password",
			wantError: true,
			setupMock: func() {
				mockUserRepo.
					EXPECT().
					GetUserByName(gomock.Any()).
					Return(&entities.User{}, nil)

				mockHashRepo.
					EXPECT().
					IsPasswordValid(gomock.Any(), gomock.Any()).
					Return(true)

				mockTokenRepo.
					EXPECT().
					GenerateAndSignToken(gomock.Any()).
					Return(nil, errors.New("repository error"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			tt.setupMock()
			_, _, err := authService.LogIn(context.Background(), tt.name, tt.password)

			if (err != nil) != tt.wantError {
				t.Errorf("LogIn() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}

func TestSignIn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTokenRepo := mock.NewMockTokenRepository(ctrl)
	mockUserRepo := mock.NewMockUserRepository(ctrl)
	mockHashRepo := mock.NewMockHashRepository(ctrl)

	conf := config.GetConfig()

	authService, err := NewAuthService(conf.Auth, mockTokenRepo, mockUserRepo, mockHashRepo)
	if err != nil {
		t.Errorf("error initialization auth service: conf = %v", conf)
	}

	tests := []struct {
		testName  string
		name      string
		password  string
		wantError bool
		setupMock func()
	}{
		{
			testName:  "success case",
			name:      "user",
			password:  "password",
			wantError: false,
			setupMock: func() {
				mockHashRepo.
					EXPECT().
					HashPassword(gomock.Any()).
					Return("password", nil)

				mockUserRepo.
					EXPECT().
					CreateUser(gomock.Any(), "password").
					Return(&entities.User{}, nil)

				mockTokenRepo.
					EXPECT().
					GenerateAndSignToken(gomock.Any()).
					Return(&entities.AuthToken{}, nil).
					Times(2)
			},
		},
		{
			testName:  "user already exists",
			name:      "user",
			password:  "password",
			wantError: true,
			setupMock: func() {
				mockHashRepo.
					EXPECT().
					HashPassword(gomock.Any()).
					Return("password", nil)

				mockUserRepo.
					EXPECT().
					CreateUser(gomock.Any(), "password").
					Return(nil, errors.New("user already exists"))
			},
		},
		{
			testName:  "empty name",
			name:      "",
			password:  "password",
			wantError: true,
			setupMock: func() {
				mockHashRepo.
					EXPECT().
					HashPassword(gomock.Any()).
					Return("password", nil)

				mockUserRepo.
					EXPECT().
					CreateUser(gomock.Any(), "password").
					Return(nil, errors.New("empty name"))
			},
		},
		{
			testName:  "user repository error",
			name:      "user",
			password:  "password",
			wantError: true,
			setupMock: func() {
				mockHashRepo.
					EXPECT().
					HashPassword(gomock.Any()).
					Return("password", nil)

				mockUserRepo.
					EXPECT().
					CreateUser(gomock.Any(), "password").
					Return(nil, errors.New("repository error"))
			},
		},
		{
			testName:  "hash repository error",
			name:      "user",
			password:  "password",
			wantError: true,
			setupMock: func() {
				mockHashRepo.
					EXPECT().
					HashPassword(gomock.Any()).
					Return("", errors.New("repository error"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			tt.setupMock()
			_, _, err := authService.SignIn(context.Background(), tt.name, tt.password)

			if (err != nil) != tt.wantError {
				t.Errorf("SignIn() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}

func TestRefreshTokens(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTokenRepo := mock.NewMockTokenRepository(ctrl)
	mockUserRepo := mock.NewMockUserRepository(ctrl)
	mockHashRepo := mock.NewMockHashRepository(ctrl)

	conf := config.GetConfig()

	authService, err := NewAuthService(conf.Auth, mockTokenRepo, mockUserRepo, mockHashRepo)
	if err != nil {
		t.Errorf("error initialization auth service: conf = %v", conf)
	}

	tests := []struct {
		testName  string
		id        int
		isAdmin   bool
		wantError bool
		setupMock func()
	}{
		{
			testName:  "success case (not admin)",
			id:        1,
			isAdmin:   false,
			wantError: false,
			setupMock: func() {
				mockTokenRepo.
					EXPECT().
					GenerateAndSignToken(gomock.Any()).
					Return(&entities.AuthToken{}, nil).
					Times(2)
			},
		},
		{
			testName:  "success case (admin)",
			id:        1,
			isAdmin:   true,
			wantError: false,
			setupMock: func() {
				mockTokenRepo.
					EXPECT().
					GenerateAndSignToken(gomock.Any()).
					Return(&entities.AuthToken{}, nil).
					Times(2)
			},
		},
		{
			testName:  "token repository error",
			id:        1,
			wantError: true,
			setupMock: func() {
				mockTokenRepo.
					EXPECT().
					GenerateAndSignToken(gomock.Any()).
					Return(nil, errors.New("repository error"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			tt.setupMock()
			_, _, err := authService.RefreshTokens(context.Background(), tt.id, tt.isAdmin)

			if (err != nil) != tt.wantError {
				t.Errorf("RefreshTokens() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}

func TestGetMe(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTokenRepo := mock.NewMockTokenRepository(ctrl)
	mockUserRepo := mock.NewMockUserRepository(ctrl)
	mockHashRepo := mock.NewMockHashRepository(ctrl)

	conf := config.GetConfig()

	authService, err := NewAuthService(conf.Auth, mockTokenRepo, mockUserRepo, mockHashRepo)
	if err != nil {
		t.Errorf("error initialization auth service: conf = %v", conf)
	}

	tests := []struct {
		testName  string
		userId    int
		wantError bool
		setupMock func()
	}{
		{
			testName:  "success case",
			userId:    1,
			wantError: false,
			setupMock: func() {
				mockUserRepo.
					EXPECT().
					GetUserById(gomock.Eq(1)).
					Return(&entities.User{}, nil)
			},
		},
		{
			testName:  "invalid id",
			userId:    -1,
			wantError: true,
			setupMock: func() {
				mockUserRepo.
					EXPECT().
					GetUserById(gomock.Eq(-1)).
					Return(nil, errors.New("invalid id"))
			},
		},
		{
			testName:  "user not found",
			userId:    1,
			wantError: true,
			setupMock: func() {
				mockUserRepo.
					EXPECT().
					GetUserById(gomock.Eq(1)).
					Return(nil, errors.New("user not found"))
			},
		},
		{
			testName:  "user repository error",
			userId:    1,
			wantError: true,
			setupMock: func() {
				mockUserRepo.
					EXPECT().
					GetUserById(gomock.Eq(1)).
					Return(nil, errors.New("repository error"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			tt.setupMock()
			_, err := authService.GetMe(context.Background(), tt.userId)

			if (err != nil) != tt.wantError {
				t.Errorf("GetMe() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}
