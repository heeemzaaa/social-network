"use server";

import { cookies } from "next/headers";

const API_URL = process.env.BACKEND_URL || 'http://localhost:8080'


export async function createPostAction(prevState, formData) {
    let state = {
        error: null,
        errors: {},
        message: null,
    };

    const content = formData.get("content")?.trim();
    const privacy = formData.get("privacy");
    const selectedFollowersRaw = formData.get("selectedFollowers");
    const img = formData.get("img")

    const maxSize = 3 * 1024 * 1024; // 3MB

    if (!privacy) {
        state.errors.privacy = "Privacy is required";
    }

    if (privacy === "private") {
        if (!selectedFollowersRaw) {
            state.errors.selectedFollowers = "Please choose friends";
        } else {
            try {
                const selectedFollowers = JSON.parse(selectedFollowersRaw);
                if (!Array.isArray(selectedFollowers) || selectedFollowers.length === 0) {
                    state.errors.selectedFollowers = "Please choose at least one friend";
                }
            } catch {
                state.errors.selectedFollowers = "Invalid selected followers data";
            }
        }
    }

    if (img && img.size > 0) {
        const allowedTypes = ["image/jpeg", "image/png", "image/gif"];
        if (!allowedTypes.includes(img.type)) {
            state.errors.img = "jpge png gif only are allowed";
        } else if (img.size > maxSize) {
            state.errors.img = "Image file size must be less than 5MB";
        }
    }

    if (img.size == 0 && !content) {
        state.errors.img = "one filed is requied to create post";
    }

    if (Object.keys(state.errors).length > 0) {
        return {
            ...prevState,
            ...state,
            error: "Please fix the highlighted fields.",
        };
    }

    const postData = {
        content,
        privacy,
        selectedFollowers:
            privacy === "private" && selectedFollowersRaw
                ? JSON.parse(selectedFollowersRaw)
                : [],
    };
    const newFormData = new FormData();
    newFormData.append("data", JSON.stringify(postData));
    if (img && img.size > 0) {
        newFormData.append("img", img);
    } else {
        if (!content) {
            state.errors.img = "one filed is required";
        }
    }
    const cookieStore = await cookies();
    const sessionCookie = cookieStore.get("session")?.value;

    const response = await fetch(`${API_URL}/api/posts/`, {
        method: "POST",
        body: newFormData,
        headers: sessionCookie ? { Cookie: `session=${sessionCookie}` } : {},
    });

    if (!response.ok) {
        const errorJson = await response.json().catch(() => null);
        console.error("Backend error:", errorJson);
        return {
            ...state,
            errors: errorJson.errors,
            error: errorJson?.message || "Failed to create post?????",
        };
    }

    const createdPost = await response.json();
    return {
        message: "Post created successfully",
        data: createdPost,
    };
}


export async function likePostAction(prevState, formData) {
    console.log(prevState)

    let url
    let body
    const postId = formData.get("postId");
    const groupId = formData.get("groupId");

    if (!postId) {
        return { ...prevState, message: "Post ID is required." };
    }

    if (groupId && postId) {
        url = `${API_URL}/api/groups/${groupId}/react/like`
        body = {
            entity_type: "post",
            entity_id: postId
        }
    } else {
        url = `${API_URL}/api/posts/like/${postId}`
    }

    console.log()
    try {
        const cookieStore = await cookies();
        const sessionCookie = cookieStore.get("session")?.value;
        const res = await fetch(url, {
            method: "POST",
            body: JSON.stringify(body || {}),
            headers: sessionCookie ? { Cookie: `session=${sessionCookie}` } : {},
        });
        const data = await res.json();
        if (res.ok) {
            let state = {
                message: "Liked successfully!",
                liked: data.liked || data.reaction,
                likes: data.total_likes || 0,
            }
            if (data.reaction) {
                console.log(data.reaction)
                data.reaction === 1 ? state.likes++ : state.likes--
            }
            return state
        } else {
            return { ...prevState, message: data.message || "Failed to like post." };
        }
    } catch (err) {
        console.error("Error liking post:", err);
        return { ...prevState, message: "Server error." };
    }
}


export async function commentPostAction(prevState, formData) {
    console.log("======> inside the comm")

    let state = {
        error: null,
        errors: {},
        message: null,
    };

    const commentContent = formData.get("content")?.trim();
    const postID = formData.get("postID");
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
        post_id: postID,
        content: commentContent,
    });

    const newFormData = new FormData();
    newFormData.append("data", jsonData);

    if (commentImg && commentImg.size > 0) {
        newFormData.append("img", commentImg);
    }

    const cookieStore = cookies();
    const sessionCookie = cookieStore.get("session")?.value;

    try {
        const resp = await fetch(`${API_URL}/api/posts/comment`, {
            method: "POST",
            credentials: "include",
            headers: sessionCookie ? { Cookie: `session=${sessionCookie}` } : {},
            body: newFormData,
        });

        if (!resp.ok) {
            return { ...state, message: "Failed to post comment." };
        }

        const response = await resp.json();
        console.log("post comment: ", response)
        const now = new Date();
        const formatted = now.toISOString().slice(0, 16).replace('T', ' ');
        return {
            ...state,
            message: "Commented successfully",
            content: response.content,
            nickname: response.user.nickname,
            fullName: response.user.fullname,
            avatar: response.user.avatar,
            created_at: formatted,
            imagePath: response.img || response.image_path,
            success: true,
        };
    } catch (err) {
        return { ...prevState, message: "Server error." };
    }
}
