package group

import (
	"database/sql"
	"fmt"

	"social-network/backend/models"
	"social-network/backend/utils"
)

// let's edit all this shit to return the data needed

type GroupRepository struct {
	db *sql.DB
}

// NewPostRepository creates a new repository
func NewGroupRepository(db *sql.DB) *GroupRepository {
	return &GroupRepository{db: db}
}

func (repo *GroupRepository) CreateGroup(group *models.Group) (*models.Group, *models.ErrorJson) {
	groupID := utils.NewUUID()
	query := `
	INSERT INTO
    groups (groupID, groupCreatorID, title, imagePath , description )
    VALUES (?, ?, ?, ?, ?) RETURNING 
	groupID,
    groupCreatorID,
    title,
    imagePath,
    description,
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
    )
	;

	`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v 2", err)}
	}
	defer stmt.Close()

	groupCreated := models.Group{}
	err = stmt.QueryRow(groupID, group.GroupCreatorId,
		group.Title, group.ImagePath, group.Description, group.GroupCreatorId, group.GroupCreatorId, group.GroupCreatorId).Scan(
		&groupCreated.GroupId,
		&groupCreated.GroupCreatorId,
		&groupCreated.Title,
		&groupCreated.ImagePath,
		&groupCreated.Description,
		&groupCreated.CreatedAt,
		&groupCreated.GroupCreator.FullName,
		&groupCreated.GroupCreator.Nickname,
		&groupCreated.GroupCreator.ImagePath,
	)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v 1", err)}
	}
	if errJson := repo.JoinGroup(&groupCreated, group.GroupCreatorId); errJson != nil {
		return nil, &models.ErrorJson{Status: errJson.Status, Message: errJson.Message, Error: errJson.Error}
	}
	return &groupCreated, nil
}

// all these functions are needed to handle the user
func (repo *GroupRepository) GetJoinedGroups(offset string, userID string) ([]models.Group, *models.ErrorJson) {
	joinedGroups := []models.Group{}
	var where string
	if offset == "0" {
		where = "groups.groupCreatorID != ? AND  group_membership.userID = ?"
	} else {
		where = `groups.groupCreatorID != ? AND  group_membership.userID = ? AND createdAt < (
			select createdAt from groups WHERE groupID = ? 
		)`
	}
	query := fmt.Sprintf(`
	WITH
    cte_members AS (
        SELECT
            group_membership.groupID AS Id,
            count(group_membership.groupID) AS Nbr_Members
        FROM
            users
            INNER JOIN group_membership ON users.userID = group_membership.userID
        GROUP BY
            Id
    )
	SELECT
	DISTINCT
		groups.groupID,
		groups.groupCreatorID,
		concat(users.firstName, ' ', users.lastName),
		users.nickname,
		title,
		imagePath,
		description,
		groups.createdAt,
		cte_members.Nbr_Members
	FROM
		groups
		INNER JOIN group_membership ON group_membership.groupID = groups.groupID
		INNER JOIN cte_members ON  cte_members.Id = groups.groupID
		INNER JOIN users ON users.userID = groups.groupCreatorID
	WHERE %v
	
    ORDER BY groups.createdAt DESC
	LIMIT
		6
	`, where)

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	args := []any{userID, userID}
	if offset != "0" {
		args = append(args, offset)
	}
	rows, err := stmt.Query(args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return joinedGroups, nil
		}
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer rows.Close()
	for rows.Next() {
		group := &models.Group{}
		errScan := rows.Scan(
			&group.GroupId,
			&group.GroupCreatorId,
			&group.GroupCreator.FullName,
			&group.GroupCreator.Nickname,
			&group.Title,
			&group.ImagePath,
			&group.Description,
			&group.CreatedAt,
			&group.Total_Members)
		if errScan != nil {
			return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", errScan)}
		}
		joinedGroups = append(joinedGroups, *group)
	}

	return joinedGroups, nil
}

func (repo *GroupRepository) GetAvailableGroups(offset string, userID string) ([]models.Group, *models.ErrorJson) {
	availabeGroups := []models.Group{}

	query := `
	WITH
    cte_members AS (
        SELECT
            gm.groupID,
            COUNT(gm.groupID) AS Nbr_Members
        FROM
            group_membership gm
            INNER JOIN cte_groups cg ON gm.groupID = cg.groupID
        GROUP BY
            gm.groupID
    ),
    cte_requested AS (
        SELECT
            groupID AS groupIdRequested
        FROM
            group_requests
        WHERE
            senderID = ?
            AND typeRequest = 'join-request'
    ),
    cte_available AS (
        SELECT 
            g.groupID AS available
        FROM
            groups g
        WHERE
            g.groupID NOT IN (
                SELECT 
                    groupID
                FROM
                    group_membership
                WHERE
                    userID = ?

                UNION ALL
                SELECT 
                    groupIdRequested
                FROM
                    cte_requested
            )
    ),
    cte_groups AS (
        SELECT
            groupIdRequested AS groupID,
            1 AS output
        FROM
            cte_requested
        UNION ALL
        SELECT
            available AS groupID,
            0 AS output
        FROM
            cte_available
    )
	SELECT
		cg.groupID,
		cg.output,
		g.groupCreatorID,
		CONCAT(u.firstName, ' ', u.lastName) AS creatorName,
		u.nickname,
		g.title,
		g.imagePath,
		g.description,
		g.createdAt,
		COALESCE(cm.Nbr_Members, 0) AS Nbr_Members
	FROM
		cte_groups cg
		INNER JOIN groups g ON g.groupID = cg.groupID
		INNER JOIN users u ON u.userID = g.groupCreatorID
		LEFT JOIN cte_members cm ON cm.groupID = cg.groupID
	ORDER BY
		g.createdAt DESC
	LIMIT
		6;
	`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	rows, errQuery := stmt.Query(userID, userID)
	if errQuery != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("===> %v", errQuery)}
	}
	defer rows.Close()
	for rows.Next() {
		group := &models.Group{}
		errScan := rows.Scan(
			&group.GroupId,
			&group.Requested,
			&group.GroupCreatorId,
			&group.GroupCreator.FullName,
			&group.GroupCreator.Nickname,
			&group.Title,
			&group.ImagePath,
			&group.Description,
			&group.CreatedAt,
			&group.Total_Members)
		if errScan != nil {
			return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", errScan)}
		}
		availabeGroups = append(availabeGroups, *group)
	}

	return availabeGroups, nil
}

// here we won't be needing to get the name

func (repo *GroupRepository) GetCreatedGroups(offset string, userID string) ([]models.Group, *models.ErrorJson) {
	var where string
	if offset == "0" {
		where = " groupCreatorID = ?"
	} else {
		where = ` groupCreatorID = ? AND createdAt < (
			select createdAt from groups WHERE groupID = ? 
		)`
	}
	createdGroups := []models.Group{}
	query := fmt.Sprintf(`
   	WITH
    cte_members AS (
        SELECT
            group_membership.groupID AS Id,
            count(group_membership.groupID) AS Nbr_Members
        FROM
            users
            INNER JOIN group_membership ON users.userID = group_membership.userID
        GROUP BY
            Id
    )
	SELECT
		groupID,
		title,
		imagePath,
		description,
		createdAt,
		cte_members.Nbr_Members
	FROM
		groups
		INNER JOIN cte_members ON groups.groupID = cte_members.Id
	WHERE
		%v  
	ORDER BY
		groups.createdAt DESC
	LIMIT
		6
	`, where)
	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	args := []any{userID}
	if offset != "0" {
		args = append(args, offset)
	}
	rows, errQuery := stmt.Query(args...)
	if errQuery != nil {
		if errQuery == sql.ErrNoRows {
			return createdGroups, nil
		}
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", errQuery)}
	}
	defer rows.Close()

	for rows.Next() {
		group := &models.Group{}
		errScan := rows.Scan(
			&group.GroupId,
			&group.Title,
			&group.ImagePath,
			&group.Description,
			&group.CreatedAt,
			&group.Total_Members,
		)
		if errScan != nil {
			return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", errScan)}
		}
		createdGroups = append(createdGroups, *group)
	}
	return createdGroups, nil
}

func (repo *GroupRepository) GetGroupById(groupID string) *models.ErrorJson {
	var found int
	query := `SELECT 1 FROM groups WHERE groupID = ?`
	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()
	
	if err = stmt.QueryRow(groupID).Scan(&found); err != nil {
		if err == sql.ErrNoRows {
			return &models.ErrorJson{Status: 404, Error: "ERROR!! Group Not Found!"}
		}
		return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	return nil
}
