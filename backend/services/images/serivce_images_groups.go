package images

import "social-network/backend/models"


func (s *ServiceImages) AccessToGroupImage(groupID, userID string) (bool, *models.ErrorJson) {
	isMember, err := s.groupService.IsMemberGroup(groupID, userID)
	if err != nil {
		return false, &models.ErrorJson{Status: err.Status, Error: err.Error}
	}
	return  isMember, nil
}