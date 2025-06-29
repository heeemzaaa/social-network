"use client"

import { useEffect } from "react";
import InfosDiv from "./_components/user_info";
import { getPosts } from "../api/post/route";

export default function Home() {
  const userInfos = {
    firstName: "Hamza",
    lastName: "Elkhawlani",
    email: "hamza@gmail.com",
    dateOfBirth: "20-09-2000",
    nickname: "heeemzaaa",
    followers: 153,
    following: 147,
    posts: 52,
    groups: 18
  }

  // useEffect( async ()=>{ 
  //   let data = await getPosts()
  //   console.log(data)
  // },[]) 

  return (
    <main className="home_page_section">
      <InfosDiv {...userInfos} />
    </main>
  );
}


