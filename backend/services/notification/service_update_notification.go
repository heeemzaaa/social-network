package notification

import (
	"fmt"
	"social-network/backend/models"
)

// UpdateService updates a notification based on the provided data and user ID.
func (NS *NotificationService) UpdateService(data models.Unotif, userId string) *models.ErrorJson {
	notification, errJson := NS.notifRepo.SelectNotificationById(data.NotifId)
	if errJson != nil {
		if errJson.Status == 500 {
			return &models.ErrorJson{Status: 500, Message: "notification not found", Error: "500 - Internal Server Error"}
		}
		return &models.ErrorJson{Status: 400, Error: "ERROR 400 - Bad Request", Message: errJson.Message}
	}

	if userId != notification.RecieverId {
		return &models.ErrorJson{Status: 403, Error: "ERROR 403 Acces Forbidden", Message: "Invalid Operation"}
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

// UpdateFollowPrivateProfile updates the follow request for a private profile based on the provided data.
func (NS *NotificationService) UpdateFollowPrivateProfile(data models.Unotif, notification models.Notification) *models.ErrorJson {
	switch data.Status {
	case "accept":
		if err := NS.profileService.AcceptedRequest(notification.RecieverId, notification.SenderId); err != nil {
			return models.NewErrorJson(500, "500 - cannot accept request", err)
		}

	case "reject":
		if err := NS.profileService.RejectedRequest(notification.RecieverId, notification.SenderId); err != nil {
			return models.NewErrorJson(500, "500 - cannot reject request", err)
		}

	default:
		return models.NewErrorJson(400, "400 - Bad Request", "Invalid Status")
	}
	return nil
}

// UpdateGroupJoinRequest updates the group join request based on the provided data.
func (NS *NotificationService) UpdateGroupJoinRequest(data models.Unotif, notification models.Notification) *models.ErrorJson {
	switch data.Status {
	case "accept":
		if err := NS.groupService.Approve(notification.RecieverId, notification.GroupId, notification.SenderId); err != nil {
			return models.NewErrorJson(500, "500 - cannot accept request", err)
		}

		// should check if request exists before canceling //////////////////

		if err := NS.groupService.CancelTheInvitation(notification.RecieverId, notification.GroupId, notification.SenderId); err != nil {
			if err.Status == 404 && err.Error == "Invitation not found" {
				if err := NS.DeleteService(notification.SenderId, notification.RecieverId, "group-invitation", notification.GroupId); err != nil {
					return models.NewErrorJson(500, "500 - cannot delete notification join after accept join request", err)
				}
				return nil
			}
			return models.NewErrorJson(500, fmt.Sprintf("500 - %v", err), "cannot cancel invitation request after accept join request")
		}

		if err := NS.DeleteService(notification.SenderId, notification.RecieverId, "group-invitation", notification.GroupId); err != nil {
			return models.NewErrorJson(500, "500 - cannot delete notification join after accept join request", err)
		}

	case "reject":
		if err := NS.groupService.Decline(notification.RecieverId, notification.GroupId, notification.SenderId); err != nil {
			return models.NewErrorJson(500, "500 - cannot decline request", err)
		}

	default:
		return models.NewErrorJson(400, "Bad-Request 400", "Invalid Status")
	}
	return nil
}

// UpdateGroupInvitationRequest updates the group invitation request based on the provided data.
func (NS *NotificationService) UpdateGroupInvitationRequest(data models.Unotif, notification models.Notification) *models.ErrorJson {
	switch data.Status {
	case "accept":
		if err := NS.groupService.Accept(notification.SenderId, notification.GroupId, notification.RecieverId); err != nil {
			return models.NewErrorJson(500, "500 - cannot accept request", err)
		}

		// should check if request exists before canceling //////////////////
		_, err := NS.groupService.RequestToCancel(notification.RecieverId, notification.GroupId);
		if err != nil {
			if err.Status == 404 && err.Error == "Invitation not found" {
				fmt.Println("Invitation not found, deleting notification")
				if err := NS.DeleteService(notification.SenderId, notification.RecieverId, "group-join", notification.GroupId); err != nil {
					return models.NewErrorJson(500, "500 - cannot delete notification join after accept join request", err)
				}
				return nil
			}
			if err.Status == 403 && err.Message == "ERROR!! You are already a member!" {
				fmt.Println("You are already a member, deleting notification")
				if err := NS.DeleteService(notification.SenderId, notification.RecieverId, "group-join", notification.GroupId); err != nil {
					return models.NewErrorJson(500, "500 - cannot delete notification join after accept join request", err)
				}
				return nil
			}
			return models.NewErrorJson(500, "500 - cannot cancel join request after accept join request", err)
		}

		// if the type is cancel then we do not need to delete the notification

		if err := NS.DeleteService(notification.SenderId, notification.RecieverId, "group-join", notification.GroupId); err != nil {
			return models.NewErrorJson(500, "500 - cannot delete notification join after accept join request", err)
		}

	case "reject":
		if err := NS.groupService.Reject(notification.SenderId, notification.GroupId, notification.RecieverId); err != nil {
			return models.NewErrorJson(500, "500 - cannot accept request", err)
		}

	default:
		return models.NewErrorJson(400, "Bad-Request", "Invalid Status")
	}
	return nil
}
