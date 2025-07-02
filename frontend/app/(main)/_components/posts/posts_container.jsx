// frontend/app/_components/posts/posts_container.jsx

"use client"

import { memo } from "react";
import PostCard from "./post_card";

let postTest = {
    id: "1",
    type: "user",
    author: {
        name: "Sarah Johnson",
        username: "sarahj",
        avatar: "/placeholder.svg?height=40&width=40",
    },
    content:
        "Just finished an amazing hike in the mountains! The view from the top was absolutely breathtaking. Nature never fails to inspire me. 🏔️✨",
    image: "/placeholder.svg?height=200&width=300",
    timestamp: "2h ago",
    likes: 24,
    comments: 8,
    shares: 3,
    isLiked: true,
}

function PostsContainer({ posts }) {

    return (
        <div className="posts-container">
            {posts?.map((post) => (
                <PostCard key={post.id} post={postTest} />
            ))}
        </div>
    );
}

export default memo(PostsContainer);