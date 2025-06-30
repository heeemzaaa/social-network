package models

type Profile struct {
	User              User `json:"user,omitempty"`
	NumberOfPosts     int  `json:"posts_count,omitempty"`
	NumberOfFollowing int  `json:"following_count,omitempty"`
	NumberOfFollowers int  `json:"followers_count,omitempty"`
	NumberOfGroups    int  `json:"groups_count,omitempty"`
	IsMyProfile       bool `json:"is_my_profile,omitempty"`
	IsFollower        bool `json:"is_follower,omitempty"`
	IsPrivate         bool `json:"is_private,omitempty"`
}
