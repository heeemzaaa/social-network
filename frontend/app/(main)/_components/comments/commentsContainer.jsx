import React, { useEffect, useState } from 'react'
import "./comments.css"
import Comments from './comments'
import CommentsFooter from './commentsFooter'

export  default function CommentsContainer({ id }) {
  const [comments, setComments] = useState([]);


  useEffect(() => {
    const fetchComments = async () => {
      try {
        const res = await fetch(`http://localhost:8080/api/posts/comments/${id}`, {
         credentials : 'include'
        });
        const data = await res.json();
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
      <CommentsFooter id={id} setComments={setComments} />
    </section>
  );
}
