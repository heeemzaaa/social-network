export async function fetchNotif() {
   
    try {
        // const postRequest = {
        //     method :"POST",
        //     credentials : "include",
        //     body : JSON.stringify({Type: "follow-private", Sender_Id: "562b7f42-b132-4a5d-8863-ed111c6487eb", Reciever_Id: "946257db-65e0-491e-9a09-c33f946f5092"})
        // }

        const getRequest = {
            method :"GET",
            credentials : "include"
        }
        
        const response = await fetch("http://localhost:8080/api/notification", getRequest)
        
        console.log("inside fetchNotif: ", response)
        const data = await response.json()
        return data
    } catch (error) {
        console.error("Error in API single notifications:", error.message);
    }
}