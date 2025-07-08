package chat

import (
	"fmt"
	"net/http"
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

	if type_message != "private" && type_message != "group" {
		errMessage.Type = "wrong type of message"
	}

	if trimmedMsg == "" {
		errMessage.Content = "empty message Body"
	}
	if len(trimmedMsg) > 1000 {
		errMessage.Content = "message body too large!"
	}

	if errMessage.Content != "" || errMessage.TargetID != "" || errMessage.Type != "" {
		return nil, &models.ErrorJson{Status: 400, Message: errMessage}
	}

	message.CreatedAt = time.Now().UTC().Format(time.RFC3339)

	// We can go on and insert the message in the database
	switch strings.ToLower(message.Type) {
	case "private":
		//hna I will need to check if the userID match with some of my users
		exists, err := service.repo.UserExists(message.TargetID)
		if err != nil {
			return nil, err
		}

		if !exists {
			return nil, &models.ErrorJson{Status: 400, Error: "", Message: "The user doesn't exist"}
		}

		message_created, err := service.repo.AddMessage(message)
		if err != nil {
			return nil, err
		}
		message_created.Type = type_message
		return message_created, nil
	case "group":

		exists, err := service.repo.GroupExists(message.TargetID)
		if err != nil {
			return nil, err
		}

		if !exists {
			return nil, &models.ErrorJson{Status: 400, Error: "", Message: "The group doesn't exist"}
		}

		isMember, err := service.repo.IsMember(message.SenderID, message.TargetID)
		if err != nil {
			return nil, err
		}

		if !isMember {
			return nil, &models.ErrorJson{Status: 403, Error: "", Message: "you're not a member in this group !"}
		}

		message_created, err := service.repo.AddMessage(message)
		if err != nil {
			return nil, err
		}
		return message_created, nil
	}
	return nil, nil
}

// get the messages between specific users
func (service *ChatService) GetMessages(sender_id, target_id string, offset int, type_ string) ([]models.Message, *models.ErrorJson) {
	messages, errJson := service.repo.GetMessages(sender_id, target_id, offset, type_)
	if errJson != nil {
		return nil, errJson
	}
	return messages, nil
}

// check if the user exist or not , to procced after it found
func (service *ChatService) UserExists(targetID string) (bool, *models.ErrorJson) {
	exists, err := service.repo.UserExists(targetID)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (service *ChatService) GroupExists(targetID string) (bool, *models.ErrorJson) {
	exists, err := service.repo.GroupExists(targetID)
	if err != nil {
		return false, err
	}
	return exists, nil
}

// here I will get the userIds of the members of a specefic group
func (service *ChatService) GetMembersOfGroup(groupID string) ([]string, *models.ErrorJson) {
	members, err := service.repo.GetMembersOfGroup(groupID)
	if err != nil {
		return nil, err
	}
	return members, nil
}

// get the userID from the session after pass it to the repo
func (service *ChatService) GetUserIdFromSession(r *http.Request) (string, *models.ErrorJson) {
	cookie, err := r.Cookie("session")
	if err != nil {
		return "", &models.ErrorJson{Status: 401, Error: "", Message: fmt.Sprintf("%v", err)}
	}

	userID, errQuery := service.repo.GetUserIdFromSession(cookie.Value)
	if errQuery != nil {
		return "", errQuery
	}

	return userID, nil
}

func (service *ChatService) CheckExistance(type_, target_id string) (bool, *models.ErrorJson) {
	switch type_ {
	case "private":
		exists, errJson := service.UserExists(target_id)
		if errJson != nil {
			return false, &models.ErrorJson{Status: 400, Error: "", Message: fmt.Sprintf("%v", errJson)}
		}
		// check if the user exists
		if !exists {
			return false, &models.ErrorJson{Status: 400, Error: "", Message: "The userId format is not valid"}
		}
		return exists, nil
	case "group":
		exists, errJson := service.GroupExists(target_id)
		if errJson != nil {
			return false, &models.ErrorJson{Status: 400, Error: "", Message: fmt.Sprintf("%v", errJson)}
		}

		// check if the groupID exists
		if !exists {
			return false, &models.ErrorJson{Status: 400, Error: "", Message: "The groupId format is not valid"}
		}
		return exists, nil
	}
	return false, &models.ErrorJson{Status: 400, Error: "", Message: "the type is not correct"}
}
