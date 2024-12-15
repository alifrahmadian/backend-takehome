package services

import (
	"app/internal/models"
	r "app/internal/repositories"
)

type PostService interface {
	CreatePost(post *models.Post) (*models.Post, error)
	GetPostByID(id int64) (*models.Post, error)
	GetAllPosts() ([]*models.Post, error)
	UpdatePost(post *models.Post) (*models.Post, error)
	DeletePost(id int64) error
}

type postService struct {
	PostRepo r.PostRepository
}

func NewPostService(postRepo r.PostRepository) PostService {
	return &postService{
		PostRepo: postRepo,
	}
}

func (s *postService) CreatePost(post *models.Post) (*models.Post, error) {
	newPost, err := s.PostRepo.CreatePost(post)
	if err != nil {
		return nil, err
	}

	return newPost, nil
}

func (s *postService) GetPostByID(id int64) (*models.Post, error) {
	post, err := s.PostRepo.GetPostByID(id)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (s *postService) GetAllPosts() ([]*models.Post, error) {
	posts, err := s.PostRepo.GetAllPosts()
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (s *postService) UpdatePost(post *models.Post) (*models.Post, error) {
	post, err := s.PostRepo.UpdatePost(post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (s *postService) DeletePost(id int64) error {
	return s.PostRepo.DeletePost(id)
}
