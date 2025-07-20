'use client'
import React, { useState } from 'react'
import "./comments.css"
import { FaRegHeart } from "react-icons/fa6";
import { FaHeart } from "react-icons/fa";


export default function CommentsCard({ comment }) {
    let parsedTime = new Date(comment.createdAt)
    console.log("comment",comment)
    if(comment.lastName==undefined) {
        comment.lastName = ""
    }
    return (
        <div className='comments_card w-full p3 flex-col gap-3 shadow-lg'>
            <div className='card_header flex align-center gap-1'>
                <img src={comment.userImage || "/no-profile.png"} className='user_image' />
                <span className='user_name'>{comment.firstName + " " + comment.lastName}</span>
            </div>

            <div className='card_body flex justify-center align-center gap-1'>
                {comment.imagePath && <img src={comment.imagePath} className='comments_image' />}
                <h4 className='comment'>{comment.content}</h4>
            </div>

            <div className='card_footer flex justify-between'>
                {comment.isLiked === true ?
                    <div className='flex align-center justify-center gap-1'>
                        <FaHeart size={'24px'} color='red' />
                        <span>{comment.likes}</span>
                    </div> :
                    <div className='flex align-center justify-center gap-1'>
                        <FaRegHeart size={'24px'} />
                        <span>{comment.likes}</span>
                    </div>
                }
                <p className='time text-md'>{parsedTime.toLocaleString()}</p>
            </div>
        </div>
    )
}
