import Button from '@/app/_components/button';
import Image from 'next/image';
import React from 'react';
import { HiEllipsisHorizontal, HiHeart, HiOutlineChatBubbleOvalLeft } from 'react-icons/hi2';
import Avatar from '../avatar';


//TODO:ADD REAL PATH TO THE IMG SRC 
export default function PostCard({ post }) {
    return (
        <div className='post-card-container' >
            <div className="post-card-header">
                <Avatar />
            </div>
            <div className="post-card-content">
                 <div className="post-card-content-privacy">
                    {post.privacy}
                </div>
                <div className="post-card-content-privacy">
                   <img src="" alt="postImage" />
                </div>
                <div className="post-card-content-text">
                    {post.content}
                </div>
                <div className="post-card-content-img" style={{ background: `url(${post.image})` }} />
            </div>
            <div className="post-card-stats"></div>
            <div className="post-card-actions"></div>
        </div>
    );
}