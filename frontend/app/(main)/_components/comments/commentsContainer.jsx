import React from 'react'
import "./comments.css"
import Comments from './comments'
import CommentsFooter from './commentsFooter'


export default function CommentsContainer() {
  return (
    <section className='comments_container w-full flex-col gap-2 '>
        <Comments />
        <CommentsFooter />
    </section>
  )
}
