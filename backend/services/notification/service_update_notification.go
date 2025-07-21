package notification

import (
	"log"
	"social-network/backend/models"
	"social-network/backend/repositories/notification"
)

type NotificationServiceUpdate struct {
	repo *notification.NotifRepository
}

func NewNotifServiceUpdate(repo *notification.NotifRepository) *NotificationServiceUpdate {
	return &NotificationServiceUpdate{repo: repo}
}

func (NUS *NotificationServiceUpdate) UpdateService(data models.Unotif, user_id string) *models.ErrorJson {
	log.Println("START UPDATE SERVICE ----- REQUEST DATA = ", data)

	if data.Status != "accept" && data.Status != "reject" {
		return models.NewErrorJson(400, "Bad-Request 400", "Invalid----Status")
	}

	notification, errJson := NUS.repo.SelectNotification(data.Notif_Id)
	if errJson != nil {
		return models.NewErrorJson(errJson.Status, errJson.Error, errJson.Message)
	}

	if  notification.Status != "later" { // notification.Reciever_Id == user_id ||
		return models.NewErrorJson(400, "Bad-Request 400", "Invalid----Operation")
	}
	if errJson = NUS.repo.UpdateStatus(notification.Id, data.Status); errJson != nil {
		return errJson
	}
	if errJson := NUS.repo.UpdateSeen(notification.Id); errJson != nil {
		return errJson
	}
	return nil
}
