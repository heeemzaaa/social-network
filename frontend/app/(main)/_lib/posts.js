// app/_lib/posts.js

export async function fetchPosts() {
    console.log("Calling fetchPosts ");
    const res = await fetch(`http://localhost:3000/api/posts`, { cache: "no-store" });
    const data = await res.json();
    if (!res.ok) {
        throw new Error(data.error || "Failed to fetch posts");
    }
    return data;
}