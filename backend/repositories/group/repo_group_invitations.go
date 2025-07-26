package group

import (
	"database/sql"
	"fmt"

	"social-network/backend/models"
	"social-network/backend/utils"
)

func (gRepo *GroupRepository) InviteToJoin(userId, groupId string, userToInvite string) *models.ErrorJson {
	fmt.Println(userId, groupId, userToInvite)
	invitationID := utils.NewUUID()
	query := `
	INSERT INTO group_requests (requestID, senderID, receiverID, groupID, typeRequest)
	VALUES (?,?,?,?,?)
	`
	stmt, err := gRepo.db.Prepare(query)
	if err != nil {
		return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v 1", err)}
	}
	defer stmt.Close()

	res, err := stmt.Exec(invitationID, userId, userToInvite, groupId, "invitation-request")
	if err != nil {
		return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v 1", err)}
	}
	if count, _ := res.RowsAffected(); count == 0 {
		return &models.ErrorJson{Status: 404, Error: "Invitation not found"}
	}

	return nil
}

func (gRepo *GroupRepository) CancelTheInvitation(userId, groupId string, invitedUser *models.User) *models.ErrorJson {
	fmt.Println("userId: ", userId, "groupId: ", groupId, "invited user: ", invitedUser.Id)
	query := `
	DELETE FROM group_requests WHERE 
	senderID = ? AND receiverID = ? AND groupID = ? AND typeRequest = ? 
	`
	stmt, err := gRepo.db.Prepare(query)
	if err != nil {
		return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v 1", err)}
	}
	defer stmt.Close()

	res, err := stmt.Exec(userId, invitedUser.Id, groupId, "invitation-request")
	if err != nil {
		return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v 1", err)}
	}
	// it must be here :)
	if count, _ := res.RowsAffected(); count == 0 {
		return &models.ErrorJson{Status: 404, Error: "Invitation not found"}
	}

	return nil
}

func (gRepo *GroupRepository) GetUsersToInvite(userID, groupID string) ([]models.User, *models.ErrorJson) {
	users := []models.User{}
	query := `
    WITH
    cte_followers AS (
        SELECT
            followerID as follower
        FROM
            followers
        WHERE
            userID = ?
    ),
    cte_invited AS (
        SELECT
            receiverID as invited
        FROM
            group_requests
        WHERE
            senderID = ?
            AND groupID = ?
            AND typeRequest = 'invitation-request'
            AND receiverID IN (
                SELECT
                    cte_followers.follower
                FROM
                    cte_followers
            )
    ),
    cte_not_invited AS (
        SELECT
            followerID as notInvited
        from
            followers
        WHERE
            userID = ?
            AND followerID NOT IN (
                SELECT
                    cte_invited.invited
                FROM
                    cte_invited
            )
    )
		SELECT
			cte_invited.invited,
			concat (users.firstName, " ", users.lastName) AS fullname,
			users.nickname,
			1 AS output
		FROM
			cte_invited
			INNER JOIN users ON users.userID = cte_invited.invited
		UNION
		SELECT
			cte_not_invited.notInvited,
			concat (users.firstName, " ", users.lastName) AS fullname,
			users.nickname,
			0 AS output
		FROM
			cte_not_invited
			INNER JOIN users ON users.userID = cte_not_invited.notInvited
		ORDER BY
			fullname DESC;

	`
	stmt, err := gRepo.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v 1", err)}
	}
	defer stmt.Close()

	rows, err := stmt.Query(userID, userID, groupID, userID)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v 2", err)}
	}
	// it must be here :)

	for rows.Next() {
		user := models.User{}
		if err := rows.Scan(&user.Id, &user.FullName, &user.Nickname, &user.Invited); err != nil {
			return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v 3", err)}
		}
		users = append(users, user)

	}

	return users, nil
}

// check this later
func (gRepo *GroupRepository) AlreadyInvited(groupID, userID string) *models.ErrorJson {
	var found int
	query := `SELECT 1 FROM groups WHERE groupID = ?`
	stmt, err := gRepo.db.Prepare(query)
	if err != nil {
		return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()
	if err = stmt.QueryRow(groupID).Scan(&found); err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	return nil
}
