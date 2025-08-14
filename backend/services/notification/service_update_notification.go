package notification

import (
	"social-network/backend/models"
)

// UpdateService updates a notification based on the provided data and user ID.
func (NS *NotificationService) UpdateService(data models.Unotif, userId string) *models.ErrorJson {
	notification, errJson := NS.notifRepo.SelectNotificationById(data.NotifId)
	if errJson != nil {
		return &models.ErrorJson{Status: errJson.Status, Error: errJson.Error, Message: errJson.Message}
	}

	if userId != notification.TargetID {
		return &models.ErrorJson{Status: 403, Error: "Invalid Operation"}
	}

	if notification.Status != "later" {
		return &models.ErrorJson{Status: 400, Error: "400 - Bad Request"}
	}

	switch data.Type {
	case "follow-private":
		errJson = NS.UpdateFollowPrivateProfile(data, notification)
	case "group-invitation":
		errJson = NS.UpdateGroupInvitationRequest(data, notification)
	case "group-join":
		errJson = NS.UpdateGroupJoinRequest(data, notification)
	default:
		return &models.ErrorJson{
			Status: 400,
			Error:  "invalid type",
		}
	}
	if errJson != nil {
		return errJson
	}

	if errJson = NS.notifRepo.UpdateStatus(notification.Id, data.Status); errJson != nil {
		return errJson
	}
	return nil
}

// UpdateFollowPrivateProfile updates the follow request for a private profile based on the provided data.
func (NS *NotificationService) UpdateFollowPrivateProfile(data models.Unotif, notification *models.Notification) *models.ErrorJson {
	switch data.Status {
	case "accept":
		if err := NS.profileService.AcceptedRequest(notification.TargetID, notification.SenderID); err != nil {
			return &models.ErrorJson{Status: 500, Error: "500 - cannot reject request"}
		}

	case "reject":
		if err := NS.profileService.RejectedRequest(notification.TargetID, notification.SenderID); err != nil {
			return &models.ErrorJson{Status: 500, Error: "500 - cannot reject request"}
		}

	default:
		return &models.ErrorJson{Status: 400, Error: "Invalid Status"}

	}
	return nil
}

// UpdateGroupJoinRequest updates the group join request based on the provided data.
func (NS *NotificationService) UpdateGroupJoinRequest(data models.Unotif, notification *models.Notification) *models.ErrorJson {
	// get groupId

	// groupId, errJson := NS.groupService.GetGroupIdByCreatorId(notification.TargetID)
	// if errJson != nil {
	// 	return models.NewErrorJson(500, "400 - GROUP not found", errJson.Message)
	// }
	// fmt.Println("groupID ====> ",groupId)

	switch data.Status {
	case "accept":
		if err := NS.groupService.Approve(notification.TargetID, notification.GroupID, notification.SenderID); err != nil {
			return models.NewErrorJson(500, "500 - cannot accept request", err)
		}

		// check if request exists
		if err := NS.groupService.CancelTheInvitation(notification.TargetID, notification.GroupID, notification.SenderID); err != nil {
			if err.Error == "Invitation not found" {
				if err := NS.DeleteService(notification.SenderID, notification.TargetID, "group-invitation", notification.GroupID); err != nil {
					return &models.ErrorJson{Status: 500, Error: "500 - cannot delete notification join after accept join request"}
				}
				return nil
			}
			return &models.ErrorJson{Status: 500, Error: "cannot cancel invitation request after accept join request"}

		}

		if err := NS.DeleteService(notification.SenderID, notification.TargetID, "group-invitation", notification.GroupID); err != nil {
			return &models.ErrorJson{Status: 500, Error: "500 - cannot delete notification join after accept join request"}
		}

	case "reject":
		if err := NS.groupService.Decline(notification.TargetID, notification.GroupID, notification.SenderID); err != nil {
			return &models.ErrorJson{Status: 500, Error: "500 - cannot decline request"}
		}

	default:
		return &models.ErrorJson{Status: 400, Error: "Invalid Status"}
	}
	return nil
}

// UpdateGroupInvitationRequest updates the group invitation request based on the provided data.
func (NS *NotificationService) UpdateGroupInvitationRequest(data models.Unotif, notification *models.Notification) *models.ErrorJson {
	switch data.Status {
	case "accept":
		if err := NS.groupService.Accept(notification.SenderID, notification.GroupID, notification.TargetID); err != nil {
			return models.NewErrorJson(500, "500 - cannot accept request", err)
		}

		// check if request exists
		_, err := NS.groupService.RequestToCancel(notification.TargetID, notification.GroupID)
		if err != nil {

			// if the error is because the user has not requested to join or is already a member case when has notification
			if err.Error == "Invitation not found" || err.Error == "Already a member!" {
				if err := NS.DeleteService(notification.SenderID, notification.TargetID, "group-join", notification.GroupID); err != nil {
					return &models.ErrorJson{Status: 500, Error: "500 - cannot delete notification join after accept join request"}
				}
				return nil
			}
			return &models.ErrorJson{Status: 500, Error: "500 - cannot cancel join request after accept join request"}

		}

		if err := NS.DeleteService(notification.SenderID, notification.TargetID, "group-join", notification.GroupID); err != nil {
			return &models.ErrorJson{Status: 500, Error: "500 - cannot delete notification join after accept join request"}
		}

	case "reject":
		if err := NS.groupService.Reject(notification.SenderID, notification.GroupID, notification.TargetID); err != nil {
			return &models.ErrorJson{Status: 500, Error: "500 - cannot accept request"}
		}

	default:
		return &models.ErrorJson{Status: 400, Error: "Invalid Status"}

	}
	return nil
}
