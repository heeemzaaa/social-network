import Button from '@/app/_components/button';
import Image from 'next/image';
import React from 'react';
import { useActionState } from 'react'

import { HiEllipsisHorizontal, HiHeart, HiOutlineChatBubbleOvalLeft } from 'react-icons/hi2';
import Avatar from '../avatar';
import { likePostAction } from '@/app/_actions/posts';


//TODO:ADD REAL PATH TO THE IMG SRC 
export default function PostCard({ post }) {
     const [state, likeAction] = useActionState(likePostAction, { message: '' })
    return (
        <div className='post-card-container' >
            <div className="post-card-header">
                <Avatar />
            </div>
            <div className="post-card-content-nickname">
                {post.user.nickname}
            </div>
            <div className="post-card-content-likes">
                {post.total_likes}
            </div>
            <div className="post-card-content">
                <div className="post-card-content-privacy">
                    {post.privacy}
                </div>
                <form action={likeAction}>
                    <input type="hidden" name="postId" value={post.id} />
                    <button type="submit">Like</button>
                    {state.message && <p>{state.message}</p>}
                </form>
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