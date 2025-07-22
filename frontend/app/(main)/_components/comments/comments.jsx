import React from 'react'
import CommentsCard from './commentsCard'
import "./comments.css"


export default function Comments({ comments }) {
    return (
        <section className='all_comments p2 gap-1 flex-col'>
            {comments.map((comment, index) => {
                return <CommentsCard key={index} comment={comment} />
            })}
        </section>
    )
}
