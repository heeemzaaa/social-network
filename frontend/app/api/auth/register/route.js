// frontend/app/api/auth/register/route.js
export async function POST(request) {
    console.info("inside the register api route ®️");

    try {
        // Get the form data from the incoming request
        const formData = await request.formData();
        console.log("===> Form Data: ", formData)
        // Forward the request to the Go API
        const res = await fetch("http://localhost:8080/api/auth/register", {
            method: "POST",
            body: formData,
            credentials: 'include', // Include cookies in the request to the Go API
            headers: {
                // Optionally, forward relevant headers from the incoming request
                'Content-Type': request.headers.get('Content-Type') || 'multipart/form-data',
            },
        });

        // Create a new Response object to return to the client
        const responseBody = await res.text(); // Use text() to handle any response type
        console.log("===> Response Body: ", responseBody)
        console.log("===> Response Header: ", res.Headers)
        const responseHeaders = new Headers();

        // Copy relevant headers from the Go API response, including Set-Cookie
        res.headers.forEach((value, key) => {
            if (key.toLowerCase() === 'set-cookie') {
                console.log(key + "===> " + value);
                responseHeaders.append('Set-Cookie', value); // Append cookies
            } else {
                responseHeaders.set(key, value); // Set other headers
            }
        });

        // Return a new Response with the same status, headers, and body
        return new Response(responseBody, {
            status: res.status,
            statusText: res.statusText,
            headers: responseHeaders,
        });
    } catch (error) {
        console.error("Error in register API route:", error);
        return new Response(
            JSON.stringify({ error: error.message }),
            {
                status: error.status || 500,
                headers: { 'Content-Type': 'application/json' },
            }
        );
    }
}