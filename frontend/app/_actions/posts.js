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
    const img = formData.get("img");

    const maxSize = 5 * 1024 * 1024; // 5MB

    if (!content) {
        state.errors.content = "Content is required";
    }

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
            state.errors.img = "Image must be JPEG, PNG, or GIF";
        } else if (img.size > maxSize) {
            state.errors.img = "Image file size must be less than 5MB";
        }
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
            error: errorJson?.message || "Failed to create post",
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
        const cookieStore = cookies();
        const sessionCookie = cookieStore.get("session")?.value;

        const res = await fetch(`http://localhost:8080/api/posts/like/${postId}`, {
            method: "POST",
            headers: sessionCookie ? { Cookie: `session=${sessionCookie}` } : {},
        });

        const data = await res.json();

        if (data.success) {
            return { message: "Liked successfully!" };
        } else {
            return { message: data.message || "Failed to like post." };
        }
    } catch (err) {
        console.error("Error liking post:", err);
        return { message: "Server error." };
    }
}
