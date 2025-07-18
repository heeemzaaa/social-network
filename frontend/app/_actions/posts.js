"use server";

import { cookies } from "next/headers";
export async function createPostAction(prevState, formData) {
    let state = {
        error: null,
        errors: {},
        message: null,
    };
    const content = formData.get("content");
    const privacy = formData.get("privacy");
    const selectedFollowersRaw = formData.get("selectedFollowers");
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
    if (Object.keys(state.errors).length > 0) {
        return state;
    }
    const cookieStore = await cookies();
    const sessionCookie = cookieStore.get("session")?.value;
    const response = await fetch("http://localhost:8080/api/posts", {
        method: "POST",
        body: formData,
        headers: sessionCookie ? { Cookie: `session=${sessionCookie}` } : {},
    });
    if (!response.ok) {
        const errorJson = await response.json();
        console.log("Server error:", errorJson);
        state.error = errorJson?.message || "Failed to create post";
        return state;
    }
    const createdPost = await response.json();
    return {
        message: "Post created successfully",
        data: createdPost
    };
}



export async function likePostAction(prevState,formData) {
        const postId = formData.get("postId");
        console.log("post id ", postId)
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
            return { message: "Failed to like post." };
        }
    } catch (err) {
        console.error("Error liking post:", err);
        return { message: "Server error." };
    }
}
