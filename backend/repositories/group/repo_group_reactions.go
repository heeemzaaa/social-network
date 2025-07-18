package group

import (
	"database/sql"
	"fmt"

	"social-network/backend/models"
	"social-network/backend/utils"
)

func (appRepo *GroupRepository) AddReaction(reaction *models.GroupReaction, type_reaction int) (*models.GroupReaction, *models.ErrorJson) {
	reaction_created := &models.GroupReaction{}
	reactionID := utils.NewUUID()
	query := `INSERT INTO group_reactions 
	(reactionID , entityType, entityID,reaction,userID) VALUES (?,?,?,?,?) RETURNING reaction`
	stmt, err := appRepo.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v 1", err)}
	}
	defer stmt.Close()
	err = stmt.QueryRow(reactionID, reaction.EntityType, reaction.EntityId, type_reaction, reaction.UserId).Scan(
		&reaction_created.Reaction)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v 2", err)}
	}
	return reaction_created, nil
}

func (appRepo *GroupRepository) UpdateReactionLike(reaction *models.GroupReaction) (*models.GroupReaction, *models.ErrorJson) {
	reaction_created := &models.GroupReaction{}
	query := `UPDATE group_reactions SET reaction = CASE reaction
              WHEN 0 THEN 1
			  WHEN -1 THEN 1
              ELSE 0
              END
	          WHERE reactionID = ? 
			  RETURNING reaction;`

	err := appRepo.db.QueryRow(query, reaction.Id).Scan(&reaction_created.Reaction)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v ", err)}
	}

	return reaction_created, nil
}

func (appRepo *GroupRepository) HanldeReaction(reaction *models.GroupReaction) (*models.GroupReaction, *models.ErrorJson) {
	reaction_existed := &models.GroupReaction{}
	query := `SELECT * FROM group_reactions WHERE
	 userID = ? AND entityType = ? AND entityID = ?
	 `
	stmt, err := appRepo.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v jj", err)}
	}
	if err := stmt.QueryRow(reaction.UserId,
		reaction.EntityType, reaction.EntityId).Scan(
		&reaction_existed.Id,
		&reaction_existed.EntityType,
		&reaction_existed.EntityId,
		&reaction_existed.Reaction,
		&reaction_existed.UserId); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v jjj", err)}
	}
	return reaction_existed, nil
}



