package notification

import (
	"social-network/backend/models"
)

func (NS *NotificationService) UpdateService(data models.Unotif, user_id string) *models.ErrorJson {
	notification, errJson := NS.notifRepo.SelectNotificationById(data.NotifId)
	if errJson != nil {
		return &models.ErrorJson{Status: 400, Error: "ERROR 400 - Bad Request", Message: errJson.Message}
	}

	if user_id != notification.RecieverId {
		return &models.ErrorJson{Status: 403, Error: "ERROR 403 Acces Forbidden", Message: "Invalid Operation"}
	}

	if notification.Status != "later" {
		return models.NewErrorJson(400, "400 - Bad Request", "Invalid Operation")
	}

	switch data.Type {
	case "follow-private":
		errJson = NS.UpdateFollowPrivateProfile(data, notification)
	case "group-invitation":
		errJson = NS.UpdateGroupInvitationRequest(data, notification)
	case "group-join":
		errJson = NS.UpdateGroupJoinRequest(data, notification)
	default:
		return models.NewErrorJson(400, "400 - Bad Request", "invalid type")
	}
	if errJson != nil {
		return errJson
	}

	if errJson = NS.notifRepo.UpdateStatus(notification.Id, data.Status); errJson != nil {
		return errJson
	}
	return nil
}

func (NS *NotificationService) UpdateFollowPrivateProfile(data models.Unotif, notification models.Notification) *models.ErrorJson {
	switch data.Status {
	case "accept":
		if err := NS.profileRepo.AcceptedRequest(notification.RecieverId, notification.SenderId); err != nil {
			return models.NewErrorJson(500, "500 - cannot accept request", err)
		}
	case "reject":
		if err := NS.profileRepo.RejectedRequest(notification.RecieverId, notification.SenderId); err != nil {
			return models.NewErrorJson(500, "500 - cannot reject request", err)
		}
	default:
		return models.NewErrorJson(400, "400 - Bad Request", "Invalid Status")
	}
	return nil
}

func (NS *NotificationService) UpdateGroupJoinRequest(data models.Unotif, notification models.Notification) *models.ErrorJson {
	switch data.Status {
	case "accept":
		if err := NS.groupRepo.Approve(notification.GroupId, notification.SenderId); err != nil {
			return models.NewErrorJson(500, "500 - cannot accept request", err)
		}

		// should remove invitation request if exist
		if err := NS.groupRepo.CancelTheInvitation(notification.RecieverId, notification.GroupId, notification.SenderId); err != nil {
			return models.NewErrorJson(500, "500 - cannot cancel invitation request after accept join request", err)
		}

		if err := NS.DeleteService(notification.SenderId, notification.RecieverId, "group-invitation", notification.GroupId); err != nil {
			return models.NewErrorJson(500, "500 - cannot delete notification join after accept join request", err)
		}

	case "reject":
		if err := NS.groupRepo.Decline(notification.GroupId, notification.SenderId); err != nil {
			return models.NewErrorJson(500, "500 - cannot decline request", err)
		}
	default:
		return models.NewErrorJson(400, "Bad-Request 400", "Invalid Status")
	}
	return nil
}

func (NS *NotificationService) UpdateGroupInvitationRequest(data models.Unotif, notification models.Notification) *models.ErrorJson {
	switch data.Status {
	case "accept":
		if err := NS.groupRepo.Accept(notification.SenderId, notification.GroupId, notification.RecieverId); err != nil {
			return models.NewErrorJson(500, "500 - cannot accept request", err)
		}

		// should remove join request if exist
		if err := NS.groupRepo.RequestToCancel(notification.RecieverId, notification.GroupId); err != nil {
			return models.NewErrorJson(500, "500 - cannot cancel join request after accept join request", err)
		}

		if err := NS.DeleteService(notification.SenderId, notification.RecieverId, "group-join", notification.GroupId); err != nil {
			return models.NewErrorJson(500, "500 - cannot delete notification join after accept join request", err)
		}

	case "reject":
		if err := NS.groupRepo.Reject(notification.SenderId, notification.GroupId, notification.RecieverId); err != nil {
			return models.NewErrorJson(500, "500 - cannot accept request", err)
		}
	default:
		return models.NewErrorJson(400, "Bad-Request", "Invalid Status")
	}
	return nil
}
