export async function POST(request) {

    console.info("inside the register api route ®️")

    try {
        let formData = await request.formData();
        const res = await fetch("http://localhost:8080/api/auth/register", {
            method: "POST",
            body: formData // Forward the request body
        });
        
        return res
    } catch (error) {
        return Response.json(
            { error: error.message },
            { status: error.status || 500 }
        );
    }
}