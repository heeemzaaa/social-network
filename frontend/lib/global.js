let profileID = ""

function setID(id) {
    profileID = id
}

function getID() {
    return profileID
}

export {setID , getID} 

export const userList =  {
	username: String,
	userID: String,
	online: Boolean, 
}