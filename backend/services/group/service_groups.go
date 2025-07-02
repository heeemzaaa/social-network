package group

import (
	"strings"

	"social-network/backend/models"
	"social-network/backend/repositories/group"
	"social-network/backend/utils"
)

type GroupService struct {
	grepo *group.GroupRepository
}

func NewGroupService(grepo *group.GroupRepository) *GroupService {
	return &GroupService{grepo: grepo}
}

func (gServicde *GroupService) AddGroup(group *models.Group) (*models.Group, *models.ErrorJson) {
	errGroup := models.ErrGroup{}
	trimmedTitle := strings.TrimSpace(group.Title)
	trimmedDesc := strings.TrimSpace(group.Description)
	if trimmedTitle == "" {
		errGroup.Title = "empty title field"
	}
	if trimmedDesc == "" {
		errGroup.Title = "empty description field"
	}
	if err := utils.ValidateTitle(trimmedTitle); err != nil {
		errGroup.Title = err.Error()
	}
	if err := utils.ValidateDesc(trimmedDesc); err != nil {
		errGroup.Description = err.Error()
	}

	if errGroup != (models.ErrGroup{}) {
		return nil, &models.ErrorJson{Status: 400, Message: errGroup}
	}
	group, errJson := gServicde.grepo.CreateGroup(group)
	if errJson != nil {
		return nil, &models.ErrorJson{Status: errJson.Status, Message: errJson.Message}
	}
	return group, nil
}

func (gService *GroupService) GetGroups(filter string, offset int64, userID string) ([]models.Group, *models.ErrorJson) {
	var groups []models.Group
	var err *models.ErrorJson
	switch filter {
	case "owned":
		groups, err = gService.grepo.GetCreatedGroups(offset, userID)
	case "available":
		groups, err = gService.grepo.GetAvailableGroups(offset, userID)
	case "joined":
		groups, err = gService.grepo.GetJoinedGroups(offset, userID)

	}

	if err != nil {
		return nil, err
	}

	return groups, nil
}
