// frontend/app/(main)/page.js 
"use client"

let LoadPosts = async () => {
  const getRequest = {
    method: "GET",
    credentials: "include"
  }
  const postRequest = {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({
        Sender_Id: "ffc54fb8-2d14-4f83-a196-062c976e3243",
        Reciever_Id: "ffc54fb8-2d14-4f83-a196-062c976e3243",
        Type: "group-join",
        Content: ""
      })
    }
  let response = await fetch("http://localhost:8080/api/notifications/", postRequest)
  let data = await response.json()
  console.log("Fetched notifications:", data)
}

export default function Home() {

  LoadPosts()
  return (
    <main className='home-page'>
      
    </main>
  );
}