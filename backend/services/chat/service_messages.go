package chat

import (
	"fmt"
	"strings"

	"social-network/backend/models"
)

func (service *ChatService) ValidateMessage(message *models.Message) (*models.Message, *models.ErrorJson) {
	errMessage := models.NewMessageErr()
	trimmedMsg := strings.TrimSpace(message.Content)
	type_message := strings.ToLower(strings.TrimSpace(message.Type))

	if type_message != "private" && type_message != "group" && type_message != "notification" {
		errMessage.Type = "wrong type of message"
	}

	if type_message != "notification" {
		if trimmedMsg == "" {
			errMessage.Content = "empty message Body"
		}
		if len(trimmedMsg) > 1000 {
			errMessage.Content = "message body too large!"
		}
	}

	if errMessage.Content != "" || errMessage.Type != "" {
		return nil, &models.ErrorJson{Status: 400, Message: errMessage}
	}




	// We can go on and insert the message in the database
	switch strings.ToLower(message.Type) {
	case "private":
		// hna I will need to check if the userID match with some of my users
		exists, err := service.repo.UserExists(message.TargetID)
		if err != nil {
			return nil, &models.ErrorJson{Status: err.Status, Error: err.Error}
		}

		if !exists {
			return nil, &models.ErrorJson{Status: 400, Error: "The user doesn't exist"}
		}

		message_created, err := service.repo.AddMessage(message)
		if err != nil {
			return nil, &models.ErrorJson{Status: err.Status, Error: err.Error}
		}
		return message_created, nil
	case "group":

		exists, err := service.repo.GroupExists(message.TargetID)
		if err != nil {
			return nil, &models.ErrorJson{Status: err.Status, Error: err.Error}
		}

		if !exists {
			return nil, &models.ErrorJson{Status: 400, Error: "The group doesn't exist"}
		}

		isMember, err := service.repo.IsMember(message.SenderID, message.TargetID)
		if err != nil {
			return nil, &models.ErrorJson{Status: err.Status, Error: err.Error}
		}

		if !isMember {
			return nil, &models.ErrorJson{Status: 403, Error: "you're not a member in this group !"}
		}

		message_created, err := service.repo.AddMessage(message)
		if err != nil {
			return nil, &models.ErrorJson{Status: err.Status, Error: err.Error}
		}
		return message_created, nil
	case "notification":
		if message.Notification == (models.Notification{}) {
			return nil, &models.ErrorJson{Status: 400, Error: "The notification is empty"}
		}
		if strings.TrimSpace(message.Notification.Content) == "" {
			return nil, &models.ErrorJson{Status: 400, Message: models.Notification{Content: "empty content!"}}
		}

		if strings.HasPrefix(message.Notification.Type, "group") {
			fmt.Printf("%#v", message.Notification.Type)
			exists, err := service.repo.GroupExists(message.Notification.GroupID)
			if err != nil {
				return nil, &models.ErrorJson{Status: err.Status, Error: err.Error}
			}

			if !exists {
				return nil, &models.ErrorJson{Status: 400, Error: "The group doesn't exist"}
			}
		} else {
			exists, err := service.repo.UserExists(message.Notification.TargetID)
			if err != nil {
				return nil, &models.ErrorJson{Status: err.Status, Error: err.Error}
			}

			if !exists {
				return nil, &models.ErrorJson{Status: 400, Error: "The user doesn't exist"}
			}
		}

		message.Notification.SenderID = message.SenderID

		if errJson := service.PostService(message.Notification); errJson != nil {
			return nil, &models.ErrorJson{Status: 400, Error: errJson.Error}
		}
		return message, nil

	}
	return nil, nil
}

// get the messages between specific users
func (service *ChatService) GetMessages(sender_id, target_id, type_ string) ([]models.Message, *models.ErrorJson) {
	if sender_id == "" || target_id == "" || type_ == "" {
		return []models.Message{}, &models.ErrorJson{Status: 400, Error: "Invalid data format !"}
	}

	if sender_id == target_id {
		return []models.Message{}, &models.ErrorJson{Status: 400, Error: "Invalid target ID !"}
	}

	messages, errJson := service.repo.GetMessages(sender_id, target_id, type_)
	if errJson != nil {
		return nil, &models.ErrorJson{Status: errJson.Status, Error: errJson.Error}
	}
	return messages, nil
}

func (service *ChatService) EditReadStatus(sender_id, target_id string) *models.ErrorJson {
	if err := service.repo.EditReadStatus(sender_id, target_id); err != nil {
		return &models.ErrorJson{Status: err.Status, Error: err.Error}
	}
	return nil
}
