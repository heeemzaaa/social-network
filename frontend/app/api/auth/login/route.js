export async function POST(request) {
    try {
        const requestBody = await request.json(); // Parse incoming request body
        const res = await fetch("http://localhost:8080/api/auth/login", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(requestBody) // Forward the request body
        });
        return res
    } catch (error) {
        console.error("Error in login API route:", error.message);
    }
}