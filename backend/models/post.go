package models

type Post struct {
	ID            string `json:"id"`
	UserID        string `json:"user_id"`
	GroupID       string `json:"group_id"`
	Content       string `json:"content"`
	CreatedAt     string `json:"created_at"`
	LikesCount    int    `json:"likes_count"`
	CommentsCount int    `json:"comments_count"`
	LikedByUser   bool   `json:"is_liked"`
	Privacy       string `json:"privacy"`
}

func NewPost() *Post {
	return &Post{}
}

type Reaction struct {
	Id           int    `json:"id,omitempty"`
	EntityTypeId int    `json:"entity_type_id,omitempty"`
	EntityType   string `json:"entity_type,omitempty"`
	EntityId     int    `json:"entity_id,omitempty"`
	Reaction     int    `json:"reaction"`
	UserId       int    `json:"user_id,omitempty"`
}

func NewReaction() *Reaction {
	return &Reaction{}
}
