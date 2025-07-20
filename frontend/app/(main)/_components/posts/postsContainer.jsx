"use client"

import { memo } from "react";
import PostCard from "./postCard";

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