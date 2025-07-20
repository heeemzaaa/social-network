import React from 'react';
import Avatar from '../avatar';


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