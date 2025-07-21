package group

import (
	"strings"

	"social-network/backend/models"
	"social-network/backend/repositories/group"
	"social-network/backend/services/notification"
	"social-network/backend/services/profile"
	"social-network/backend/utils"
)

type GroupService struct {
	gRepo    *group.GroupRepository
	sProfile *profile.ProfileService
	sNotif   *notification.NotificationService
}

func NewGroupService(grepo *group.GroupRepository, sProfile *profile.ProfileService, sNotif *notification.NotificationService) *GroupService {
	return &GroupService{gRepo: grepo, sProfile: sProfile, sNotif: sNotif}
}

func (gService *GroupService) AddGroup(group *models.Group) (*models.Group, *models.ErrorJson) {
	errGroup := models.ErrGroup{}
	trimmedTitle := strings.TrimSpace(group.Title)
	trimmedDesc := strings.TrimSpace(group.Description)
	if err := gService.ValidateGroupTitle(trimmedTitle); err != nil {
		errGroup.Title = err.Error()
	}
	if err := utils.ValidateDesc(trimmedDesc); err != nil {
		errGroup.Description = err.Error()
	}

	if errGroup != (models.ErrGroup{}) {
		if group.ImagePath != "" {
			if err := utils.RemoveImage(group.ImagePath); err != nil {
				return nil, &models.ErrorJson{Status: 500, Error: err.Error()}
			}
		}
		return nil, &models.ErrorJson{Status: 400, Message: errGroup}
	}
	group.Title, group.Description = trimmedTitle, trimmedDesc
	groupCreated, errJson := gService.gRepo.CreateGroup(group)
	if errJson != nil {
		if group.ImagePath != "" {
			if err := utils.RemoveImage(group.ImagePath); err != nil {
				return nil, &models.ErrorJson{Status: 500, Error: err.Error()}
			}
		}
		return nil, &models.ErrorJson{Status: errJson.Status, Message: errJson.Message, Error: errJson.Error}
	}
	return groupCreated, nil
}

func (gService *GroupService) GetGroups(filter string, offset int64, userID string) ([]models.Group, *models.ErrorJson) {
	var groups []models.Group
	var err *models.ErrorJson
	switch filter {
	case "owned":
		groups, err = gService.gRepo.GetCreatedGroups(offset, userID)
	case "available":
		groups, err = gService.gRepo.GetAvailableGroups(offset, userID)
	case "joined":
		groups, err = gService.gRepo.GetJoinedGroups(offset, userID)

	}

	if err != nil {
		return nil, &models.ErrorJson{Status: err.Status, Error: err.Error, Message: err.Message}
	}

	return groups, nil
}
