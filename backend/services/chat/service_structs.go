package chat

import (
	"social-network/backend/repositories/chat"
)

type ChatService struct {
	repo *chat.ChatRepository
}

func NewChatService(repo *chat.ChatRepository) *ChatService {
	return &ChatService{repo: repo}
}
