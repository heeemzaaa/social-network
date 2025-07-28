"use client "
import { FaRegHeart, FaHeart, FaRegComment } from "react-icons/fa"
import "./style.css"
import Avatar from "../avatar"
import { useModal } from "../../_context/ModalContext"
import { likePostAction } from "@/app/_actions/posts"
import { useActionState, useState } from "react"
import CommentsContainer from "../comments/commentsContainer"
import { useRouter } from "next/navigation"
import { timeAgo } from "@/app/_utils/time"
import { HiOutlineClock } from "react-icons/hi2"

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
    const [totalComments, setTotalComments] = useState(total_comments)
    console.log(user)
    const handleCommentMessage = (msg) => {
        setTotalComments(prev => prev + 1)
    }
    const { openModal } = useModal()
    const initialState = {
        liked: liked === 1,
        likes: total_likes,
        message: null,
    }

    const router = useRouter()
    const navigateToProfile = (profileId) => {
        router.push(`/profile/${profileId}`);
    }


    const [state, formAction] = useActionState(likePostAction, initialState)
    return (
        <div className="post-card">
            <div className="post-card-body">
                <div className="post-card-header">
                    <div className="flex align-center gap-1">
                        <Avatar img={user.avatar} size="42" />
                        <div onClick={() => navigateToProfile(user.id)}>
                            <h3 className="post-user hover-pointer">
                                {user.fullname}
                            </h3>
                            <span className="hover-pointer text-sm font-medium " style={{opacity:".8"}}>
                                {user.nickname && `@${user.nickname}`}
                            </span>
                        </div>
                    </div>
                    <span className="post-privacy">{privacy}</span>
                </div>
                <p className="post-content">{content}</p>
                {image_path && (
                    <div className="post-card-img" >
                        <img src={`http://localhost:8080/static/${image_path}`} alt={image_path} className="rounded-md" style={{width: '100%', height: 'max-content'}} />
                    </div>
                )}

                <div className="post-actions flex gap-2 align-center flex-wrap" >
                    <form action={formAction}>
                        <input type="hidden" name="postId" value={id} />
                        <div className="post-actions flex gap-2 align-center">
                            <button type="submit" style={actionStyle}>
                                {state.liked ? <FaHeart color="red" /> : <FaRegHeart />}
                                <span>{state.likes || 0}</span>
                            </button>
                        </div>
                    </form>
                    <div onClick={() => { openModal(<CommentsContainer id={id} onCommentMessage={handleCommentMessage} />) }}>
                        <div style={actionStyle}>
                            <FaRegComment />
                            <span>
                                {totalComments}
                            </span>
                        </div>
                    </div>
                    <div style={{opacity:".5", gap:"5px", paddingLeft:"3px", marginLeft:"auto"}} className="flex align-end">
                        <HiOutlineClock size={24} />
                        <span>{timeAgo(created_at)}</span>
                    </div>
                </div>
            </div>
        </div>
    )
}

const actionStyle = {
    fontSize: '20px',
    height: "min-content",
    display: "flex",
    alignItems: "center",
    gap: "4px",
    padding: ".3rem .8rem",
    borderRadius: "var(--border-radius-sm)",
    cursor: "pointer"
}