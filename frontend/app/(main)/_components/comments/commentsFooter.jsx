import "./comments.css"
import { useEffect } from 'react'
import { MdPermMedia } from "react-icons/md";
import { FaPaperPlane } from "react-icons/fa";
import { useActionState } from 'react'
import { commentPostAction } from '@/app/_actions/posts'
import { commentGroupPostAction } from '@/app/_actions/groupPosts'
import { useNotification } from "../../_context/NotificationContext";

export default function CommentsFooter({ id, groupID = null, setComments, onCommentMessage }) {
    const initialState = {
        group: groupID ? true : false,
        groupID: groupID,
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
    const {showNotification } = useNotification()

    useEffect(() => {
        let state = groupID ? groupActionState : postActionState
        console.log("comment state: ", state)
        if (state?.message) {
            const newComment = {
                content: state.content,
                nickName: state.nickname,
                fullName: state.fullName,
                userImage: state.avatar,
                createdAt: new Date(),
                imagePath: state.imagePath,
            };
            console.log("new comment: ", state)
            setComments(prev => [newComment, ...prev]);
            showNotification({Content:state.message, Status:"success"})

            if (onCommentMessage) {
                onCommentMessage("A new comment was added");
            }
        } else if (state?.errors || state?.error) {
            console.log("gggggggggggggg", state.error)
            showNotification({Content:state.error, Status:"error"})
        }
    }, [postActionState, groupActionState])
    return (
        <form
            action={groupID ? groupAction : postAction}
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
            {groupID && <input type="hidden" name="groupId" value={groupID} />}

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