package models

type Followers struct {
	UserID     string `db:"userID" json:"followed"`
	FollowerID string `db:"followerID" json:"follower"`
}
