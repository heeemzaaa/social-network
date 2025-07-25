package group

import (
	"fmt"

	"social-network/backend/models"
)

// function to check if a specific item is there based on a specific value
// generic somehow
// we need to specify  the type aftewards ;)
// it will be used for the nickname , session and also the email checking
func (gRepo *GroupRepository) GetItem(typ string, field string, value string) ([]any, bool, *models.ErrorJson) {
	data := make([]any, 0)
	query := fmt.Sprintf(`SELECT %v FROM %v WHERE %v=?`, field, typ, field)
	stmt, err := gRepo.db.Prepare(query)
	if err != nil {
		return nil, false, &models.ErrorJson{Status: 500, Error: "ERROR!! Internal Server error"}
	}
	defer stmt.Close()
	rows, err := stmt.Query(value)
	if err != nil {
		return nil, false, &models.ErrorJson{Status: 500, Error: "ERROR!! Internal Server error"}
	}
	for rows.Next() {
		var row any
		rows.Scan(&row)
		data = append(data, row)
	}

	defer rows.Close()

	if len(data) != 0 {
		return data, true, nil
	}
	return nil, false, nil
}

func (gRepo *GroupRepository) IsAdmin(groupId, userId string) (bool, *models.ErrorJson) {
	query := `SELECT groups.groupCreatorID FROM groups WHERE groups.groupID = ?`
	stmt, err := gRepo.db.Prepare(query)
	if err != nil {
		return false, &models.ErrorJson{Status: 500, Error: "ERROR!! Internal Server error"}
	}
	defer stmt.Close()
	var groupAdminId string
	err = stmt.QueryRow(groupId).Scan(&groupAdminId)
	if err != nil {
		return false, &models.ErrorJson{Status: 500, Error: "ERROR!! Internal Server error"}
	}

	return groupAdminId == userId, nil
}
