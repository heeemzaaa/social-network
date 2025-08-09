import React, { useEffect, useState } from 'react'
import "./comments.css"
import Comments from './comments'
import CommentsFooter from './commentsFooter'

export default function CommentsContainer({ id, onCommentMessage, groupID, creatorID }){
  const [comments, setComments] = useState([]);

  const postComment = `http://localhost:8080/api/posts/comments/${id}`
  const groupComment = `http://localhost:8080/api/groups/${groupID}/posts/${id}/comments?offset=0`
  useEffect(() => {
    const fetchComments = async () => {
      try {
        const res = await fetch(groupID ? groupComment : postComment, {
          method : 'GET',
          credentials: 'include',
        });
        const raw = await res.json();

        const data = raw.map(comment => ({
          content: comment.content,
          fullName: comment.user?.fullname,
          nickName: comment.user?.nickname,
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
    <section className="comments_container w-full h-full flex-col justify-between gap-2">
      {comments.length  === 0 ? <img src='/no-comments.svg' className='no_comments'/> :   <Comments comments={comments} id={id} groupID={groupID} creatorID={creatorID} />}
      <CommentsFooter id={id} groupID={groupID} setComments={setComments}  onCommentMessage={onCommentMessage}/>
    </section>
  );
}
