package services

import "fmt"

func (s *PostService) CreateComment(userID string, postID string, content string, image_url string) (string, string, error) {
	fmt.Println("entred service comment")
	return s.repo.CreateComment(userID, postID, content, image_url)
}
