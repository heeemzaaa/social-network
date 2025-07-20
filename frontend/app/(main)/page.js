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
        Reciever_Id: "247c4572-09cb-49e8-b75f-353cf4068466",
        Sender_Id: "247c4572-09cb-49e8-b75f-353cf4068466",
        Type: "follow-private",
        Content: ""

      })
    }
  let response = await fetch("http://localhost:8080/api/notifications/", postRequest)
  let data = await response.json()
  console.log("Fetched post new notification:", data)
  let res = await fetch("http://localhost:8080/api/notifications/", getRequest)
  let ddd = await res.json()
  console.log("Fetched get seen notifications:", ddd)
  
}

export default function Home() {

  LoadPosts()
  return (
    <main className='home-page'>
      
    </main>
  );
}