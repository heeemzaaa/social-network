import "./comments.css"
import React from 'react'
import CommentsCard from './commentsCard'


export default function Comments({ comments }) {
    return (
        <section className='all_comments p2 gap-1 flex-col'>
            {Array.isArray(comments) && comments.map((comment, index) => (
                <CommentsCard key={index} comment={comment} />
            ))}
        </section>

    )
}
