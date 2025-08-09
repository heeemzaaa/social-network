"use server";


import { cookies } from "next/headers";

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

    const response = await fetch("http://localhost:8080/api/posts", {
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
    const postId = formData.get("postId");
    if (!postId) {
        return { ...prevState, message: "Post ID is required." };
    }
    try {
        const cookieStore = await cookies();
        const sessionCookie = cookieStore.get("session")?.value;
        const res = await fetch(`http://localhost:8080/api/posts/like/${postId}`, {
            method: "POST",
            headers: sessionCookie ? { Cookie: `session=${sessionCookie}` } : {},
        });
        const data = await res.json();
        if (data.success) {
            return {
                message: "Liked successfully!",
                liked: data.liked,
                likes: data.total_likes,
            };
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
        const resp = await fetch("http://localhost:8080/api/posts/comment", {
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
