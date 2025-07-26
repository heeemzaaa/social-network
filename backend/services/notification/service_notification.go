package notification

import (
	"fmt"
	"sort"
	"strconv"
	"time"

	"social-network/backend/models"
	"social-network/backend/repositories/auth"
	"social-network/backend/repositories/notification"
	"social-network/backend/utils"

	"github.com/google/uuid"
)

type NotificationService struct {
	repo2 *auth.AuthRepository
	repo  *notification.NotifRepository
}

func NewNotifService(repo *notification.NotifRepository, repo2 *auth.AuthRepository) *NotificationService {
	return &NotificationService{
		repo:  repo,
		repo2: repo2,
	}
}

func (NS *NotificationService) ToggleAllSeenFalse(notifications []models.Notification) *models.ErrorJson {
	for _, notification := range notifications {
		if errJson := NS.repo.UpdateSeen(notification.Id); errJson != nil {
			return errJson
		}
	}
	return nil
}

// toggle all notifications status by type
func (NS *NotificationService) ToggleAllStaus(notifications []models.Notification, value, notifType string) *models.ErrorJson {
	for _, notification := range notifications {
		if errJson := NS.repo.UpdateStatusById(notification.Id, value); errJson != nil {
			return errJson
		}
	}
	return nil
}

// toggle all notifications status by type
func (NS *NotificationService) ToggleStaus(userID, reciever, value, notifType string) *models.ErrorJson {
	if errJson := NS.repo.UpdateStatusByType(userID, reciever, value, notifType); errJson != nil {
		return errJson
	}
	return nil
}

func (NS *NotificationService) GetAllNotifService(user_id, notifType string) ([]models.Notification, *models.ErrorJson) {
	all, err := NS.repo.SelectAllNotification(user_id)
	if err != nil {
		return nil, err
	}
	return all, nil
}

func (NS *NotificationService) DeleteService(reciever, sender, notifType, value string) *models.ErrorJson {
	if notifType != "follow-private" {
		if errJson := NS.repo.DeleteGroupNotification(sender, reciever, notifType, value); errJson != nil {
			return errJson
		}
	} else {
		if errJson := NS.repo.DeleteFollowNotification(sender, reciever, notifType, value); errJson != nil {
			return errJson
		}
	}
	return nil
}

func (NS *NotificationService) IsHasSeenFalse(user_id string) (bool, *models.ErrorJson) {
	isValid, errJson := NS.repo.IsHasSeenFalse(user_id)
	if errJson != nil {
		return false, errJson
	}
	return isValid, nil
}

func (NS *NotificationService) GetService(user_id, queryParam string) ([]models.Notification, *models.ErrorJson) {

	all, err := NS.repo.SelectAllNotification(user_id)
	if err != nil {
		return nil, err
	}
	len := len(all)
	nbr, _ := strconv.Atoi(queryParam)
	sort.Slice(all, func(i, j int) bool {
		// Check if one is "later" and the other isn't
		if all[i].Status == "later" && all[j].Status != "later" {
			return true
		}
		if all[i].Status != "later" && all[j].Status == "later" {
			return false
		}
		// If both are "later" or both are not, sort by CreatedAt descending
		return all[i].CreatedAt.After(all[j].CreatedAt)
	})

	switch {
	case len <= 10:
		return all, nil
	case nbr <= 0:
		return all[:10], nil
	case len <= nbr:
		return []models.Notification{}, nil
	case len < nbr+10:
		return all[nbr:], nil
	}
	return all[nbr : nbr+10], nil
}

// insert new notification after event hapen
func (NS *NotificationService) PostService(data models.Notif) *models.ErrorJson {

	name, errJson := NS.repo2.GetUserNameById(data.SenderId)
	if errJson  != nil {
		return errJson
	}
	data.SenderFullName = name

	switch data.Type {
	case "follow-private":
		errJson = NS.FollowPrivateProfile(data)
	case "follow-public":
		errJson = NS.FollowPublicProfile(data)
	case "group-invitation":
		errJson = NS.GroupInvitationRequest(data)
	case "group-join":
		errJson = NS.GroupJoinRequest(data)
	case "group-event":
		errJson = NS.GroupEventRequest(data)
	default:
		return models.NewErrorJson(400, "Bad Request - 400", "invalid type")
	}

	if errJson != nil {
		return errJson
	}

	fmt.Println("ADD NOTIFICATION SUCCES --------- !")
	return nil
}

// - follow private profile request
func (NS *NotificationService) FollowPrivateProfile(data models.Notif) *models.ErrorJson {

	notification := models.Notification{}
	notification.SenderId = data.SenderId
	notification.Id = uuid.New().String()
	notification.Seen = false
	notification.Type = data.Type
	notification.RecieverId = data.RecieverId
	notification.GroupId = data.GroupId
	notification.Status = "later"
	notification.Content = fmt.Sprintf("%v sent follow request", data.SenderFullName)
	notification.CreatedAt = time.Now()

	if err := NS.repo.InsertNewNotification(notification); err != nil {
		fmt.Println("error private = insertion ---------", err)
		return err
	}
	return nil
}

// - follow public profile request
func (NS *NotificationService) FollowPublicProfile(data models.Notif) *models.ErrorJson {
	///////////////////////////////////////////////////  golna madich nkhedmo 3la had l case //////////
	notification := models.Notification{}
	notification.SenderId = data.SenderId
	notification.Id = utils.NewUUID()
	notification.Seen = false
	notification.GroupId = data.GroupId
	notification.Type = data.Type
	notification.Status = "none"
	notification.RecieverId = data.RecieverId
	notification.Content = data.SenderFullName + " follow you"
	notification.CreatedAt = time.Now()

	if err := NS.repo.InsertNewNotification(notification); err != nil {
		fmt.Println("error public = insertion ---------", err)
		return err
	}
	return nil
}

// - group invitation request
func (NS *NotificationService) GroupInvitationRequest(data models.Notif) *models.ErrorJson {

	notification := models.Notification{}
	notification.SenderId = data.SenderId
	notification.Id = uuid.New().String()
	notification.Seen = false
	notification.Type = data.Type
	notification.Status = "later"
	notification.RecieverId = data.RecieverId
	notification.Content = data.SenderFullName + " invite you to join club " + data.GroupName
	notification.GroupId = data.GroupId
	notification.CreatedAt = time.Now()

	if err := NS.repo.InsertNewNotification(notification); err != nil {
		fmt.Println("error invitation = insertion ---------", err)
		return err
	}
	return nil
}

// - group join request [admin]
func (NS *NotificationService) GroupJoinRequest(data models.Notif) *models.ErrorJson {

	notification := models.Notification{}
	notification.SenderId = data.SenderId
	notification.Id = utils.NewUUID()
	notification.Seen = false
	notification.Type = data.Type
	notification.GroupId = data.GroupId
	notification.Status = "later"
	notification.RecieverId = data.RecieverId
	notification.Content = data.SenderFullName + " want join club " + data.GroupName
	notification.CreatedAt = time.Now()

	if err := NS.repo.InsertNewNotification(notification); err != nil {
		fmt.Println("error join = insertion ---------", err)
		return err
	}
	return nil
}

// 5 - group event created [group-members]
//
//	var notif = models.Notif{
//		SenderId: "current user id == event maker",
//		SenderFullName: "sender full name",
//		RecieverId: "reciever id == one of group member",
//		ReceiverFullName: "reciever full name",
//		Type: "group-event",
//		GroupName: "group name",
//	}
func (NS *NotificationService) GroupEventRequest(data models.Notif) *models.ErrorJson {

	notification := models.Notification{}

	notification.SenderId = data.SenderId
	notification.Id = uuid.New().String()
	notification.Seen = false
	notification.GroupId = data.GroupId
	notification.Type = data.Type
	notification.Status = "later"

	notification.Content = data.SenderFullName + " create event at club " + data.GroupName
	notification.CreatedAt = time.Now()

	notification.RecieverId = data.RecieverId

	if err := NS.repo.InsertNewNotification(notification); err != nil {
		fmt.Println("error event = insertion ---------", err)
		return err
	}
	return nil
}
