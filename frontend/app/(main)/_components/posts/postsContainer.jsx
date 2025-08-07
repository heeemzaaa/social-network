"use client"

import { memo, useEffect, useState } from "react";
import PostCard from "./postCard";

export function PostsContainer({ post }) {
    const [posts, setPosts] = useState([])

    useEffect(() => {
        async function fetchPosts() {
            try {
                const resp = await fetch("http://localhost:8080/api/posts", {
                    method: "GET",
                    credentials: "include",
                });

                if (!resp.ok) {
                    console.log("error fetching posts 1");
                    return;
                }
                const data = await resp.json();
                setPosts(data); 
            } catch (error) {
                console.log("error fetching posts", error);
            }
        }

        fetchPosts(); 
    }, []);

    useEffect(() => {
        if (!post) return;
        setPosts(prev => [post, ...prev])
    }, [post])

    return (
        <div className="posts-container">
            {posts?.map((post) => (
                <PostCard key={post.id} post={post} />
            ))}
        </div>
    );
}

export default memo(PostsContainer);
