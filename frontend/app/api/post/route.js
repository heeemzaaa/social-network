// frontend/app/api/post/route.js
import { NextResponse } from "next/server";

const posts = [
    {
        "id": "post_001",
        "userId": "user_123",
        "groupId": "group_456",
        "content": "Just tried the new VR headset, and it's mind-blowing! Anyone else excited about virtual reality?",
        "createdAt": "2025-06-29T10:15:30Z",
        "updatedAt": "2025-06-29T10:20:00Z",
        "likes": 150,
        "likedByUser": true,
        "commentsCount": 1,
    },
    {
        "id": "post_002",
        "userId": "user_789",
        "groupId": "group_123",
        "content": "Reading 'The Great Gatsby' again. This book never gets old!",
        "createdAt": "2025-06-28T09:00:00Z",
        "updatedAt": "2025-06-28T09:00:00Z",
        "likes": 85,
        "likedByUser": false,
        "commentsCount": 0,
    }
]


// To handle a GET request to /api
export async function getPosts(request) {
    return Response.json(posts, { status: 200 });
}

// To handle a POST request to /api
export async function Post(request) {
    return NextResponse.json({ message: "Hello World" }, { status: 200 });
}