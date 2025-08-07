package group

import (
	"database/sql"
	"fmt"

	"social-network/backend/models"
	"social-network/backend/utils"
)

// get the event created of a specific group
func (gRepo *GroupRepository) GetGroupEvents(groupId, userId, offset string) ([]models.Event, *models.ErrorJson) {
	events := []models.Event{}
	var where string
	if offset == "0" {
		where = `WHERE group_events.groupID =  ?`
	} else {
		where = `WHERE group_events.groupID =  ? AND group_events.createdAt < (
			select createdAt from group_events WHERE eventID = ? 
		)`
	}
	query := fmt.Sprintf(`
	WITH
    cte_interested AS (
        SELECT
            group_events.eventID as ID,
            group_events.title,
            group_event_users.actionChosen AS chosen
        FROM
            group_events
            INNER JOIN group_event_users ON group_event_users.eventID = group_events.eventID
            WHERE group_event_users.userID = ?
        GROUP BY
            group_events.eventID
    )


    SELECT
        group_events.groupID,
        group_events.eventID,
        group_events.eventCreatorID,
        concat (users.firstName, " ", users.lastName) AS FullName,
        users.nickname,
        users.avatarPath,
        group_events.title,
        group_events.description,
        group_events.eventTime,
    	group_events.createdAt,
        coalesce(cte_interested.chosen,0)'ba39d493-543b-4c9b-8258-81c7f4df8203'
    FROM
        group_events
        INNER JOIN users ON group_events.eventCreatorID = users.userID
        LEFT JOIN cte_interested ON cte_interested.ID = group_events.eventID
    %v
    ORDER BY group_events.createdAt DESC
    LIMIT 20
	`, where)
	stmt, err := gRepo.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v 1", err)}
	}
	defer stmt.Close()
	args := []any{userId, groupId}
	if offset != "0" {
		args = append(args, offset)
	}

	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v 2", err)}
	}
	defer rows.Close()
	for rows.Next() {
		var event models.Event
		if err := rows.Scan(
			&event.Group.GroupId,
			&event.EventId,
			&event.EventCreator.Id,
			&event.EventCreator.FullName,
			&event.EventCreator.Nickname,
			&event.EventCreator.ImagePath,
			&event.Title,
			&event.Description,
			&event.EventDate,
			&event.CreatedAt,
			&event.Going,
		); err != nil {
			return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v 2", err)}
		}
		events = append(events, event)
	}

	return events, nil
}

// add an event in a specific group
func (gRepo *GroupRepository) AddGroupEvent(event *models.Event) (*models.Event, *models.ErrorJson) {
	fmt.Println("event ", event)
	eventId := utils.NewUUID()
	query := `
	INSERT INTO
    group_events (
        eventID,
        eventCreatorID,
        groupID,
        title,
        description,
        eventTime
    )
    VALUES
    (?, ?, ?, ?, ?, ?) RETURNING eventID,
    eventCreatorID,
    groupID,
    title,
    description,
    eventTime,
    createdAt,
    (
        SELECT
            concat (firstName, ' ', lastName)
        FROM
            users
        WHERE
            users.userID = ?
    ) AS fullName,
    (
        SELECT
            nickname
        FROM
            users
        WHERE
            users.userID = ?
    ),
	(
        SELECT
            avatarPath
        FROM
            users
        WHERE
            users.userID = ?
    ),
	(
        SELECT
            title
        FROM
            groups
        WHERE
            groups.groupID = ?
    )

	;
	`
	stmt, err := gRepo.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v huna", err)}
	}
	defer stmt.Close()

	event_created := models.Event{}
	if err = stmt.QueryRow(eventId, event.EventCreator.Id,
		event.Group.GroupId, event.Title, event.Description, event.EventDate,
		event.EventCreator.Id, event.EventCreator.Id, event.EventCreator.Id, event.Group.GroupId).Scan(
		&event_created.EventId,
		&event_created.EventCreator.Id,
		&event_created.Group.GroupId,
		&event_created.Title,
		&event_created.Description,
		&event_created.EventDate,
		&event_created.CreatedAt,
		&event_created.EventCreator.FullName,
		&event_created.EventCreator.Nickname,
		&event_created.EventCreator.ImagePath,
		&event_created.Group.Title,
	); err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	fmt.Printf("event_created: %v\n", event_created)
	// add the user to the user_events_table
	eventUserId := utils.NewUUID()
	queryAdded := `INSERT INTO group_event_users 
	(ID , eventID , groupID, userID) VALUES (?,?,?,?)`
	stmt, err = gRepo.db.Prepare(queryAdded)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	_, err = stmt.Exec(eventUserId, eventId, event.Group.GroupId, event.EventCreator.Id)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	return &event_created, nil
}

// Get the details ( the card of a specific event )
// userid will be needed to see if he has done some interest on the event
func (gRepo *GroupRepository) GetEventDetails(eventId, groupId, userId string) (*models.Event, *models.ErrorJson) {
	event := &models.Event{}
	query := `
	WITH
    cte_liked AS (
        SELECT
            group_event_users.actionChosen as action_chosen
        FROM
            users
            INNER JOIN group_event_users ON users.userID = group_event_users.userID
            AND users.userID = ?
    )
	SELECT
		eventID,
		groupID,
		title,
		description,
		eventTime,
		group_events.createdAt,
		concat (users.firstName, " ", users.lastName) AS fullName,
		(
			SELECT
				cte_liked.action_chosen
			FROM
				cte_liked
		) as actionChosen
	FROM
		group_events
		INNER JOIN users ON group_events.eventCreatorID = users.userID
	WHERE
		eventID = ?
		and groupID = ?

	`
	stmt, err := gRepo.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v hna", err)}
	}
	defer stmt.Close()

	if err := stmt.QueryRow(userId, eventId, groupId).Scan(&event.EventId,
		&event.Group.GroupId,
		&event.Title,
		&event.Description,
		&event.EventDate,
		&event.CreatedAt,
		&event.EventCreator,
		&event.Going,
	); err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v lhiih", err)}
	}
	return event, nil
}

// show the interest to a specific event

// here will be adding the action chosen for the first time and then we need to update it again
func (gRepo *GroupRepository) AddAction(actionChosen *models.UserEventAction) (*models.UserEventAction, *models.ErrorJson) {
	action_created := &models.UserEventAction{}
	actionID := utils.NewUUID()
	query := `INSERT INTO group_event_users 
	(ID , eventID , groupID, userID, actionChosen) VALUES (?,?,?,?,?) RETURNING actionChosen`
	stmt, err := gRepo.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v 12", err)}
	}
	defer stmt.Close()

	err = stmt.QueryRow(actionID, actionChosen.EventId, actionChosen.GroupId, actionChosen.UserId, actionChosen.Action).Scan(
		&action_created.Action)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v 2", err)}
	}
	action_created.Id = actionID
	return action_created, nil
}

// here we'll be updating the action we had chosen before and then proceed and return it
// here we'll be going to update to going whatever the case

func (gRepo *GroupRepository) UpdateToGoing(actionChosen *models.UserEventAction) (*models.UserEventAction, *models.ErrorJson) {
	action_created := &models.UserEventAction{}
	query := `UPDATE group_event_users SET actionChosen = CASE actionChosen
              WHEN 0 THEN 1
			  WHEN -1 THEN 1
              ELSE 0
              END
	          WHERE eventID = ? AND groupID= ? AND userID= ?
			  RETURNING actionChosen;`

	// finahyaa l preparation a diik oumaaaaaaaaayma :)

	stmt, err := gRepo.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	err = stmt.QueryRow(actionChosen.EventId, actionChosen.GroupId, actionChosen.UserId).Scan(&action_created.Action)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	return action_created, nil
}

func (gRepo *GroupRepository) UpdateToNotGoing(actionChosen *models.UserEventAction) (*models.UserEventAction, *models.ErrorJson) {
	action_created := &models.UserEventAction{}
	query := `UPDATE group_event_users SET actionChosen = CASE actionChosen
              WHEN 0 THEN -1
			  WHEN 1 THEN -1
              ELSE 0
              END
	          WHERE eventID = ? AND groupID= ? AND userID= ?
			  RETURNING actionChosen;`

	// finahyaa l preparation a diik oumaaaaaaaaayma :)

	stmt, err := gRepo.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v 2", err)}
	}
	defer stmt.Close()

	err = stmt.QueryRow(actionChosen.EventId, actionChosen.GroupId, actionChosen.UserId).Scan(&action_created.Action)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v 3", err)}
	}
	return action_created, nil
}

func (gRepo *GroupRepository) HanldeAction(actionChosen *models.UserEventAction) (*models.UserEventAction, *models.ErrorJson) {
	query := `SELECT * FROM group_event_users WHERE
	eventID = ? AND groupID = ? AND userID = ?
	`
	stmt, err := gRepo.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v dddd ", err)}
	}
	defer stmt.Close()

	reaction_existed := &models.UserEventAction{}
	if err := stmt.QueryRow(actionChosen.EventId, actionChosen.GroupId, actionChosen.UserId).Scan(
		&reaction_existed.Id,
		&reaction_existed.EventId,
		&reaction_existed.GroupId,
		&reaction_existed.UserId,
		&reaction_existed.Action); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v jjj", err)}
	}
	return reaction_existed, nil
}
