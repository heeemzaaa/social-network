// app/api/posts/route.js

export async function GET(request) {
    console.log("Request to /api/auth/logout: ", request)
    try {
        const res = await fetch("http://localhost:8080/api/auth/loggedin");
        return res
    } catch (error) {
        console.error("Error in API route:", error.message);
        return Response.json(
            { error: error.message },
            { status: error.status || 500 }
        );
    }
}