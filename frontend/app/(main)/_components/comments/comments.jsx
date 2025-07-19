import React from 'react'
import CommentsCard from './commentsCard'
import "./comments.css"


export default function Comments() {
    let comments = [
        {
            "firstName": "Amina",
            "lastName": "El Idrissi",
            "userImage": "/no-profile.png",
            "content": "This is absolutely amazing! Great work.",
            "imagePath": "/no-profile.png",
            "createdAt": "2025-07-18T08:32:00Z",
            "likes": 12
        },
        {
            "firstName": "Youssef",
            "lastName": "Benali",
            "userImage": "/no-profile.png",
            "content": "Nice explanation, very clear and useful.",
            "createdAt": "2025-07-18T07:45:00Z",
            "likes": 5
        },
        {
            "firstName": "Salma",
            "lastName": "Meknassi",
            "userImage": "/no-profile.png",
            "content": "Where can I find more details about this?",
            "imagePath": "/no-profile.png",
            "createdAt": "2025-07-18T06:30:00Z"
        },
        {
            "firstName": "Hamza",
            "lastName": "Zouhair",
            "userImage": "/no-profile.png",
            "content": "Thanks for sharing this post.",
            "createdAt": "2025-07-18T05:12:00Z",
            "likes": 3
        },
        {
            "firstName": "Lina",
            "lastName": "Chakiri",
            "userImage": "/no-profile.png",
            "content": "Could you explain this part again?",
            "createdAt": "2025-07-18T04:50:00Z"
        },
        {
            "firstName": "Omar",
            "lastName": "Naji",
            "userImage": "/no-profile.png",
            "content": "I've been waiting for something like this!",
            "imagePath": "/no-profile.png",
            "createdAt": "2025-07-18T03:33:00Z",
            "likes": 9
        },
        {
            "firstName": "Sara",
            "lastName": "Kabbaj",
            "userImage": "/no-profile.png",
            "content": "Great insight, thanks for this.",
            "createdAt": "2025-07-18T02:20:00Z"
        }
    ]

    return (
        <section className='all_comments p2 gap-1 flex-col'>
            {comments.map((comment, index) => {
                return <CommentsCard key={index} comment={comment} />
            })}
        </section>
    )
}
