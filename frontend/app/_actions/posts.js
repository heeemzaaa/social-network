"use server"

export async function createPostAction(prevState, formData) {

    let state = {
        error: null,
        errors: {},
        message: null,
    }

    const data = Object.fromEntries(formData.entries());
    let title = formData.get("title")
    let content = formData.get("content")

    if (title == "") {
        state.errors.title = "title is requied"
    }
    if (content == "") {
        state.errors.title = "contenet is requied"
    }

    if (Object.keys(state.errors).length > 0) {
        return state;
    }

    const response = await fetch('/api/posts', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(data),
    });

    if (!response.ok) {
        throw new Error('Failed to create post');
    }

    const createdPost = await response.json();
    console.log(createdPost ,"--")
    return createdPost;
}
