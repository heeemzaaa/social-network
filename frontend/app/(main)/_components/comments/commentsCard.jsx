'use client'
import React, { useState } from 'react'
import "./comments.css"

export default function CommentsCard({ comment }) {
    let parsedTime = new Date(comment.createdAt)
    console.log("comment",comment)
    if(comment.lastName==undefined) {
        comment.lastName = ""
    }
    return (
        <div className='comments_card w-full p3 flex-col gap-3 shadow-lg'>
            <div className='card_header flex align-center gap-1'>
                <img src={`http://localhost:8080/static/${comment.userImage}` || "/no-profile.png"} className='user_image' />
                <span className='user_name'>{comment.firstName + " " + comment.lastName}</span>
            </div>

            <div className='card_body flex justify-center align-center gap-1'>
                {comment.imagePath && <img src={`http://localhost:8080/static/${comment.imagePath}`}  className='comments_image' />}
                <h4 className='comment'>{comment.content}</h4>
            </div>

            <div className='card_footer flex justify-between'>
                <p className='time text-md'>{parsedTime.toLocaleString()}</p>
            </div>
        </div>
    )
}
