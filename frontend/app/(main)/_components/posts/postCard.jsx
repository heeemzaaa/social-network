"use client "
import { FaRegHeart, FaHeart, FaRegComment } from "react-icons/fa";
import "./style.css"
import Avatar from "../avatar";
import { useModal } from "../../_context/ModalContext";
import { useState } from "react";

export default function PostCard({
    id,
    user,
    content,
    created_at,
    img,
    total_likes,
    total_comments,
    liked,
    privacy
}) {

    const { openModal } = useModal()
    const [isLiked, setIsLiked] = useState(liked === 1);
    const [likes, setLikes] = useState(total_likes);

    const handleToggleLike = (id) => {
        setIsLiked(prev => !prev);
        setLikes(prev => isLiked ? prev - 1 : prev + 1);

        // TODO: send to backend
    };

    return (

        <div className="post-card">
            <div className="post-card-body">
                <div className="post-card-header">
                    <div className="flex align-center gap-1">
                        <Avatar size="42" />
                        <h3 className="post-user">
                            {user.firstname} {user.lastname}
                        </h3>
                    </div>
                    <span className="post-privacy">{privacy}</span>
                </div>
                <p className="post-content">{content}</p>
                {img && (
                    <div className="post-card-img">
                        <img src={`http://localhost:8080/static/${img}`} alt={img} />
                    </div>
                )}
                <span>{new Date(created_at).toISOString().slice(0, 16).replace('T', ' ')}</span>
                <div className="post-actions flex gap-2 align-center" >
                    <div style={actionStyle} onClick={handleToggleLike} >
                        {isLiked ? <FaHeart color="red" /> : <FaRegHeart />}
                        <span>
                            {likes}
                        </span>
                    </div>
                    <div className="glass-bg" onClick={() => { openModal("pass comments component here") }}>
                        <div style={actionStyle}>
                            <FaRegComment />
                            <span>
                                {total_comments}
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