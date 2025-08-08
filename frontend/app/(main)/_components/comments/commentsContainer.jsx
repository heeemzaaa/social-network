import "./comments.css"
import Comments from './comments'
import CommentsFooter from './commentsFooter'
import {
  useEffect,
  useState
} from 'react'

export default function CommentsContainer({ id, onCommentMessage }) {
  const [comments, setComments] = useState([]);


  useEffect(() => {
    const fetchComments = async () => {
      try {
        const res = await fetch(`http://localhost:8080/api/posts/comments/${id}`, {
          method: 'GET',
          credentials: 'include',
        });
        const raw = await res.json();
        const data = raw.map(comment => ({
          content: comment.content,
          fullName: comment.user?.fullname,
          nickName: comment.user.nickname,
          imagePath: comment.img,
          userImage: comment.user.avatar,
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
      {comments.length === 0 ? <img src='/no-comments.svg' className='no_comments' /> : <Comments comments={comments} />}
      <CommentsFooter id={id} setComments={setComments} onCommentMessage={onCommentMessage} />
    </section>
  );
}
