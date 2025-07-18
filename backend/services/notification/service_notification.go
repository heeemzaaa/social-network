package notification

import (
	"fmt"
	"slices"
	"social-network/backend/models"
	"social-network/backend/repositories/notification"
	"social-network/backend/utils"
	"strconv"

	"github.com/google/uuid"
)

type NotificationService struct {
	repo *notification.NotifRepository
}

func NewNotifService(repo *notification.NotifRepository) *NotificationService {
	return &NotificationService{repo: repo}
}

func (NS *NotificationService) GetService(userid, queryParam string) ([]models.Notification, *models.ErrorJson) {
	all, err := NS.repo.SelectAllNotification(userid)
	if err != nil {
		return nil, err
	}

	len := len(all)
	nbr, _ := strconv.Atoi(queryParam)

	switch {
	case len <= 10:
		return all, nil

	case nbr <= 0:
		fmt.Println("nbr <= 0")
		return all[:10], nil

	case len <= nbr:
		fmt.Println("greates than notifications : len <= nbr")
		return []models.Notification{}, nil

	case len < nbr+10:
		fmt.Println("len < nbr + 10")
		return all[nbr:], nil

	default:
		fmt.Println("default = [nbr : nbr + 10]")
		return all[nbr : nbr+10], nil
	}
}

func (NS *NotificationService) PostService(data models.Notif, user_id string) (models.Notification, *models.ErrorJson) {
	fmt.Println("POST SERVICE 000 : data reciever = : ", data)

	if user_id != data.Sender_Id {
		fmt.Println("error : sender id : bad request")
	}

	// get user_name by user_id //////////////////////
	sender_name := "AMINE"

	switch data.Type {
	case "follow-private":
		NS.FollowPrivateProfile(data, sender_name)
	case "follow-public":
		NS.FollowPublicProfile(data, sender_name)
	case "group-invitation":
		NS.GroupInvitationRequest(data, sender_name)
	case "group-join":
		NS.GroupJoinRequest(data, sender_name)
	case "group-event":
		NS.GroupEventRequest(data, sender_name)
	default:
		fmt.Println("error: bad request = invalid type ---------")
		return models.Notification{}, nil // (w, models.ErrorJson{Status: 400, Error: "Bad Request - 400", Message: "invalid type"})
	}

	// check if reciever has web socket con //// to brodcast notification
	result, err := NS.repo.SelectNotification(user_id)
	if err != nil {
		return models.Notification{}, err
	}
	fmt.Println("SERVICE 333 : --------- : ", result)
	return result, nil
}

// 1 - follow private profile request
// event : follow profile button : onclick()
func (NS *NotificationService) FollowPrivateProfile(data models.Notif, sender_name string) {
	fmt.Println("SERVICE 111 : private : INSERT START")

	// check reciever id //////////////////////////:

	notification := models.Notification{}
	notification.Sender_Id = data.Sender_Id
	notification.Id = uuid.New().String()
	notification.Seen = false
	notification.Type = data.Type
	notification.Reciever_Id = data.Reciever_Id
	notification.Status = "later"
	notification.Content = fmt.Sprintf("%v sent follow request", sender_name)

	err := NS.repo.InsertNewNotification(notification)
	if err != nil {
			fmt.Println("error = insertion -----")
		// return models.Notification{}, err
	}
	fmt.Println("SERVICE 222 : private : INSERT END ")

	// check if reciever id has web socket connection ///////////////////

	// brodcast ==> notification
}

// 2 - follow public profile request
// event : follow profile button : onclick()
func (NS *NotificationService) FollowPublicProfile(data models.Notif, sender_name string) {
	fmt.Println("SERVICE 111 : public : INSERT START")

	// chech reciever id ///////////////

	notification := models.Notification{}
	notification.Sender_Id = data.Sender_Id
	notification.Id = utils.NewUUID()
	notification.Seen = false
	notification.Type = data.Type
	notification.Status = "none"
	notification.Reciever_Id = data.Reciever_Id
	notification.Content = sender_name + " follow you"

	err := NS.repo.InsertNewNotification(notification)
	if err != nil {
			fmt.Println("error = insertion -----")
		// return models.Notification{}, err
	}
	fmt.Println("SERVICE 222 : public : INSERT END ")
	
	// check if reciever id has web socket connection ///////////////////

	// brodcast ==> notification
}

// 3 - group invitation request
// event : select invitation button : onclick()
func (NS *NotificationService) GroupInvitationRequest(data models.Notif, sender_name string) {
	fmt.Println("SERVICE 111 : invitation : INSERT START")

	// check if group exist from groups by group_name = data.Content //////::
	
	// check reciever id ////////////////:::::

	// check if sender id in group /////////////::

	group_name := "GROUP_NAME"

	notification := models.Notification{}
	notification.Sender_Id = data.Sender_Id
	notification.Id = uuid.New().String()
	notification.Seen = false
	notification.Type = data.Type
	notification.Status = "later"
	notification.Reciever_Id = data.Reciever_Id
	notification.Content = sender_name + " invite you to join club " + group_name

	err := NS.repo.InsertNewNotification(notification)
	if err != nil {
		// return models.Notification{}, err*
			fmt.Println("error = insertion -----")

	}
	fmt.Println("SERVICE 222 : invitaion : INSERT END ")
	
	// check if reciever id has web socket connection ///////////////////

	// brodcast ==> notification
}

// 4 - group join request [admin] //// ------ get reciever id from admin_group
// event : group join button : onclick()
func (NS *NotificationService) GroupJoinRequest(data models.Notif, sender_name string) {
	fmt.Println("SERVICE 111 : join : INSERT START")

	// check if group exist from groups by group_name = data.Content ////////////////:
	
	// get adminID from groups ////////////////////

	group_name := "GROUP_NAME"

	notification := models.Notification{}
	notification.Sender_Id = data.Sender_Id
	notification.Id = uuid.New().String()
	notification.Seen = false
	notification.Type = data.Type
	notification.Status = "later"
	notification.Reciever_Id = "admin___user___id"
	notification.Reciever_Id = data.Reciever_Id // reciever id = group id = admin id
	notification.Content = sender_name + " want join club " + group_name

	err := NS.repo.InsertNewNotification(notification)
	if err != nil {
		// return models.Notification{}, err
			fmt.Println("error = insertion -----")

	}
	fmt.Println("SERVICE 222 : join : INSERT END ")
	
	// check if reciever id has web socket connection ///////////////////

	// brodcast ==> notification
}

// 5 - group event created [group-members]
// event : create group event button : onclick()
func (NS *NotificationService) GroupEventRequest(data models.Notif, sender_name string) {
	fmt.Println("SERVICE 111 : event : INSERT START")

	// check if group exist from groups by group_name = data.Content /////////////:

	// get group members ///////////////:

	group_name := "GROUP_NAME"
	groupMembers := []string{"67e3fbfb-e5f9-4eb1-ab7e-c738f69d3580", "b3cc45f5-b1bb-4c65-871c-75dfe8b5692f", "9f57a431-adf6-42b2-94f3-89c4496d8932", "562b7f42-b132-4a5d-8863-ed111c6487eb", "946257db-65e0-491e-9a09-c33f946f5092"} // get group_admin_user_id from groups where group_name = data.content [group_name]

	// check if the sender in the group
	if !slices.Contains(groupMembers, data.Sender_Id) {
		fmt.Println("error : bad request")
	}

	for _, member_Id := range groupMembers {

		if member_Id == data.Sender_Id {
			continue
		}

		notification := models.Notification{}

		notification.Sender_Id = data.Sender_Id
		notification.Id = uuid.New().String()
		notification.Seen = false
		notification.Type = data.Type
		notification.Status = "later"

		notification.Content = sender_name + " create event at club " + group_name

		notification.Reciever_Id = member_Id

		err := NS.repo.InsertNewNotification(notification)
		if err != nil {
			// return models.Notification{}, err
			fmt.Println("error = insertion -----")
		}
		
		// check if reciever id has web socket connection ///////////////////

		// brodcast ==> notification
	}

	fmt.Println("SERVICE 222 : event : INSERT END ")
}

// 1 -- function check if user has web socket connection

// 2 -- function brodcast notification data

// 3 -- function check if user has notification containe false seen

// 4 -- function update status

// 5 -- function update seen
