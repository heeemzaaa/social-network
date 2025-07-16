"use client"

import { memo, useEffect, useState } from "react";
import PostCard from "./postCard";

export function PostsContainer({ post }) {
    const [posts, setPosts] = useState([])

    useEffect(() => {
        console.log("fetch posts here.")
        // todo : for fetching posts
        try {
            
        } catch (error) {
            
        }
    },[])

    useEffect(()=>{
        if (!post) return;
        setPosts(prev => [post,...prev])
    },[post])

    return (
        <div className="posts-container">
            {posts?.map((post) => (
                <PostCard key={post.id} post={post} />
            ))}
        </div>
    );
}

export default memo(PostsContainer);
