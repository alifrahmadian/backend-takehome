package services

import (
	"app/internal/models"
	r "app/internal/repositories"
)

type CommentService interface {
	AddComment(comment *models.Comment) (*models.Comment, error)
	GetCommentsByPostID(postID int64) ([]*models.Comment, error)
}

type commentService struct {
	CommentRepo r.CommentRepository
	PostRepo    r.PostRepository
}

func NewCommentService(commentRepo r.CommentRepository, postRepo r.PostRepository) CommentService {
	return &commentService{
		CommentRepo: commentRepo,
		PostRepo:    postRepo,
	}
}

func (s *commentService) AddComment(comment *models.Comment) (*models.Comment, error) {
	_, err := s.PostRepo.GetPostByID(comment.PostID)
	if err != nil {
		return nil, err
	}

	newComment, err := s.CommentRepo.AddComment(comment)
	if err != nil {
		return nil, err
	}

	return newComment, nil
}

func (s *commentService) GetCommentsByPostID(postID int64) ([]*models.Comment, error) {
	comments, err := s.CommentRepo.GetCommentsByPostID(postID)
	if err != nil {
		return nil, err
	}

	return comments, nil
}
