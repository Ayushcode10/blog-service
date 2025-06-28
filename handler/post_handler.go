package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Ayushcode10/blog-service/models"
	"github.com/Ayushcode10/blog-service/service"
	"github.com/gorilla/mux"
)

type PostHandler struct {
	postService *service.PostService
}

func NewPostHandler(s *service.PostService) *PostHandler {
	return &PostHandler{postService: s}
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if err := h.postService.CreatePost(post); err != nil {
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Post created!"})
}

func (h *PostHandler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := h.postService.GetAllPosts()
	if err != nil {
		http.Error(w, "Error fetching posts", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(posts)
}

func (h *PostHandler) GetPostByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	post, err := h.postService.GetPostByID(id)
	if err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(post)
}
