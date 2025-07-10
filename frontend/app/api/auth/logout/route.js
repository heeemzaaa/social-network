
export async function POST(request) {
    console.log("===> Request to /api/auth/logout: ", request.headers)
    try {
        console.log("+ requeset headers :", request.headers)
        const res = await fetch("http://localhost:8080/api/auth/logout", {
            method: "POST",
            credentials: "include",
            headers: request.headers,
        });
        return res
    } catch (error) {
        console.error("Error in API route:", error.message);
    }
}