package dtos

import "time"

type AddCommentRequest struct {
	Content string `json:"content" binding:"required"`
}

type AddCommentResponse struct {
	ID         int64     `json:"id"`
	PostID     int64     `json:"post_id"`
	AuthorName string    `json:"author_name"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
}

type GetCommentResponse struct {
	ID         int64     `json:"id"`
	AuthorName string    `json:"author_name"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
}
