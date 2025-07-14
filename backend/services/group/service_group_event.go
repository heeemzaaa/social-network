package group

import "social-network/backend/models"

func (service *GroupService) GetEventDetails(eventId, userId, groupId string) (*models.Event, *models.ErrorJson) {
	event, errJson := service.gRepo.GetEventDetails(eventId, userId, groupId)
	if errJson != nil {
		return nil, &models.ErrorJson{Status: errJson.Status, Error: errJson.Error, Message: errJson.Message}
	}
	return event, nil
}
