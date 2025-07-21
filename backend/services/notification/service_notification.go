package notification

import (
	"fmt"
	"slices"
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

func (NS *NotificationService) ToggleSeenFalse(notifications []models.Notification) *models.ErrorJson {
	for _, notification := range notifications {
		if errJson := NS.repo.UpdateSeen(notification.Id); errJson != nil {
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
		return all[i].CreatedAt.After(all[j].CreatedAt) // sort notification by time
	})

	switch {
	case len <= 10:
		return all, nil
	case nbr <= 0:
		// fmt.Println("nbr <= 0")
		return all[:10], nil
	case len <= nbr:
		// fmt.Println("greates than notifications : len <= nbr")
		return []models.Notification{}, nil
	case len < nbr+10:
		// fmt.Println("len < nbr + 10")
		return all[nbr:], nil
	}
	// fmt.Println("default = [nbr : nbr + 10]")
	return all[nbr : nbr+10], nil
	// return []models.Notification{}, nil
}

func (NS *NotificationService) PostService(data models.Notif, user_id string) *models.ErrorJson {
	if user_id != data.SenderId {
		return models.NewErrorJson(400, "bad - request - 400", "sender != current user")
	}

	sender_name, errJson := NS.repo2.GetUserNameById(user_id)
	if errJson != nil {
		return &models.ErrorJson{Status: errJson.Status, Message: errJson.Message}
	}
	// reciever_name, errJson := NH.US.GetUserName(user_Id.String())
	// if errJson != nil {
	// 	utils.WriteJsonErrors(w, models.ErrorJson{Status: errJson.Status, Message: errJson.Message})
	// 	return
	// }

	var errInse *models.ErrorJson
	switch data.Type {
	case "follow-private":
		errInse = NS.FollowPrivateProfile(data, sender_name)
	case "follow-public":
		errInse = NS.FollowPublicProfile(data, sender_name)
	case "group-invitation":
		errInse = NS.GroupInvitationRequest(data, sender_name)
	case "group-join":
		errInse = NS.GroupJoinRequest(data, sender_name)
	case "group-event":
		errInse = NS.GroupEventRequest(data, sender_name)
	default:
		return models.NewErrorJson(400, "Bad Request - 400", "invalid type")
	}

	if errInse != nil {
		return errInse
	}

	// check if reciever has web socket con //// to brodcast notification
	// result, err := NS.repo.SelectNotification(user_id)
	// if err != nil {
	// 	return models.Notification{}, err
	// }
	fmt.Println("SERVICE 333 : --------- : ")
	return nil
}

// 1 - follow private profile request
// event : follow profile button : onclick()
func (NS *NotificationService) FollowPrivateProfile(data models.Notif, sender_name string) *models.ErrorJson {
	// fmt.Println("SERVICE 111 : private : INSERT START")

	// check reciever id //////////////////////////:

	notification := models.Notification{}
	notification.Sender_Id = data.SenderId
	notification.Id = uuid.New().String()
	notification.Seen = false
	notification.Type = data.Type
	notification.Reciever_Id = data.RecieverId
	notification.Status = "later"
	notification.Content = fmt.Sprintf("%v sent follow request", sender_name)
	notification.CreatedAt = time.Now()

	if err := NS.repo.InsertNewNotification(notification); err != nil {
		fmt.Println("error = insertion ---------", err)
		return err
	}
	// fmt.Println("SERVICE 222 : private : INSERT END ")

	// check if reciever id has web socket connection ///////////////////

	// brodcast ==> notification

	return nil
}

// 2 - follow public profile request
// event : follow profile button : onclick()
func (NS *NotificationService) FollowPublicProfile(data models.Notif, sender_name string) *models.ErrorJson {
	// fmt.Println("SERVICE 111 : public : INSERT START")

	// chech reciever id ///////////////

	notification := models.Notification{}
	notification.Sender_Id = data.SenderId
	notification.Id = utils.NewUUID()
	notification.Seen = false
	notification.Type = data.Type
	notification.Status = "none"
	notification.Reciever_Id = data.RecieverId
	notification.Content = sender_name + " follow you"
	notification.CreatedAt = time.Now()

	if err := NS.repo.InsertNewNotification(notification); err != nil {
		fmt.Println("error = insertion ---------", err)
		return err
	}
	// fmt.Println("SERVICE 222 : public : INSERT END ")

	// check if reciever id has web socket connection ///////////////////

	// brodcast ==> notification
	return nil
}

// 3 - group invitation request
// event : select invitation button : onclick()
func (NS *NotificationService) GroupInvitationRequest(data models.Notif, sender_name string) *models.ErrorJson {
	// fmt.Println("SERVICE 111 : invitation : INSERT START")

	// check if group exist from groups by group_name = data.Content //////::

	// check reciever id ////////////////:::::

	// check if sender id in group /////////////::

	group_name := "GROUP_NAME"

	notification := models.Notification{}
	notification.Sender_Id = data.SenderId
	notification.Id = uuid.New().String()
	notification.Seen = false
	notification.Type = data.Type
	notification.Status = "later"
	notification.Reciever_Id = data.RecieverId
	notification.Content = sender_name + " invite you to join club " + group_name
	notification.CreatedAt = time.Now()

	if err := NS.repo.InsertNewNotification(notification); err != nil {
		fmt.Println("error = insertion ---------", err)
		return err
	}
	// fmt.Println("SERVICE 222 : invitaion : INSERT END ")

	// check if reciever id has web socket connection ///////////////////

	// brodcast ==> notification
	return nil
}

// 4 - group join request [admin] //// ------ get reciever id from admin_group
// event : group join button : onclick()
func (NS *NotificationService) GroupJoinRequest(data models.Notif, sender_name string) *models.ErrorJson {
	// fmt.Println("SERVICE 111 : join : INSERT START")

	// check if group exist from groups by group_name = data.Content ////////////////:

	// get adminID from groups ////////////////////

	group_name := "GROUP_NAME"

	notification := models.Notification{}
	notification.Sender_Id = data.SenderId
	notification.Id = utils.NewUUID()
	notification.Seen = false
	notification.Type = data.Type
	notification.Status = "later"
	notification.Reciever_Id = data.RecieverId // reciever id = group id = admin id
	notification.Content = sender_name + " want join club " + group_name
	notification.CreatedAt = time.Now()

	if err := NS.repo.InsertNewNotification(notification); err != nil {
		fmt.Println("error = insertion ---------", err)
		return err
	}
	// fmt.Println("SERVICE 222 : join : INSERT END ")

	// check if reciever id has web socket connection ///////////////////

	// brodcast ==> notification
	return nil
}

// 5 - group event created [group-members]
// event : create group event button : onclick()
func (NS *NotificationService) GroupEventRequest(data models.Notif, sender_name string) *models.ErrorJson {
	// fmt.Println("SERVICE 111 : event : INSERT START")

	// check if group exist from groups by group_name = data.Content /////////////:

	// get group members ///////////////:

	group_name := "GROUP_NAME"
	groupMembers := []string{"ffc54fb8-2d14-4f83-a196-062c976e3243", "e2e47e63-ff05-473d-a2af-4c5d9366c1a7", "61b2ae60-1aae-48e6-8997-e2943ebb4e74", "80710ff4-8b05-4b5d-8fb9-b34b98022e0c", "c56c4546-b5ae-4c96-8c70-8f6b2e36f69c", "05dd2f42-bb69-4375-a661-353217a0d574"} // get group_admin_user_id from groups where group_name = data.content [group_name]

	// check if the sender in the group
	if !slices.Contains(groupMembers, data.SenderId) {
		fmt.Println("error : bad request")
	}

	for _, member_Id := range groupMembers {

		if member_Id == data.SenderId {
			continue
		}

		notification := models.Notification{}

		notification.Sender_Id = data.SenderId
		notification.Id = uuid.New().String()
		notification.Seen = false
		notification.Type = data.Type
		notification.Status = "later"

		notification.Content = sender_name + " create event at club " + group_name
		notification.CreatedAt = time.Now()

		notification.Reciever_Id = member_Id

		if err := NS.repo.InsertNewNotification(notification); err != nil {
			fmt.Println("error = insertion ---------", err)
			return err
		}

		// check if reciever id has web socket connection ///////////////////

		// brodcast ==> notification
	}

	// fmt.Println("SERVICE 222 : event : INSERT END ")
	return nil
}
