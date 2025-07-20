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
    group_posts (postID, groupID, userID, content, imagePath)
    VALUES (?, ?, ?, ?, ?) RETURNING postID,
    groupID,
    userID,
    content,
    imagePath,
    createdAt,
    (
        SELECT
            concat (firstName, ' ', lastName)
        FROM
            users
        WHERE
            users.userID = userID
    ) AS fullName,
    (
        SELECT
            nickname
        FROM
            users
        WHERE
            users.userID = userID
    );

	`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()
	groupCreated := models.Group{}
	err = stmt.QueryRow(groupID, group.GroupCreatorId,
		group.Title, group.ImagePath, group.Description).Scan(
		&groupCreated.GroupId,
		&groupCreated.GroupCreatorId,
		&groupCreated.Title,
		&groupCreated.ImagePath,
		&groupCreated.Description,
		&groupCreated.CreatedAt,
		&groupCreated.GroupCreatorFullName,
		&groupCreated.GroupCreatorNickname,
	)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	if errJson := repo.JoinGroup(&groupCreated, group.GroupCreatorId); errJson != nil {
		return nil, &models.ErrorJson{Status: errJson.Status, Message: errJson.Message, Error: errJson.Error}
	}
	return &groupCreated, nil
}

// all these functions are needed to handle the user
func (repo *GroupRepository) GetJoinedGroups(offset int64, userID string) ([]models.Group, *models.ErrorJson) {
	joinedGroups := []models.Group{}
	query := `
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
	WHERE groups.groupCreatorID != ?
	AND  group_membership.userID = ?
    ORDER BY groups.createdAt DESC
	LIMIT
		20
	OFFSET
		?
	`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	rows, err := stmt.Query(userID, userID, offset)
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
			&group.GroupCreatorFullName,
			&group.GroupCreatorNickname,
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

func (repo *GroupRepository) GetAvailableGroups(offset int64, userID string) ([]models.Group, *models.ErrorJson) {
	availabeGroups := []models.Group{}
	query := `
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
		groups.groupCreatorID,
		concat(users.firstName, ' ', users.lastName),
		users.nickname,
		title,
		imagePath,
		description,
		groups.createdAt ,
		cte_members.Nbr_Members
	FROM
		groups
		INNER JOIN cte_members ON groups.groupID = cte_members.Id
		INNER JOIN users ON users.userID = groups.groupCreatorID
	WHERE
		groups.groupID NOT IN (
			SELECT
				groups.groupID
			FROM
				groups
				INNER JOIN group_membership ON group_membership.groupID = groups.groupID
				AND group_membership.userID = ?
		)
	ORDER BY groups.createdAt DESC
	LIMIT
		20
	OFFSET
		?
	`
	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	rows, errQuery := stmt.Query(userID, offset)
	if errQuery != nil {
		if errQuery == sql.ErrNoRows {
			return availabeGroups, nil
		}
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", errQuery)}
	}
	defer rows.Close()
	for rows.Next() {
		group := &models.Group{}
		errScan := rows.Scan(
			&group.GroupId,
			&group.GroupCreatorId,
			&group.GroupCreatorFullName,
			&group.GroupCreatorNickname,
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

func (repo *GroupRepository) GetCreatedGroups(offset int64, userID string) ([]models.Group, *models.ErrorJson) {
	createdGroups := []models.Group{}
	query := `
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
		groupCreatorID = ?
	ORDER BY
		groups.createdAt DESC
	LIMIT
		20
	OFFSET
		?

	`
	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	rows, errQuery := stmt.Query(userID, offset)
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
