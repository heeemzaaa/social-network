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

func (gRepo *GroupRepository) IntersetedOrNot(event *models.Event, userId string) *models.ErrorJson {
	
	return nil
}
