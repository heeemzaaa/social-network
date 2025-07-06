package chat

import (
	"social-network/backend/models"
	"social-network/backend/repositories/chat"
	"strings"
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

	if type_message != "message" && type_message != "read" {
		errMessage.Type = "wrong type of message"
	}

	if trimmedMsg == "" {
		errMessage.Content = "empty message Body"
	}
	if len(trimmedMsg) > 1000 {
		errMessage.Content = "message body too large!"
	}
	// hna I will need to check if the userID match with some of my users
	// if username, _ := service.repo.GetUserNameById(message.TargetID); username == "" {
	// 	errMessage.ReceiverID = "The receiver specified does dot exist!!"
	// }

	// if message.CreatedAt.IsZero() {
	// 	errMessage.CreatedAt = "the date is not set up!"
	// }

	if errMessage.Content != "" || errMessage.ReceiverID != "" || errMessage.Type != "" || errMessage.CreatedAt != "" {
		return nil, &models.ErrorJson{Status: 400, Message: errMessage}
	}

	// We can go on and insert the message in the database
	switch strings.ToLower(message.Type) {
	case "message":
		message_created, err := service.repo.AddMessage(message)
		if err != nil {
			return nil, err
		}
		message_created.Type = type_message
		return message_created, nil
	case "read":
		service.EditReadStatus(message.SenderID, message.TargetID)
		return message, nil
	case "typing":

	}

	// so in this case we only need to update the database (the message exists already)
	// if message.Type == "read" {
	// }
	return nil, nil
}

// from the unread to the read status
func (service *ChatService) EditReadStatus(sender_id, target_id string) *models.ErrorJson {
	if err := service.repo.EditReadStatus(sender_id, target_id); err != nil {
		return err
	}
	return nil
}
