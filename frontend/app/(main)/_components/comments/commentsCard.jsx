'use client'
import "./comments.css"
import Avatar from "../avatar"
import { timeAgo } from '@/app/_utils/time'

export default function CommentsCard({ comment, id, groupID, creatorID }) {
    return (
        <div className='comments_card w-full p3 flex-col gap-3 shadow-lg'>
            <div className='card_header flex align-center gap-1'>
                <Avatar img={comment.userImage} size="42" />
                <div className='flex-col'>
                    <span className='user_name text-md'>{comment.fullName}</span>
                    <span className='nickname_comment text-sm' style={{ opacity: '.5' }}>{comment.nickName && `@${comment.nickName}`}</span>
                </div>
            </div>

            <div className='card_body flex-col justify-start gap-1'>
                <h4 className='comment'>{comment.content}</h4>
                <div className='w-full flex justify-center'>
                    {comment.imagePath && <img src={`http://localhost:8080/static/${comment.imagePath}?creator=${creatorID || ""}&postID=${id || ""}&groupID=${groupID || ""}`} className='comments_image' style={{ width: '70%', height: 'max-content' }} />}
                </div>
            </div>

            <div className='card_footer flex justify-end'>
                <p className='time text-md'>{timeAgo(comment.createdAt)}</p>
            </div>
        </div>
    )
}
