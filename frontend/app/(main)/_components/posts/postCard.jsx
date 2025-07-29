"use client "
import { FaRegHeart, FaHeart, FaRegComment } from "react-icons/fa";
import "./style.css"
import Avatar from "../avatar";
import { useModal } from "../../_context/ModalContext";
import { likePostAction } from "@/app/_actions/posts";
import { useActionState, useState } from "react";
import CommentsContainer from "../comments/commentsContainer";
export default function PostCard({
    id,
    user,
    content,
    created_at,
    image_path,
    total_likes,
    total_comments,
    liked,
    privacy
}) {
    const [totalComments, setTotalComments] = useState(total_comments);

    const handleCommentMessage = (msg) => {
        console.log(msg); 
        setTotalComments(prev => prev + 1);
    };
    const { openModal } = useModal()
    const initialState = {
        liked: liked === 1,
        likes: total_likes,
        message: null,
    };

    const [state, formAction] = useActionState(likePostAction, initialState);
    console.log('user.avatar', user.avatar)
    return (
        <div className="post-card">
            <div className="post-card-body">
                <div className="post-card-header">
                    <div className="flex align-center gap-1">
                        <Avatar img={user.avatar} size="42" />
                        <h3 className="post-user">
                            {user.fullname}
                        </h3>
                    </div>
                    <span className="post-privacy">{privacy}</span>
                </div>
                <p className="post-content">{content}</p>
                {image_path && (
                    <div className="post-card-img">
                        <img src={`http://localhost:8080/static/${image_path}`} alt={image_path} />
                    </div>
                )}
                <span>{new Date(created_at).toISOString().slice(0, 16).replace('T', ' ')}</span>
                <div className="post-actions flex gap-2 align-center" >
                    <form action={formAction}>
                        <input type="hidden" name="postId" value={id} />
                        <div className="post-actions flex gap-2 align-center">
                            <button type="submit" style={actionStyle}>
                                {state.liked ? <FaHeart color="red" /> : <FaRegHeart />}
                                <span>{state.likes || 0}</span>
                            </button>
                        </div>
                    </form>
                    <div className="glass-bg" onClick={() => { openModal(<CommentsContainer id={id} onCommentMessage={handleCommentMessage} />) }}>
                        <div style={actionStyle}>
                            <FaRegComment />
                            <span>
                                {totalComments}
                            </span>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
}

const actionStyle = {
    border: "1px solid",
    color: "black",
    background: "#eee",
    height: "min-content",
    display: "flex",
    alignItems: "center",
    gap: "4px",
    padding: ".3rem .8rem",
    borderRadius: "var(--border-radius-sm)",
    cursor: "pointer"
}