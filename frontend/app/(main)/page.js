// frontend/app/(main)/page.js 

"use client"
import InfosDiv from "./_components/user_info";
import PostsContainer from "./_components/posts/posts_container";
import { useEffect, useState } from "react";
import Loading from "./loading";
import { fetchPosts, fetchNotif } from "./_lib/posts";

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

  const [posts, setPosts] = useState(null)
  const [loading, setLoading] = useState(true)

  // using the golang api directely
  // useEffect( () => {
  //   console.log("use Effect runs");
  //   let LoadPosts = async () => {
  //     let response = await fetch("http://localhost:8080/api/posts")
  //     let data = await response.json()
  //     setPosts(data)
  //     setLoading(false)
  //   }
  //   LoadPosts()
  // }, []);



  // useEffect( () => {
  //   console.log("use Effect runs");
  //   let LoadPosts = async () => {
  //     let posts = await fetchPosts() 
  //     console.log(posts)
  //     setPosts()
  //     setLoading(false)
  //   }
  //   LoadPosts()
  // }, []);

  //  using the golang api directely
 useEffect(() => {
  const getRequest = {
    method: "GET",
    credentials: "include"
  }
  let LoadPosts = async () => {
    let response = await fetch("http://localhost:8080/api/notification", getRequest)
    let data = await response.json()
    console.log("Fetched notifications:", data)
    setPosts(data)
    setLoading(false)
  }
  LoadPosts()
}, []);
  
  // useEffect( () => {
  //   console.log("use Effect runs");
  //   let LoadPosts = async () => {
  //     let posts = await fetchNotif() 
  //     console.log(posts)
  //   }
  //   LoadPosts(false)
  // }, []);

  return (
    <main className='home-page'>
      <InfosDiv {...userInfos} />
      {loading ? <Loading /> :  <PostsContainer posts={posts} />}
    </main>
  );
}