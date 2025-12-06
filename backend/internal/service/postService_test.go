package service

import (
	"github.com/SemgaTeam/blog/mock"
	"github.com/SemgaTeam/blog/internal/entities"
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
