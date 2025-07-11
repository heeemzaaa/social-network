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
        if (!res.ok) {
            throw new Error(`API request failed with status ${res.status}`);
        }
        const data = await res.json();
        return Response.json(data);
    } catch (error) {
        console.error("Error in API route:", error.message);
        return Response.json(
            { error: error.message },
            { status: error.status || 500 }
        );
    }
}