"use server";

import { cookies } from "next/headers";


export async function commentGroupPostAction(prevState, formData) {    
    let state = {
        error: null,
        errors: {},
        message: null,
    };

    const commentContent = formData.get("content")?.trim();
    const postID = formData.get("postID");
    const groupID = formData.get("groupId")
    const commentImg = formData.get("commentImg");
    const maxSize = 3 * 1024 * 1024;

    if (!commentContent && commentImg.size === 0) {
        state.errors.commentContent = "Input comment is required";
        return state;
    }
    if (!postID) {
        state.errors.postID = "Post ID is required";
        return state;
    }
    if (commentImg && commentImg.size > 0) {
        const allowedTypes = ["image/jpeg", "image/png", "image/gif"];
        if (!allowedTypes.includes(commentImg.type)) {
            state.errors.commentImg = "jpeg png gif only are allowed";
            return state;
        } else if (commentImg.size > maxSize) {
            state.errors.commentImg = "Image file size must be less than 3MB";
            return state;
        }
    }

    const jsonData = JSON.stringify({
        id: postID,
        group_id: groupID,
        content: commentContent,
    });

    console.log("---------------> ", jsonData);
    

    const newFormData = new FormData();
    newFormData.append("data", jsonData);

    if (commentImg && commentImg.size > 0) {
        newFormData.append("img", commentImg);
    }

    const cookieStore = await cookies();
    const sessionCookie = cookieStore.get("session")?.value;

    try {
        const resp = await fetch(`http://localhost:8080/api/groups/${groupID}/posts/${postID}/comments/`, {
            method: "POST",
            credentials: "include",
            headers: sessionCookie ? { Cookie: `session=${sessionCookie}` } : {},
            body: newFormData,
        });

        if (!resp.ok) {
            console.log("error fetching request");
            return { ...state, message: "Failed to post comment." };
        }
        
        const response = await resp.json();
        console.log('response', response)        
        const now = new Date();
        const formatted = now.toISOString().slice(0, 16).replace('T', ' ');
        return {
            ...state,
            message: "Commented successfully",
            content: response.content,
            nickname: response.user.nickname,
            fullName: response.user.fullname,
            avatar: response.user.avatar,
            success: true,
            createdAt: formatted,
            imagePath: response.img,
        };
    } catch (err) {
        return { ...prevState, message: "Server error." };
    }
}
