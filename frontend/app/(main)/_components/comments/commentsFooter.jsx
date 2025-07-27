import React, { useEffect } from 'react'
import "./comments.css"
import { MdPermMedia } from "react-icons/md";
import { FaPaperPlane } from "react-icons/fa";

import { useActionState } from 'react'
import { commentPostAction } from '@/app/_actions/posts'
import Button from '@/app/_components/button';

export default function CommentsFooter({ id, setComments, onCommentMessage }) {
    const initialState = {
        message: '',
        success: false,
        content: '',
        firstName: '',
        imagePath: '',
        createdAt: '',
        userImage: '',
        likes: 0,
    };

    const [state, formAction] = useActionState(commentPostAction, initialState)

    useEffect(() => {
        if (state.success) {
            const newComment = {
                content: state.content,
                firstName: state.firstName,
                lastName: "",
                imagePath: state.userImage,
                createdAt: state.createdAt || new Date().toISOString(),
                likes: state.likes || 0,
            };

            setComments(prev => [newComment, ...prev]);

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
                required
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
