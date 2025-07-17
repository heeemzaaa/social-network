package group

import (
	"database/sql"
	"fmt"

	"social-network/backend/models"
	"social-network/backend/utils"
)

// get the event created of a specific group
func (gRepo *GroupRepository) GetGroupEvents(groupId string, offset int64) ([]models.Event, *models.ErrorJson) {
	events := []models.Event{}
	query := `SELECT eventID, eventCreatorID,  concat(users.firstName, " " , users.lastName) AS FullName, 
	group_events.title, group_events.description, group_events.eventTime 
	FROM group_events INNER JOIN users 
	ON group_events.eventCreatorID = users.userID
	WHERE groupID = ?
	LIMIT 20 OFFSET ?
	`
	stmt, err := gRepo.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v 1", err)}
	}
	defer stmt.Close()

	rows, err := stmt.Query(groupId, offset)
	if err != nil {
		if err == sql.ErrNoRows {
			return events, nil
		}
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v 2", err)}
	}
	defer rows.Close()
	for rows.Next() {
		var event models.Event
		if err := rows.Scan(
			&event.EventId,
			&event.EventCreatorId,
			&event.EventCreator,
			&event.Title,
			&event.Description,
			&event.EventDate); err != nil {
			return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v 2", err)}
		}
		events = append(events, event)
	}

	return events, nil
}

// add an event in a specific group
func (gRepo *GroupRepository) AddGroupEvent(event *models.Event) (*models.Event, *models.ErrorJson) {
	eventId := utils.NewUUID()
	event.EventId = eventId
	query := `INSERT INTO group_events 
	(eventID,eventCreatorID,groupID,title,description,eventTime)
	VALUES (?,?,?,?,?,?)
	`
	stmt, err := gRepo.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()
	if _, err = stmt.Exec(event.EventId, event.EventCreatorId,
		event.GroupId, event.Title, event.Description, event.EventDate); err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}

	return event, nil
}

// Get the details ( the card of a specific event )

func (gRepo *GroupRepository) GetEventDetails(eventId, userId, groupId string) (*models.Event, *models.ErrorJson) {
	event := &models.Event{}
	query := `
	SELECT * FROM group_events 
	WHERE eventID = ? and groupID = ? 
	`
	stmt, err := gRepo.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	if err := stmt.QueryRow(eventId, groupId).Scan(&event.EventId,
		&event.EventCreatorId,
		&event.GroupId,
		&event.Title,
		&event.Description,
		&event.EventDate,
		&event.CreatedAt); err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	return event, nil
}

// show the interest to a specific event

// here will be adding the action chosen for the first time and then we need to update it again
func (gRepo *GroupRepository) AddAction(userId, groupId, eventId string, action int) (*models.UserEventAction, *models.ErrorJson) {
	action_created := &models.UserEventAction{}
	actionID := utils.NewUUID()
	query := `INSERT INTO group_event_users 
	(ID , eventID , groupID, userID, actionChosen) VALUES (?,?,?,?) RETURNING actionChosen`
	stmt, err := gRepo.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v 1", err)}
	}
	defer stmt.Close()
	err = stmt.QueryRow(actionID, eventId, groupId, userId, action).Scan(
		&action_created.Action)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v 2", err)}
	}
	action_created.Id = actionID
	return action_created, nil
}


// here we'll be updating the action we had chosen before and then proceed and return it
// here we'll be going to update to going whatever the case 

func (gRepo *GroupRepository) UpdateToGoing(userId, groupId, eventId string) (*models.UserEventAction, *models.ErrorJson) {
	action_created := &models.UserEventAction{}
	query := `UPDATE group_reactions SET actionChosen = CASE actionChosen
              WHEN 0 THEN 1
			  WHEN -1 THEN 1
              ELSE 0
              END
	          WHERE eventID = ? AND groupID= ? AND userID= ?
			  RETURNING actionChosen;`

	// finahyaa l preparation a diik oumaaaaaaaaayma :)

	stmt, err := gRepo.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v ", err)}
	}
	defer stmt.Close()

	err = stmt.QueryRow(eventId, groupId, userId).Scan(&action_created.Action)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v ", err)}
	}

	return action_created, nil
}

func (gRepo *GroupRepository) UpdateToNotGoing(userId, groupId, eventId string) (*models.UserEventAction, *models.ErrorJson) {
	action_created := &models.UserEventAction{}
	query := `UPDATE group_reactions SET actionChosen = CASE actionChosen
              WHEN 0 THEN -1
			  WHEN 1 THEN -1
              ELSE 0
              END
	          WHERE eventID = ? AND groupID= ? AND userID= ?
			  RETURNING actionChosen;`

	// finahyaa l preparation a diik oumaaaaaaaaayma :)

	stmt, err := gRepo.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v ", err)}
	}
	defer stmt.Close()

	err = stmt.QueryRow(eventId, groupId, userId).Scan(&action_created.Action)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v ", err)}
	}

	return action_created, nil
}




func (gRepo *GroupRepository) HanldeAction(eventID, userID, groupID string) (*models.UserEventAction, *models.ErrorJson) {
	query := `SELECT * FROM group_event_users WHERE
	eventID = ? AND groupID = ? AND userID = ?
	`
	stmt, err := gRepo.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v ", err)}
	}
	reaction_existed := &models.UserEventAction{}
	if err := stmt.QueryRow(eventID, groupID, userID).Scan(
		&reaction_existed.Id,
		&reaction_existed.EventId,
		&reaction_existed.GroupId,
		&reaction_existed.UserId,
		&reaction_existed.Action); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v jjj", err)}
	}
	return reaction_existed, nil
}
