import React, { useEffect } from 'react'
import "./comments.css"
import { MdPermMedia } from "react-icons/md";
import { FaPaperPlane } from "react-icons/fa";

import { useActionState } from 'react'
import { commentPostAction } from '@/app/_actions/posts'

export default function CommentsFooter({ id, setComments, onCommentMessage }) {
    const initialState = {
        message: '',
        content: '',
        nickname: '',
        fullName: '',
        avatar: '',
        success: false,
        createdAt: '',
        commentImage: '',
    };

    const [state, formAction] = useActionState(commentPostAction, initialState)

    useEffect(() => {
        if (state.success) {
            const newComment = {
                content: state.content,
                nickName: state.nickname,
                fullName: state.fullName,
                userImage: state.avatar,
                createdAt: new Date(),
                ImagePath: state.img,
            };
            console.log('newComment', newComment)
            setComments(prev => [...prev , newComment]);

            if (onCommentMessage) {
                onCommentMessage("A new comment was added");
            }
        }
    }, [state]); // This will run every time state changes (after submission)

    return (
        <form
            action={formAction}
            className='comments_footer flex justify-center align-center p1 gap-1'
        >
            <label htmlFor="commentImg">
                <MdPermMedia size="24px" style={{ cursor: 'pointer' }} />
            </label>

            <input
                type="file"
                id='commentImg'
                name='commentImg'
                style={{ display: 'none' }}
            />

            <input type="hidden" name="postID" value={id} />

            <input
                type="text"
                name="content"
                className="w-full p1 rounded-md"
                placeholder="Write a comment..."
            />

            <button type="submit" className='submit_comment'>
                <FaPaperPlane size="24px" />
            </button>

            {state.errors?.commentContent && (
                <span className="field-error">{state.errors.commentContent}</span>
            )}
        </form>
    );
}
