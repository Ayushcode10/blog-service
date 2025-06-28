package store

import (
	"database/sql"

	"github.com/Ayushcode10/blog-service/models"
)

type Store interface {
	Create(post models.Post) error
	GetAll() ([]models.Post, error)
	GetByID(id int) (models.Post, error)
}

type MySQLStore struct {
	DB *sql.DB
}

func NewMySQLStore(db *sql.DB) *MySQLStore {
	return &MySQLStore{DB: db}
}

func (s *MySQLStore) Create(post models.Post) error {
	_, err := s.DB.Exec("INSERT INTO posts (title, content, author) VALUES (?, ?, ?)", post.Title, post.Content, post.Author)
	return err
}

func (s *MySQLStore) GetAll() ([]models.Post, error) {
	rows, err := s.DB.Query("SELECT id, title, content, author FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Author)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (s *MySQLStore) GetByID(id int) (models.Post, error) {
	row := s.DB.QueryRow("SELECT id, title, content, author FROM posts WHERE id = ?", id)

	var post models.Post
	err := row.Scan(&post.ID, &post.Title, &post.Content, &post.Author)
	return post, err
}
