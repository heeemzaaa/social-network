package chat

import (
	"social-network/backend/models"
)

// check if the user exist or not , to procced after it found
func (service *ChatService) UserExists(targetID string) (bool, *models.ErrorJson) {
	exists, err := service.repo.UserExists(targetID)
	if err != nil {
		return false, &models.ErrorJson{Status: err.Status, Error: err.Error}
	}
	return exists, nil
}

// check if the group exist or not , to procced
func (service *ChatService) GroupExists(targetID string) (bool, *models.ErrorJson) {
	exists, err := service.repo.GroupExists(targetID)
	if err != nil {
		return false, &models.ErrorJson{Status: err.Status, Error: err.Error}
	}
	return exists, nil
}

// check the existance of the target whether its a groupID or userID
func (service *ChatService) CheckExistance(type_, target_id string) (bool, *models.ErrorJson) {
	switch type_ {
	case "private":
		exists, errJson := service.UserExists(target_id)
		if errJson != nil {
			return false, &models.ErrorJson{Status: errJson.Status, Error: errJson.Error}
		}
		// check if the user exists
		if !exists {
			return false, &models.ErrorJson{Status: 400, Error: "The userId doesn't exist"}
		}
		return exists, nil
	case "group":
		exists, errJson := service.GroupExists(target_id)
		if errJson != nil {
			return false, &models.ErrorJson{Status: errJson.Status, Error: errJson.Error}
		}

		// check if the groupID exists
		if !exists {
			return false, &models.ErrorJson{Status: 400, Error: "The userId doesn't exist"}
		}
		return exists, nil
	}
	return false, &models.ErrorJson{Status: 400, Error: "the type is not correct"}
}
