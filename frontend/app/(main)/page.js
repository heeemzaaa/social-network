"use client"

import { useEffect, useState } from "react";
import { useModal } from "./_context/ModalContext";
import PostCardList from "./_components/posts/postCardList";



export default function Home() {

  const [post, setPost] = useState(null) 
  const {getModalData, setModalData} = useModal()
  useEffect(()=> {
    let postData = getModalData()  
    setPost(postData)
  },[setModalData])
  return (
    <main className='home-page' style={{overflow : "auto"}} >
      <PostCardList post={post} />
    </main>
  );
}