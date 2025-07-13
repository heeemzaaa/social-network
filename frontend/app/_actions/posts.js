"use server"

export async function createPostAction(prevState, formData) {

    let state = {
        error: null,
        errors: {},
        message: null,
    }
    console.log("post action run by the next server")
    const data = Object.fromEntries(formData.entries());

    let title = formData.get("title")
    //   let title = formData.get("title")

    if (title == "") {
        state.errors.title = "title is requied"
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

    return createdPost;
}
