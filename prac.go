package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	UserID int    `json:"user_id"`
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Endpoint to fetch user by ID
	r.Get("/users/{userId}", func(w http.ResponseWriter, r *http.Request) {
		userID := chi.URLParam(r, "userId")

		// Simulate fetching user data from a database (replace with actual logic)
		user := User{
			ID:   1,
			Name: "John Doe",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	})

	// Endpoint to fetch posts (assuming filtering by user ID is supported)
	r.Get("/posts", func(w http.ResponseWriter, r *http.Request) {
		userID := r.URL.Query().Get("userId")

		// Simulate fetching posts from a database (replace with actual logic)
		posts := []Post{
			{ID: 1, Title: "Post 1", UserID: 1},
			{ID: 2, Title: "Post 2", UserID: 2},
		}

		var userPosts []Post
		for _, post := range posts {
			if post.UserID == parseUserID(userID) { // Assuming helper function to parse user ID
				userPosts = append(userPosts, post)
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(userPosts)
	})

	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", r)
}

func parseUserID(userID string) (int, error) {
	// Implement logic to parse user ID from string to integer
	// Handle potential errors during parsing
}
