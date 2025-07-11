// app/_lib/posts.js
export async function fetchPosts() {
    console.log("Calling fetchPosts ");
    const res = await fetch(`http://localhost:3000/api/posts`,{
        method :"GET",
        credentials : "include"
    });
    console.log("inside fetchPosts: ", res)
    const data = await res.json();
    return data;
}