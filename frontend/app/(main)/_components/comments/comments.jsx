import React from 'react'
import CommentsCard from './commentsCard'
import "./comments.css"


export default function Comments({ comments, id, groupID, creatorID }) {
    return (
        <section className='all_comments p2 gap-1 flex-col'>
            {Array.isArray(comments) && comments.map((comment, index) => (
                <CommentsCard key={index} comment={comment} id={id} groupID={groupID} creatorID={creatorID} />
            ))}
        </section>

    )
}
