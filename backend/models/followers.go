package models

type Followers struct {
	UserID string `json:"followed"`
	FollowerID string `json:"follower"`
}