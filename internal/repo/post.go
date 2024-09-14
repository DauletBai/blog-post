package repo

import (
	"database/sql"
	"blog-post/internal/models"
)

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) CreatePost(post *models.Post) error {
	// SQL query to create a post
	_, err := r.db.Exec("INSERT INTO posts (title, content) VALUES (?, ?)", post.Title, post.Content)
	return err
}

// ...