export async function postAction(prevState, formData) {
  const state = {
    errors: {},
    error: null,
    message: null,
  };

  const title = formData.get("title")?.trim();
  const content = formData.get("content")?.trim();

  if (!title) state.errors.title = "Title can't be empty";
  if (!content) state.errors.content = "Content can't be empty";

  if (Object.keys(state.errors).length > 0) {
    return state; 
  }

  try {
    const res = await fetch("/api/posts", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ title, content }),
    });

    const data = await res.json();

    if (!res.ok) {
      state.error = data.error || "Failed to create post";
      state.errors = data.errors || null;
      return state;
    }

    state.message = "Post created successfully!";
    state.post = data; 

  } catch (error) {
    console.error(error);
    state.error = "An unexpected error occurred";
  }

  return state;
}
