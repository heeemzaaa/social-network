package chat

import (

	"social-network/backend/models"
)

// here I will get the userIds of the members of a specefic group
func (service *ChatService) GetMembersOfGroup(groupID string) ([]string, *models.ErrorJson) {
	members, err := service.repo.GetMembersOfGroup(groupID)
	if err != nil {
		return members, &models.ErrorJson{Status: err.Status, Error: err.Error}
	}
	return members, nil
}

// here I will get the users that I have connection with
func (service *ChatService) GetUsers(authUserID string) ([]models.User, *models.ErrorJson) {
	users, err := service.repo.GetUsers(authUserID)
	if err != nil {
		return users, &models.ErrorJson{Status: err.Status, Error: err.Error}
	}

	return users, nil
}

// here I will get the groups that I'm member in
func (service *ChatService) GetGroups(authUserID string) ([]models.Group, *models.ErrorJson) {
	groups, err := service.repo.GetGroups(authUserID)
	if err != nil {
		return []models.Group{}, &models.ErrorJson{Status: err.Status, Error: err.Error}
	}
	return groups, nil
}
