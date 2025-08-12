"use client"

import { useRouter } from "next/navigation";
import PostCard from "./postCard";
import {
    memo,
    useEffect,
    useState
} from "react";


const API_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080'

export function PostsContainer({ post }) {
    const [posts, setPosts] = useState([])
    const router = useRouter()

    useEffect(() => {
        async function fetchPosts() {
            try {
                const resp = await fetch(`${API_URL}/api/posts`, {
                    method: "GET",
                    credentials: "include",
                });

                if (!resp.ok) {
                    if (resp.status === 401) {
                        router.push("/login")
                        return
                    }
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
