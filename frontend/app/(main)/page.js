// frontend/app/(main)/page.js 
"use client"

let LoadPosts = async () => {
  // const getRequest = {
  //   method: "GET",
  //   credentials: "include"
  // }
  const postRequest = {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({
        Reciever_Id: "ec1ab391-7a83-4683-a86f-cbe3869140e7",
        Sender_Id: "ec1ab391-7a83-4683-a86f-cbe3869140e7",
        Type: "follow-public",
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