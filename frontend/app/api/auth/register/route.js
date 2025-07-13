// frontend/app/api/auth/register/route.js
export async function POST(request) {
    try {
        // Get the form data from the incoming request
        const formData = await request.formData()
        
        // Forward the request to the Go API
        const res = await fetch("http://localhost:8080/api/auth/register", {
            method: "POST",
            body: formData,
            credentials: 'include',
        })
        return res
    } catch (error) {
        console.error("Error in register API route: ", error)
    }
}