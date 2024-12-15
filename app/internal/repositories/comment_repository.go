package repositories

import (
	"app/internal/models"
	"database/sql"
)

type CommentRepository interface {
	AddComment(comment *models.Comment) (*models.Comment, error)
	GetCommentsByPostID(postID int64) ([]*models.Comment, error)
}

type commentRepository struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepository{
		DB: db,
	}
}

func (r *commentRepository) AddComment(comment *models.Comment) (*models.Comment, error) {
	query := "INSERT INTO comments(post_id, author_name, content, created_at) VALUES (?, ?, ?, ?)"

	result, err := r.DB.Exec(
		query,
		comment.PostID,
		comment.AuthorName,
		comment.Content,
		comment.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	comment.ID = id

	return comment, nil
}

func (r *commentRepository) GetCommentsByPostID(postID int64) ([]*models.Comment, error) {
	var comments []*models.Comment

	query := "SELECT id, post_id, author_name, content, created_at FROM comments WHERE post_id = ?"

	rows, err := r.DB.Query(query, postID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		comment := &models.Comment{}
		err := rows.Scan(
			&comment.ID,
			&comment.PostID,
			&comment.AuthorName,
			&comment.Content,
			&comment.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return comments, nil
}
