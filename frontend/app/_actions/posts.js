"use server";

import { cookies } from "next/headers";

export async function createPostAction(prevState, formData) {
    let state = {
        error: null,
        errors: {},
        message: null,
    };

    const title = formData.get("title");
    const content = formData.get("content");
    const privacy = formData.get("privacy");
    const selectedFollowersRaw = formData.get("selectedFollowers");

    if (!title) {
        state.errors.title = "Title is required";
    }
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

    const sessionCookie = cookies().get("session")?.value;

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
    console.log(createdPost, "post created successfully");

    return {
        message: "Post created successfully",
        post: createdPost,
    };
}
