package service

import (
	"github.com/SemgaTeam/blog/internal/dto"
	"github.com/SemgaTeam/blog/internal/entities"
	"github.com/SemgaTeam/blog/mock"
	"github.com/golang/mock/gomock"

	"context"
	"errors"
	"testing"
)

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock.NewMockUserRepository(ctrl)
	userService := NewUserService(mockUserRepo)	

	tests := []struct{
		testName string
		name string 
		password string
		wantError bool
		setupMock func()
	}{
		{ 
			testName: "success case", 
			name: "user",
			password: "password",
			wantError: false,
			setupMock: func() {
				mockUserRepo.
					EXPECT().
					CreateUser(gomock.Not(""), gomock.Any()).
					Return(&entities.User{}, nil)
			},
		},
		{ 
			testName: "empty name", 
			wantError: true,
			setupMock: func() {
				mockUserRepo.
					EXPECT().
					CreateUser(gomock.Eq(""), gomock.Any()).
					Return(nil, errors.New("empty name"))
			},
		},
		{
			testName: "repository error",
			name: "user",
			password: "password",
			wantError: true,
			setupMock: func() {
				mockUserRepo.
					EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Return(nil, errors.New("repository error"))	
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			tt.setupMock()
			_, err := userService.CreateUser(context.Background(), tt.name, tt.password)

			if (err != nil) != tt.wantError {
				t.Errorf("CreateUser() error = %v, wantError %v", err, tt.wantError)
			}
		}) 
	}
}

func TestGetUserById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock.NewMockUserRepository(ctrl)
	userService := NewUserService(mockUserRepo)	

	tests := []struct{
		testName string
		id int
		wantError bool
		setupMock func()
	}{
		{ 
			testName: "success case", 
			id: 1,
			wantError: false,
			setupMock: func() {
				mockUserRepo.
					EXPECT().
					GetUserById(gomock.Eq(1)).
					Return(&entities.User{}, nil)
			},
		},
		{
			testName: "repository error",
			wantError: true,
			setupMock: func() {
				mockUserRepo.
					EXPECT().
					GetUserById(gomock.Any()).
					Return(nil, errors.New("repository error"))	
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			tt.setupMock()
			_, err := userService.GetUserById(context.Background(), tt.id)

			if (err != nil) != tt.wantError {
				t.Errorf("GetUserById() error = %v, wantError %v", err, tt.wantError)
			}
		}) 
	}
}

func TestGetUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock.NewMockUserRepository(ctrl)
	userService := NewUserService(mockUserRepo)

	tests := []struct{
		name string
		params dto.GetUserParams
		total int64
		wantError bool
		setupMock func()
	}{
		{ 
			name: "valid input: IDs", 
			params: dto.GetUserParams{
				IDs: []int{1, 2, 3},
			},
			total: 3,
			wantError: false,
			setupMock: func() {
				mockUserRepo.
					EXPECT().
					GetUsers(gomock.Any()).
					Return([]entities.User{}, int64(3), nil)
			},
		},
		{ 
			name: "valid input: Name", 
			params: dto.GetUserParams{
				Name: "user",
			},
			total: 3,
			wantError: false,
			setupMock: func() {
				mockUserRepo.
					EXPECT().
					GetUsers(gomock.Any()).
					Return([]entities.User{}, int64(3), nil)
			},
		},
		{
			name: "empty IDs",
			params: dto.GetUserParams{
				IDs: []int{},
			},
			total: 0,
			wantError: false,
			setupMock: func() {
				mockUserRepo.
					EXPECT().
					GetUsers(gomock.Any()).
					Return(nil, int64(0), nil)
			},
		},
		{ 
			name: "repository error", 
			params: dto.GetUserParams{
				IDs: []int{1, 2, 3},
			},
			total: 0,
			wantError: true,
			setupMock: func() {
				mockUserRepo.
					EXPECT().
					GetUsers(gomock.Any()).
					Return(nil, int64(0), errors.New("repository error"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()
			_, total, err := userService.GetUsers(context.Background(), tt.params)

			if tt.total != total {
				t.Errorf("GetUsers() total = %v, but want %v", total, tt.total)
			}

			if (err != nil) != tt.wantError {
				t.Errorf("GetUsers() error = %v, wantError %v", err, tt.wantError)
			}
		}) 
	}
}

func TestUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock.NewMockUserRepository(ctrl)
	userService := NewUserService(mockUserRepo)	

	tests := []struct{
		testName string
		id int
		name string
		password string
		wantError bool
		setupMock func()
	}{
		{ 
			testName: "valid input", 
			id: 1,
			name: "user",
			password: "password",
			wantError: false,
			setupMock: func() {
				mockUserRepo.
					EXPECT().
					UpdateUser(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(&entities.User{}, nil)
			},
		},
		{ 
			testName: "empty name", 
			id: 1,
			name: "",
			password: "password",
			wantError: false,
			setupMock: func() {
				mockUserRepo.
					EXPECT().
					UpdateUser(gomock.Any(), gomock.Eq(""), gomock.Any()).
					Return(&entities.User{}, nil)
			},
		},
		{ 
			testName: "empty password", 
			id: 1,
			name: "user",
			password: "",
			wantError: false,
			setupMock: func() {
				mockUserRepo.
					EXPECT().
					UpdateUser(gomock.Any(), gomock.Any(), gomock.Eq("")).
					Return(&entities.User{}, nil)
			},
		},
		{ 
			testName: "invalid id", 
			id: -1,
			name: "user",
			password: "password",
			wantError: true,
			setupMock: func() {
				mockUserRepo.
					EXPECT().
					UpdateUser(gomock.Eq(-1), gomock.Any(), gomock.Any()).
					Return(nil, errors.New("invalid id"))
			},
		},
		{ 
			testName: "repository error", 
			name: "user",
			password: "password",
			wantError: true,
			setupMock: func() {
				mockUserRepo.
					EXPECT().
					UpdateUser(gomock.Not(""), gomock.Any(), gomock.Any()).
					Return(nil, errors.New("repository error"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()
			_, err := userService.UpdateUser(context.Background(), tt.id, tt.name, tt.password)

			if (err != nil) != tt.wantError {
				t.Errorf("UpdateUser() error = %v, wantError %v", err, tt.wantError)
			}
		}) 
	}
}

func TestDeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock.NewMockUserRepository(ctrl)
	userService := NewUserService(mockUserRepo)	

	tests := []struct{
		name string
		id int
		wantError bool
		setupMock func()
	}{
		{ 
			name: "valid input", 
			id: 1,
			wantError: false,
			setupMock: func() {
				mockUserRepo.
					EXPECT().
					DeleteUser(gomock.Any()).
					Return(1, nil)
			},
		},
		{
			name: "invalid id",
			id: -1,
			wantError: true,
			setupMock: func() {
				mockUserRepo.
					EXPECT().
					DeleteUser(gomock.Any()).
					Return(0, errors.New("invalid id"))	
			},
		},
		{ 
			name: "repository error", 
			id: 1, 
			wantError: true,
			setupMock: func() {
				mockUserRepo.
					EXPECT().
					DeleteUser(gomock.Any()).
					Return(0, errors.New("repository error"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()
			_, err := userService.DeleteUser(context.Background(), tt.id)

			if (err != nil) != tt.wantError {
				t.Errorf("DeleteUser() error = %v, wantError %v", err, tt.wantError)
			}
		}) 
	}
}
