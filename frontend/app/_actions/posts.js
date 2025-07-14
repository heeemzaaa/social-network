"use server"

import { cookies } from "next/headers"


export async function createPostAction(prevState, formData) {
    let state = {
        error: null,
        errors: {},
        message: null,
    };

    const title = formData.get("title");
    const content = formData.get("content");

    if (!title) {
        state.errors.title = "Title is required";
    }
    if (!content) {
        state.errors.content = "Content is required";
    }

    if (Object.keys(state.errors).length > 0) {
        return state;
    }
    const sessionCookie = cookies().get("session")?.value;
    const response = await fetch("http://localhost:8080/api/posts", {
        method: "POST",
        body: formData,
        headers: sessionCookie ? { Cookie: `session=${sessionCookie}` } : {}
    });

    console.log(await response.json())
    if (!response.ok) {
        throw new Error("Failed to create post");
    }

    const createdPost = await response.json();
    console.log(createdPost, "--");
    return createdPost;
}
