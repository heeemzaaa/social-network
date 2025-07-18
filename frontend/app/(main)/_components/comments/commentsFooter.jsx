import React from 'react'
import "./comments.css"
import { MdPermMedia } from "react-icons/md";
import { FaPaperPlane } from "react-icons/fa";



export default function CommentsFooter() {
    return (
        <form className='comments_footer flex justify-between align-center p1 gap-2'>
            <MdPermMedia size={'24px'} />
            <input type="text" className='w-full p1 rounded-md'/>
            <FaPaperPlane size={'24px'}/>
        </form>
    )
}
