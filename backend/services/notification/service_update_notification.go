package notification

import (
	"log"
	"social-network/backend/models"
	"social-network/backend/repositories/notification"
	"social-network/backend/repositories/profile"
)

type NotificationServiceUpdate struct {
	repo  *notification.NotifRepository
	repo2 *profile.ProfileRepository
}

func NewNotifServiceUpdate(repo *notification.NotifRepository, repo2 *profile.ProfileRepository) *NotificationServiceUpdate {
	return &NotificationServiceUpdate{
		repo:  repo,
		repo2: repo2,
	}
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

	switch data.Status {
		case "accept": 
			err := NUS.repo2.AcceptedRequest(user_id, notification.Sender_Id)
			if err != nil {
				return models.NewErrorJson(500, "500 - cannot accept request", err)
			}
		case "reject":
			err := NUS.repo2.RejectedRequest(user_id, notification.Sender_Id)
			if err != nil {
				return models.NewErrorJson(500, "500 - cannot accept request", err)
			}
		default:
		return models.NewErrorJson(400, "Bad-Request 400", "Invalid----Status")
	}


	if errJson = NUS.repo.UpdateStatusById(notification.Id, data.Status); errJson != nil {
		return errJson
	}
	if errJson := NUS.repo.UpdateSeen(notification.Id); errJson != nil {
		return errJson
	}
	return nil
}
