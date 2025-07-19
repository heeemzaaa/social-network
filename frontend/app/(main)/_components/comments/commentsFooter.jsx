import React from 'react'
import "./comments.css"
import { MdPermMedia } from "react-icons/md";
import { FaPaperPlane } from "react-icons/fa";

import { useActionState } from 'react'
import { commentPostAction } from '@/app/_actions/posts'
import post from '../posts';


export default function CommentsFooter({ id }) {
    const initialState = {
        message: '',
        success: false

    }
    const [state, formAction] = useActionState(commentPostAction, initialState)

    return (
        <form
            action={formAction}
            className='comments_footer flex justify-between align-center p1 gap-2'
        >   <label htmlFor="commentImg">
                <MdPermMedia size="24px" style={{ cursor: 'pointer' }} />
            </label>
            <input
                type="file"
                id='commentImg'
                name='commentImg'
                style={{display:'none'}}
            />
            <input type="hidden" name="postID" value={id} />
            <input
                type="text"
                name="content"
                className="w-full p1 rounded-md"
                placeholder="Write a comment..."
                required
            />
            <button type="submit">
                <FaPaperPlane size="24px" />
            </button>
            {state.errors?.commentContent && <span className="field-error">{state.errors.commentContent}</span>}
        </form>
    )
}
