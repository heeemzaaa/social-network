package chat

import (
	"social-network/backend/models"
	"social-network/backend/repositories/chat"
	"strings"
	"time"
)

type ChatService struct {
	repo *chat.ChatRepository
}

func NewChatService(repo *chat.ChatRepository) *ChatService {
	return &ChatService{repo: repo}
}

func (s *ChatService) GetSessionByTokenEnsureAuth(token string) (*models.Session, *models.ErrorJson) {
	session, err := s.repo.GetSessionbyTokenEnsureAuth(token)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (service *ChatService) ValidateMessage(message *models.Message) (*models.Message, *models.ErrorJson) {
	errMessage := models.NewMessageErr()
	trimmedMsg := strings.TrimSpace(message.Content)
	type_message := strings.ToLower(strings.TrimSpace(message.Type))

	if type_message != "message" && type_message != "group" {
		errMessage.Type = "wrong type of message"
	}

	if trimmedMsg == "" {
		errMessage.Content = "empty message Body"
	}
	if len(trimmedMsg) > 1000 {
		errMessage.Content = "message body too large!"
	}
	//hna I will need to check if the userID match with some of my users
	firstName, lastName, err := service.repo.GetFullNameById(message.TargetID)
	if firstName == "" && lastName == "" {
		errMessage.TargetID = "The receiver specified does dot exist!!"
		return nil, err
	}

	
	if errMessage.Content != "" || errMessage.TargetID != "" || errMessage.Type != "" {
		return nil, &models.ErrorJson{Status: 400, Message: errMessage}
	}
	
	message.CreatedAt = time.Now().String()
	
	// We can go on and insert the message in the database
	switch strings.ToLower(message.Type) {
	case "message":
		message_created, err := service.repo.AddMessage(message)
		if err != nil {
			return nil, err
		}
		message_created.Type = type_message
		return message_created, nil
	case "group":

	}
	return nil, nil
}
