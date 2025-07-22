package models

type Post struct {
	Id            string   `json:"id,omitempty"`
	User          User     `json:"user"`
	Content       string   `json:"content,omitempty"`
	CreatedAt     string   `json:"created_at,omitempty"`
	Img           string   `json:"image_path,omitempty"`
	TotalLikes    int      `json:"total_likes,omitempty"`
	TotalComments int      `json:"total_comments"`
	Privacy       string   `json:"privacy"`
	Liked         int      `json:"liked,omitempty"`
	SelectedUsers []string `json:"selected_users,omitempty"`
}

func NewPost() *Post {
	return &Post{}
}

type Reaction struct {
	Id         string `json:"id,omitempty"`
	EntityType string `json:"entity_type,omitempty"`
	EntityId   string `json:"entity_id,omitempty"`
	Reaction   int    `json:"reaction"`
	UserId     string `json:"user_id,omitempty"`
}

func NewReaction() *Reaction {
	return &Reaction{}
}
