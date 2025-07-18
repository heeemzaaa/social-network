package group

import "social-network/backend/utils"

//  This is will be for design and then implementing the requests in the correct way
//  admin of the group has the right to invite its followers
// BUT we also add the admin as a memeber of the group
// so the whole list of members will be able to invite its followers


func (gRepo *GroupRepository) RequestToJoin(userId, groupId string){
	requestId := utils.NewUUID()
	query := `
	INSERT INTO (requestID,senderID,receiverID,groupID)
	`

}



func (gRepo *GroupRepository) RequestToCancel(userId, groupId string){
	query:= `
	
	
	`

}