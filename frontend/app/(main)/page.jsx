"use client"

import PostCardList from "./_components/posts/postCardList";
export default function Home() {
  return (
    <main className='home-page' style={{overflow : "auto"}} >
      <PostCardList />
    </main>
  );
}