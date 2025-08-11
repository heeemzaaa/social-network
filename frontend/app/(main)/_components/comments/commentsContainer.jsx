"use client"

import "./comments.css"
import Comments from './comments'
import CommentsFooter from './commentsFooter'
import { useEffect, useState } from 'react'
import { useNotification } from "../../_context/NotificationContext"
import Loader from "../loader"
import Error from "../error"

const API_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080'

export default function CommentsContainer({ id, onCommentMessage, groupID, creatorID }) {
  const [comments, setComments] = useState([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState(null)
  const postComment = `http://localhost:8080/api/posts/comments/${id}`
  const groupComment = `http://localhost:8080/api/groups/${groupID}/posts/${id}/comments?offset=0`

  useEffect(() => {
    const fetchComments = async () => {
      try {
        setLoading(true)
        setError(null)
        const res = await fetch(groupID ? groupComment : postComment, {
          method: 'GET',
          credentials: 'include',
        })
        if (!res.ok) {
          setError('Failed to fetch comments')
        } else {
          const raw = await res.json()
          const data = raw.map(comment => ({
            content: comment.content,
            fullName: comment.user?.fullname,
            nickName: comment.user?.nickname,
            imagePath: comment.img || comment.image_path,
            userImage: comment.user.avatar,
            createdAt: comment.created_at || new Date().toISOString(),
            likes: comment.likes || 0,
          }))
          setComments(data)
        }
        setLoading(false)
      } catch (err) {
        setLoading(false)
        console.error("Error fetching comments:", err)
        setError('Failed to load comments. Please try again later.')
      }
    }
    fetchComments()
  }, [id, groupID])

  if (loading) return <Loader />
  if (error) return <Error error={error} />

  return (
    <section className="comments_container w-full h-full flex-col justify-between gap-2">
      {comments.length === 0 ? <img src='/no-comments.svg' className='no_comments' alt="No comments available" /> : <Comments comments={comments} id={id} groupID={groupID} creatorID={creatorID} />}
      <CommentsFooter id={id} groupID={groupID} setComments={setComments} onCommentMessage={onCommentMessage} />
    </section>
  )
}