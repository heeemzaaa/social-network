import React, { useEffect } from 'react'
import "./comments.css"
import { MdPermMedia } from "react-icons/md";
import { FaPaperPlane } from "react-icons/fa";
import { useActionState } from 'react'
import { commentPostAction } from '@/app/_actions/posts'
import { commentGroupPostAction } from '@/app/_actions/groupPosts'

export default function CommentsFooter({ id, groupId = null,  setComments, onCommentMessage }) {
    console.log("+=====> group id: ", groupId)
    const initialState = {
        group: groupId ? true : false,
        groupId: groupId,
        message: '',
        content: '',
        nickname: '',
        fullName: '',
        avatar: '',
        success: false,
        createdAt: '',
        imagePath: '',
    };

    const [postActionState, postAction] = useActionState(commentPostAction, initialState)
    const [groupActionState, groupAction] = useActionState(commentGroupPostAction, initialState)
    useEffect(() => {
        let state = postActionState.success ?  postActionState :  groupActionState.success ? groupActionState : null
        if (state?.success) {
            const newComment = {
                content: state.content,
                nickName: state.nickname,
                fullName: state.fullName,
                userImage: state.avatar,
                createdAt: new Date(),
                imagePath: state.imagePath,
            };
            setComments(prev => [newComment, ...prev]);

            if (onCommentMessage) {
                onCommentMessage("A new comment was added");
            }
        }
    }, [postActionState, groupActionState]); // This will run every time state changes (after submission)
    return (
        <form
            action={groupId ? groupAction : postAction}
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
            {groupId && <input type="hidden" name="groupId" value={groupId} />}

            <input
                type="text"
                name="content"
                className="w-full p1 rounded-md"
                placeholder="Write a comment..."
            />

            <button type="submit" className='submit_comment'>
                <FaPaperPlane size="24px" />
            </button>

        </form>
    );
}
