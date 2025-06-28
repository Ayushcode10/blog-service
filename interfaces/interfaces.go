package interfaces

import "github.com/Ayushcode10/blog-service/models"

type PostStore interface {
	Create(post models.Post) (int64, error)
	GetAll() ([]models.Post, error)
	GetByID(id int64) (models.Post, error)
}

type PostService interface {
	CreatePost(post models.Post) (int64, error)
	GetAllPosts() ([]models.Post, error)
	GetPostsByID(id int64) (models.Post, error)
}
