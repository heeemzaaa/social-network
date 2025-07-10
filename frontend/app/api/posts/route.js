// app/api/posts/route.js

export async function GET(request) {
    console.log("=====> requesting api/auth/logout ", request.headers)
    try {
        const res = await fetch("http://localhost:8080/api/posts", {
            method : "GET",
            credentials : "include",
            headers: request.headers,
        });
        return res
    } catch (error) {
        console.error("Error in API route:", error.message);
    }
}