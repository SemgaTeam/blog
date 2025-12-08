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

func TestCreatePostValidation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPostRepo := mock.NewMockPostRepository(ctrl)
	postService := NewPostService(mockPostRepo)	

	tests := []struct{
		name string
		title string 
		content string
		authorID int
		wantError bool
		setupMock func()
	}{
		{ 
			name: "valid post", 
			title: "title", 
			content: "content", 
			authorID: 1, 
			wantError: false,
			setupMock: func() {
				mockPostRepo.
					EXPECT().
					CreatePost(gomock.Not(""), gomock.Any(), gomock.Any()).
					Return(&entities.Post{}, nil)
			},
		},
		{ 
			name: "empty title", 
			title: "", 
			content: "content", 
			authorID: 1, 
			wantError: true,
			setupMock: func() {
				mockPostRepo.
					EXPECT().
					CreatePost(gomock.Eq(""), gomock.Any(), gomock.Any()).
					Return(nil, errors.New("empty title"))
			},
		},
		{
			name: "invalid author id",
			title: "title",
			content: "content",
			authorID: -1,
			wantError: true,
			setupMock: func() {
				mockPostRepo.
					EXPECT().
					CreatePost(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil, errors.New("invalid author id"))	
			},
		},
		{ 
			name: "repository error", 
			title: "title", 
			content: "content", 
			authorID: 1, 
			wantError: true,
			setupMock: func() {
				mockPostRepo.
					EXPECT().
					CreatePost(gomock.Not(""), gomock.Any(), gomock.Any()).
					Return(nil, errors.New("repository error"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()
			_, err := postService.CreatePost(context.Background(), tt.title, tt.content, tt.authorID)

			if (err != nil) != tt.wantError {
				t.Errorf("CreatePost() error = %v, wantError %v", err, tt.wantError)
			}
		}) 
	}
}

func TestGetPostValidation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPostRepo := mock.NewMockPostRepository(ctrl)
	postService := NewPostService(mockPostRepo)	

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
				mockPostRepo.
					EXPECT().
					GetPost(gomock.Any()).
					Return(&entities.Post{}, nil)
			},
		},
		{ 
			name: "repository error", 
			id: 1,
			wantError: true,
			setupMock: func() {
				mockPostRepo.
					EXPECT().
					GetPost(gomock.Any()).
					Return(nil, errors.New("repository error"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()
			_, err := postService.GetPost(context.Background(), tt.id)

			if (err != nil) != tt.wantError {
				t.Errorf("GetPost() error = %v, wantError %v", err, tt.wantError)
			}
		}) 
	}
}

func TestGetPostsValidation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPostRepo := mock.NewMockPostRepository(ctrl)
	postService := NewPostService(mockPostRepo)	

	tests := []struct{
		name string
		params dto.GetPostParams
		total int64
		wantError bool
		setupMock func()
	}{
		{ 
			name: "valid input: IDs", 
			params: dto.GetPostParams{
				IDs: []int{1, 2, 3},
			},
			total: 3,
			wantError: false,
			setupMock: func() {
				mockPostRepo.
					EXPECT().
					GetPosts(gomock.Any()).
					Return([]entities.Post{}, int64(3), nil)
			},
		},
		{ 
			name: "valid input: Name", 
			params: dto.GetPostParams{
				Name: "post",
			},
			total: 3,
			wantError: false,
			setupMock: func() {
				mockPostRepo.
					EXPECT().
					GetPosts(gomock.Any()).
					Return([]entities.Post{}, int64(3), nil)
			},
		},
		{ 
			name: "valid input: AuthorID", 
			params: dto.GetPostParams{
				AuthorID: 1,
			},
			total: 1,
			wantError: false,
			setupMock: func() {
				mockPostRepo.
					EXPECT().
					GetPosts(gomock.Any()).
					Return([]entities.Post{}, int64(1), nil)
			},
		},
		{
			name: "empty IDs",
			params: dto.GetPostParams{
				IDs: []int{},
			},
			total: 0,
			wantError: false,
			setupMock: func() {
				mockPostRepo.
					EXPECT().
					GetPosts(gomock.Any()).
					Return(nil, int64(0), nil)
			},
		},
		{
			name: "invalid input: AuthorID",
			params: dto.GetPostParams{
				AuthorID: -1,
			},
			total: 0,
			wantError: true,
			setupMock: func() {
				mockPostRepo.
					EXPECT().
					GetPosts(gomock.Any()).
					Return(nil, int64(0), errors.New("invalid author id"))
			},
		},
		{ 
			name: "repository error", 
			params: dto.GetPostParams{
				IDs: []int{1, 2, 3},
			},
			total: 0,
			wantError: true,
			setupMock: func() {
				mockPostRepo.
					EXPECT().
					GetPosts(gomock.Any()).
					Return(nil, int64(0), errors.New("repository error"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()
			_, total, err := postService.GetPosts(context.Background(), tt.params)

			if tt.total != total {
				t.Errorf("GetPosts() total = %v, but want %v", total, tt.total)
			}

			if (err != nil) != tt.wantError {
				t.Errorf("GetPosts() error = %v, wantError %v", err, tt.wantError)
			}
		}) 
	}
}

func TestUpdatePostValidation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPostRepo := mock.NewMockPostRepository(ctrl)
	postService := NewPostService(mockPostRepo)	

	tests := []struct{
		name string
		id int
		title string 
		content string
		wantError bool
		setupMock func()
	}{
		{ 
			name: "valid input", 
			id: 1,
			title: "title", 
			content: "content", 
			wantError: false,
			setupMock: func() {
				mockPostRepo.
					EXPECT().
					UpdatePost(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(&entities.Post{}, nil)
			},
		},
		{ 
			name: "empty title", 
			id: 1,
			title: "", 
			content: "content", 
			wantError: true,
			setupMock: func() {
				mockPostRepo.
					EXPECT().
					UpdatePost(gomock.Any(), gomock.Eq(""), gomock.Any()).
					Return(nil, errors.New("empty title"))
			},
		},
		{
			name: "invalid id",
			id: -1,
			title: "title",
			content: "content",
			wantError: true,
			setupMock: func() {
				mockPostRepo.
					EXPECT().
					UpdatePost(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil, errors.New("invalid id"))	
			},
		},
		{ 
			name: "repository error", 
			id: 1, 
			title: "title", 
			content: "content", 
			wantError: true,
			setupMock: func() {
				mockPostRepo.
					EXPECT().
					UpdatePost(gomock.Not(""), gomock.Any(), gomock.Any()).
					Return(nil, errors.New("repository error"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()
			_, err := postService.UpdatePost(context.Background(), tt.id, tt.title, tt.content)

			if (err != nil) != tt.wantError {
				t.Errorf("UpdatePost() error = %v, wantError %v", err, tt.wantError)
			}
		}) 
	}
}
