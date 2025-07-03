export async function POST(request) {

    console.info("inside the register api route ®️")

    try {
        
        let formData = await request.formData();
        console.log("request form data aaaaaaaaaaaaaaaaa: ", formData)

        const res = await fetch("http://localhost:8080/api/auth/register", {
            method: "POST",
            body: await request.body // Forward the request body
        });

        if (!res.ok) {
            throw new Error(`API request failed with status ${res.status}`);
        }
        const data = await res.json();
        return Response.json(data);
    } catch (error) {
        console.error("Error in API route: 1", error.message);
        return Response.json(
            { error: error.message },
            { status: error.status || 500 }
        );
    }
}