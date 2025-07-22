"use client "

import { FaRegHeart, FaHeart, FaRegComment } from "react-icons/fa";
import "./style.css"
import Avatar from "../avatar";
import { useModal } from "../../_context/ModalContext";
import { useState } from "react";
import  CommentsContainer  from "@/app/(main)/_components/comments/commentsContainer"

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
    console.log('user', user)
    const { openModal } = useModal()
    const [isLiked, setIsLiked] = useState(liked === 1);
    const [likes, setLikes] = useState(total_likes || 0);

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
                        <Avatar img={user.avatar ? `http://localhost:8080/static/${user.avatar}` : '/no-profile.png'} size="42"  />
                        <div className="flex-col text-sm">
                            <span className="post-user">
                                {user.fullname ? `${user.fullname}` : `${user.firstname} ${user.lastname}`}
                            </span>
                            <span>{`@${user.nickname}`}</span>
                        </div>
                    </div>
                    <span className="post-privacy">{privacy}</span>
                </div>
                <p className="post-content">{content}</p>
                {image_path && (
                    <div className="post-card-img">
                        <img src={`http://localhost:8080/static/${image_path}`} alt="Post" />
                    </div>
                )}
                <div className="post-actions flex gap-2 align-center" >
                    <div style={actionStyle} onClick={handleToggleLike} >
                        {isLiked ? <FaHeart color="red" /> : <FaRegHeart />}
                        <span>
                            {likes}
                        </span>
                    </div>
                    <div className="glass-bg" onClick={() => { openModal(<CommentsContainer id={id} />) }}>
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