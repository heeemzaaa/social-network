"use client"

import { HiOutlineDocumentPlus } from "react-icons/hi2";
import Button from "../_components/button";
import PostCardList from "./_components/posts/postCardList";
import { useModal } from "./_context/ModalContext";
import CreatePost from "./_components/posts/createPost";
import { createPostAction } from "../_actions/posts";
export default function Home() {
  const { openModal } = useModal()
  return (
    <main className='home-page flex-col' >
      <div style={{position:"sticky", top:"0", borderBottom:"solid 1px", paddingBottom:".5rem", margin:".5rem"}} >
        <Button style={{marginLeft:"auto"}} onClick={() => openModal(<CreatePost postAction={createPostAction} />)}>
          <HiOutlineDocumentPlus size={24} />
          <span>Add New Post</span>
        </Button>
      </div>
      <PostCardList />
    </main >
  );
}