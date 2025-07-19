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
        Notif_Id: "de6bd98b-7aab-4899-ae1d-e2f14555d30b",
        Status: "reject",
      })
    }
  let response = await fetch("http://localhost:8080/api/notifications/update/", postRequest)
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