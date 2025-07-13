package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"social-network/backend/models"
)

func NewPostHandler(w http.ResponseWriter, r *http.Request) {
	
	fmt.Println("inside the get posts handler.")

	// fake post
	posts := []models.Post{
		{
			ID:            "1",
			UserID:        "u123",
			GroupID:       "g1",
			Content:       "This is my first post!",
			CreatedAt:     "2025-06-30T10:00:00Z",
			LikesCount:    15,
			CommentsCount: 3,
			LikedByUser:   true,
			Privacy:       "public",
		},
		{
			ID:            "2",
			UserID:        "u124",
			GroupID:       "g2",
			Content:       "Enjoying the summer ☀️",
			CreatedAt:     "2025-06-29T14:30:00Z",
			LikesCount:    24,
			CommentsCount: 6,
			LikedByUser:   false,
		},
		{
			ID:            "3",
			UserID:        "u125",
			GroupID:       "g3",
			Content:       "Working on my Go project 👨‍💻",
			CreatedAt:     "2025-06-28T09:15:00Z",
			LikesCount:    42,
			CommentsCount: 10,
			LikedByUser:   true,
		},
	}

	json.NewEncoder(w).Encode(posts)
}
