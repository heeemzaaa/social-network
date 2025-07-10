package group

import (
	"fmt"

	"social-network/backend/models"

	"github.com/google/uuid"
)

func (gRepo *GroupRepository) GetGroupEvents(groupId string, offset int64) ([]models.Event, *models.ErrorJson) {
	query := `SELECT concat(users.firstName, " " , users.lastName) AS FullName, 
	group_events.title, group_events.description, group_events.eventTime 
	FROM group_events INNER JOIN users 
	ON group_events.eventCreatorID = users.userID
	WHERE groupID = ?
	LIMIT 20, OFFSET ?
	`
	stmt, err := gRepo.db.Prepare(query)
	if err != nil {
	}
	defer stmt.Close()

	rows, err := stmt.Query(groupId, offset)
	if err != nil {
	}
	defer rows.Close()
	var events []models.Event
	for rows.Next() {
		var event models.Event
		if err := rows.Scan(&event.EventCreator, &event.Title, &event.Description, &event.EventDate); err != nil {
			return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
		}
		events = append(events, event)
	}

	return events, nil
}

func (gRepo *GroupRepository) AddGroupEvent(event *models.Event) (*models.Event, *models.ErrorJson) {
	eventId := uuid.New()
	query := `INSERT INTO group_events 
	(eventID,eventCreatorID,groupID,title,description,eventTime)
	VALUES (?,?,?,?,?,?)
	`
	stmt, err := gRepo.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()
	if _, err = stmt.Exec(eventId, event.EventCreatorId,
		event.GroupId, event.Title, event.Description, event.EventDate); err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}

	return event, nil
}
