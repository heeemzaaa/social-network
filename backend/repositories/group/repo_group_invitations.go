package group

import (
	"fmt"

	"social-network/backend/models"
)

func (gRepo *GroupRepository) InviteToJoin(userId, groupId string) *models.ErrorJson {
	query := `


	`
	fmt.Printf("query: %v\n", query)
	return nil
}

func (gRepo *GroupRepository) CancelTheInvitation(userId, groupId string) *models.ErrorJson {
	query := `
	
	
	`
	fmt.Printf("query: %v\n", query)
	return nil
}
