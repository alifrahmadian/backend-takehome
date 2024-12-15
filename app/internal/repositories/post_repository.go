package repositories

import (
	"app/internal/models"
	"app/pkg/errors"
	"database/sql"
)

type PostRepository interface {
	CreatePost(post *models.Post) (*models.Post, error)
	GetPostByID(id int64) (*models.Post, error)
}

type postRepository struct {
	DB *sql.DB
}

func NewPostRepository(db *sql.DB) PostRepository {
	return &postRepository{
		DB: db,
	}
}

func (r *postRepository) CreatePost(post *models.Post) (*models.Post, error) {
	query := "INSERT INTO posts(title, content, author_id, created_at) VALUES (?, ?, ?, ?)"

	result, err := r.DB.Exec(
		query,
		post.Title,
		post.Content,
		post.AuthorID,
		post.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	post.ID = id

	return post, nil
}

func (r *postRepository) GetPostByID(id int64) (*models.Post, error) {
	post := &models.Post{}

	query := `
		SELECT
			posts.id,
			title,
			content,
			author_id,
			users.id,
			users.name,
			users.email,
			posts.created_at,
			posts.updated_at
		FROM
			posts
		LEFT JOIN
			users ON author_id = users.id
		WHERE
			posts.id = ?
		LIMIT 1
	`

	err := r.DB.QueryRow(
		query,
		id,
	).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.AuthorID,
		&post.User.ID,
		&post.User.Name,
		&post.User.Email,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrPostNotFound
		}

		return nil, err
	}

	return post, nil
}
