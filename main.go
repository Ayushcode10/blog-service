package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"github.com/Ayushcode10/blog-service/handler"
	"github.com/Ayushcode10/blog-service/service"
	"github.com/Ayushcode10/blog-service/store"
)

func main() {
	// Connect to MySQL
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3307)/blogdb")
	if err != nil {
		log.Fatal("DB Open Error:", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("DB Ping Error:", err)
	}
	fmt.Println("✅ Connected to MySQL")

	// Init Store → Service → Handler
	postStore := store.NewMySQLStore(db)
	postService := service.NewPostService(postStore)
	postHandler := handler.NewPostHandler(postService)

	// Setup router
	r := mux.NewRouter()
	r.HandleFunc("/posts", postHandler.CreatePost).Methods("POST")
	r.HandleFunc("/posts", postHandler.GetAllPosts).Methods("GET")
	r.HandleFunc("/posts/{id}", postHandler.GetPostByID).Methods("GET")

	fmt.Println("🚀 Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
