package services

import (
	"blog-post/internal/models"
	"blog-post/internal/repo"
)

type PostService struct {
	repo repo.PostRepository
}

func NewPostServices(repo repo.PostRepository) *PostServices {
	return &PostService{repo: repo}
}

func (s *PostService) CreatePost(post *models.Post) error {
	return s.repo.CreatePost(post)
}

// ...