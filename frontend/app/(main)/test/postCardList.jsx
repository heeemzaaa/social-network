"use client";
import PostCard from "./postCard";

export default function PostCardList() {
    const posts = [
        {
            id: "1",
            user: {
                first_name: "Alice",
                last_name: "Johnson",
                avatar_path: "/avatar1.png",
            },
            content: "Just had a great day at the beach!",
            created_at: new Date().toISOString(),
            img: "/no-profile.png",
            total_likes: 12,
            total_comments: 3,
            privacy: "public",
            liked: 1,
        },
        {
            id: "2",
            user: {
                first_name: "Bob",
                last_name: "Smith",
                avatar_path: "/avatar2.png",
            },
            content: "Reading a fascinating book about space exploration.",
            created_at: new Date().toISOString(),
            img: "/no-profile.png",
            total_likes: 8,
            total_comments: 0,
            privacy: "private",
            liked: 0,
        },
        {
            id: "3",
            user: {
                first_name: "Claire",
                last_name: "Lee",
                avatar_path: "/avatar3.png",
            },
            content: "Here’s my latest painting. Hope you like it!",
            created_at: new Date().toISOString(),
            total_likes: 20,
            total_comments: 5,
            privacy: "public",
            liked: 1,
        },
    ];

    return (
        <div className="list-container " style={{ overflowY: "auto" }}>
            {posts.map((post) => (
                <PostCard key={post.id} {...post} />
            ))}
        </div>
    );
}
