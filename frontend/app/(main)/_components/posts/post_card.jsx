// import Button from '@/app/_components/button';
import Image from 'next/image';
import React from 'react';
import { HiEllipsisHorizontal, HiHeart, HiOutlineChatBubbleOvalLeft } from 'react-icons/hi2';
import Avatar from '../avatar';

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
    image: "/no-profile.png",
    timestamp: "2h ago",
    likes: 24,
    comments: 8,
    shares: 3,
    isLiked: true,
}

export default function PostCard({ post }) {

    return (
        <div className='post-card-container'>
            <div className="post-card-header">
                <Avatar/>
            </div>
            <div className="post-card-content">
                <div className="post-card-content-text">
                    {post.content}
                </div>
                <div className="post-card-content-img" style={{background:`url(${post.image})`}}/>
            </div>
            <div className="post-card-stats"></div>
            <div className="post-card-actions"></div>
        </div>
    );
}