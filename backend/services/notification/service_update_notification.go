package notification

import (
	"log"
	"social-network/backend/models"
	GR "social-network/backend/repositories/group"
	"social-network/backend/repositories/notification"
	"social-network/backend/repositories/profile"
	// GS "social-network/backend/service/group"
)

type NotificationServiceUpdate struct {
	repo  *notification.NotifRepository
	repo2 *profile.ProfileRepository
	gr    *GR.GroupRepository
	// gs *GS.GroupService
}

func NewNotifServiceUpdate(repo *notification.NotifRepository, repo2 *profile.ProfileRepository, gr *GR.GroupRepository) *NotificationServiceUpdate {
	return &NotificationServiceUpdate{
		repo:  repo,
		repo2: repo2,
		gr:    gr,
	}
}
func (NUS *NotificationServiceUpdate) UpdateFollowPrivateProfile(data models.Unotif, notification models.Notification) *models.ErrorJson {
	switch data.Status {
	case "accept":
		err := NUS.repo2.AcceptedRequest(notification.Reciever_Id, notification.Sender_Id)
		if err != nil {
			return models.NewErrorJson(500, "500 - cannot accept request", err)
		}
	case "reject":
		err := NUS.repo2.RejectedRequest(notification.Reciever_Id, notification.Sender_Id)
		if err != nil {
			return models.NewErrorJson(500, "500 - cannot accept request", err)
		}
	default:
		return models.NewErrorJson(400, "400 - Bad Request", "Invalid Status:" + data.Status)
	}
	return nil
}

func (NUS *NotificationServiceUpdate) UpdateGroupJoinRequest(data models.Unotif, notification models.Notification) *models.ErrorJson {
	switch data.Status {
	case "accept":
		err := NUS.gr.Approve(notification.GroupId, notification.Sender_Id)
		if err != nil {
			return models.NewErrorJson(500, "500 - cannot accept request", err)
		}
	case "reject":
		err := NUS.gr.Decline(notification.GroupId, notification.Sender_Id)
		if err != nil {
			return models.NewErrorJson(500, "500 - cannot accept request", err)
		}
	default:
		return models.NewErrorJson(400, "Bad-Request 400", "Invalid----Status")
	}
	return nil
}
func (NUS *NotificationServiceUpdate) UpdateService(data models.Unotif, user_id string) *models.ErrorJson {
	log.Println("START UPDATE SERVICE ----- REQUEST DATA = ", data)

	notification, errJson := NUS.repo.SelectNotification(data.Notif_Id)
	if errJson != nil {
		return models.NewErrorJson(errJson.Status, errJson.Error, errJson.Message)
	}

	// notification.Reciever_Id == user_id ||
	if notification.Status != "later" {
		return models.NewErrorJson(400, "Bad-Request 400", "Invalid----Operation")
	}

	switch data.Type {
	case "follow-private":
		errJson := NUS.UpdateFollowPrivateProfile(data, notification)
		if errJson != nil {
			return errJson
		}
	case "follow-public":
		// errJson = NS.FollowPublicProfile(data)
	case "group-invitation":
		// errJson = NS.GroupInvitationRequest(data)
	case "group-join":
		errJson := NUS.UpdateGroupJoinRequest(data, notification)
		if errJson != nil {
			return errJson
		}
	case "group-event":
		// errJson = NS.GroupEventRequest(data)
	default:
		return models.NewErrorJson(400, "Bad Request - 400", "invalid type")
	}


	if errJson = NUS.repo.UpdateStatusById(notification.Id, data.Status); errJson != nil {
		return errJson
	}
	if errJson := NUS.repo.UpdateSeen(notification.Id); errJson != nil {
		return errJson
	}
	return nil
}
