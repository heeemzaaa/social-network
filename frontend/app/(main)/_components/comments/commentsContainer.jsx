import React, { useEffect, useState } from 'react'
import "./comments.css"
import Comments from './comments'
import CommentsFooter from './commentsFooter'

export default function CommentsContainer({ id, onCommentMessage }){
  const [comments, setComments] = useState([]);


  useEffect(() => {
    const fetchComments = async () => {
      try {
        const res = await fetch(`http://localhost:8080/api/posts/comments/${id}`, {
          method : 'GET',
          credentials: 'include',
        });
        const raw = await res.json();

        const data = raw.map(comment => ({
          content: comment.content,
          firstName: comment.user.nickname = "" || `${comment.user?.firstname || ""} ${comment.user?.lastname || ""}`,
          imagePath: comment.img,
          userImage : comment.user.avatar,
          createdAt: comment.created_at || new Date().toISOString(),
          likes: comment.likes || 0,
        }));

        setComments(data);
      } catch (err) {
        console.error("Error fetching comments:", err);
      }
    };
    fetchComments();
  }, [id]);


  return (
    <section className="comments_container w-full flex-col gap-2">
      <Comments comments={comments} />
      <CommentsFooter id={id} setComments={setComments}  onCommentMessage={onCommentMessage}/>
    </section>
  );
}
