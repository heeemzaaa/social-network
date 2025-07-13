package group

import "social-network/backend/models"

func (service *GroupService) GetEventDetails(eventId, userId, groupId string) (*models.Event, *models.ErrorJson) {
	return service.gRepo.GetEventDetails(eventId, userId, groupId)
}
