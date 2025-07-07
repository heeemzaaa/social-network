export async function GET(request) {
    // console.log(request)
    try {
        const res = await fetch("http://localhost:8080/api/test");
        if (!res.ok) throw new Error("API request failed", res.status);
        const data = await res.json();
        return Response.json(data);
    } catch (error) {
        return Response.json({ error: `${error.message}` }, { status: 500 });
    }
}