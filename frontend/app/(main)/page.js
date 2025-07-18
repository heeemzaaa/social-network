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
        Sender_Id: "67e3fbfb-e5f9-4eb1-ab7e-c738f69d3580",
        Reciever_Id: "67e3fbfb-e5f9-4eb1-ab7e-c738f69d3580",
        Type: "follow-private",
        Content: "Hello, please accept."
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