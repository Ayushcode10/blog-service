package service

import (
	"github.com/Ayushcode10/blog-service/models"
	"github.com/Ayushcode10/blog-service/store"
)

type PostService struct {
	store store.Store
}

func NewPostService(s store.Store) *PostService {
	return &PostService{store: s}
}

func (ps *PostService) CreatePost(post models.Post) error {
	// You can add validation here (e.g., check if title is empty)
	return ps.store.Create(post)
}

func (ps *PostService) GetAllPosts() ([]models.Post, error) {
	return ps.store.GetAll()
}

func (ps *PostService) GetPostByID(id int) (models.Post, error) {
	return ps.store.GetByID(id)
}
