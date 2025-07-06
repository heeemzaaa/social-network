package group

import (
	"database/sql"
	"fmt"

	"social-network/backend/models"

	"github.com/google/uuid"
)

type GroupRepository struct {
	db *sql.DB
}

// NewPostRepository creates a new repository
func NewGroupRepository(db *sql.DB) *GroupRepository {
	return &GroupRepository{db: db}
}

func (repo *GroupRepository) CreateGroup(group *models.Group) (*models.Group, *models.ErrorJson) {
	groupID := uuid.New().String()
	query := `INSERT INTO groups (groupID, groupCreatorID,title,imagePath,description)
	VALUES (?,?,?,?,?)`
	_, err := repo.db.Exec(query, groupID, group.GroupCreatorId, group.Title, group.ImagePath, group.Description)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}

	return group, nil
}

// all these functions are needed to handle the user
func (repo *GroupRepository) GetJoinedGroups(offset int64, userID string) ([]models.Group, *models.ErrorJson) {
	joinedGroups := []models.Group{}
	query := `
	SELECT title , imagePath, description 
	FROM groups INNER JOIN group_membership 
    ON group_membership.groupID = groups.groupID
	AND groups.groupCreatorID != ?
    LIMIT 20 OFFSET ?
	`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	rows, errQuery := stmt.Query(userID, offset)
	if errQuery != nil {
		if err == sql.ErrNoRows {
			return joinedGroups, nil
		}
	}

	for rows.Next() {
		var group models.Group
		errScan := rows.Scan(&group.Title, &group.ImagePath, &group.Description)
		if errScan != nil {
			return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
		}
		joinedGroups = append(joinedGroups, group)
	}

	return joinedGroups, nil
}

func (repo *GroupRepository) GetAvailableGroups(offset int64, userID string) ([]models.Group, *models.ErrorJson) {
	availabeGroups := []models.Group{}
	query := `
	SELECT title , imagePath, description 
	FROM groups INNER JOIN group_membership 
    ON group_membership.groupID != groups.groupID
	AND groups.groupCreatorID != ?
    LIMIT 20 OFFSET ?
	`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	rows, errQuery := stmt.Query(userID, offset)
	if errQuery != nil {
		if err == sql.ErrNoRows {
			return availabeGroups, nil
		}
	}

	for rows.Next() {
		var group models.Group
		errScan := rows.Scan(&group.Title, &group.ImagePath, &group.Description)
		if errScan != nil {
			return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
		}
		availabeGroups = append(availabeGroups, group)
	}

	return availabeGroups, nil
}

func (repo *GroupRepository) GetCreatedGroups(offset int64, userID string) ([]models.Group, *models.ErrorJson) {
	createdGroups := []models.Group{}
	query := `
    SELECT title , imagePath, description 
	FROM groups 
	WHERE groupCreatorID = ?
	LIMIT 20 OFFSET ?
	`
	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	rows, errQuery := stmt.Query(userID, offset)
	if errQuery != nil {
		if err == sql.ErrNoRows {
			return createdGroups, nil
		}
	}

	for rows.Next() {
		var group models.Group
		errScan := rows.Scan(&group.Title, &group.ImagePath, &group.Description)
		if errScan != nil {
			return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
		}
		createdGroups = append(createdGroups, group)
	}

	return createdGroups, nil
}

func (repo *GroupRepository) GetGroupById() {
}
