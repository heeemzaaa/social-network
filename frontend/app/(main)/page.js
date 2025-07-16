// frontend/app/(main)/page.js 
"use client"

import { useEffect, useState } from "react";
import PostsContainer from "./_components/posts/postsContainer";
import { useModal } from "./_context/ModalContext";



export default function Home() {

  const [post, setPost] = useState(null) 
  const {getModalData, setModalData} = useModal()
  useEffect(()=> {
    let postData = getModalData()  
    setPost(postData)
  },[setModalData])
  return (
    <main className='home-page'>
      <PostsContainer post={post} />
    </main>
  );
}