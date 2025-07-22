"use client";
import { useEffect, useState } from "react";
import PostCard from "./postCard";

export default function PostCardList({post}) {
    const [posts, setPosts] = useState([])
    
        useEffect(() => {
            async function fetchPosts() {
                console.log("fetch posts here.");
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
                    console.log(data)
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
        <div className="list-container " style={{ overflowY: "auto" }}>
            {posts?.map((post) => (
                <PostCard key={post.id} {...post} />
            ))}
        </div>
    );
}
