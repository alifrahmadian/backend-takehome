package models

import "time"

type Comment struct {
	ID         int64     `json:"id"`
	PostID     int64     `json:"post_id"`
	Post       Post      `json:"post"`
	AuthorName string    `json:"author_name"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
}
