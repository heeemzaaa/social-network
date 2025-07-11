export async function GET(req, response) {
    // const cookieHeader = req.headers.cookie || '';
    // console.log('Test API: Cookies received:', cookieHeader);
    try {
        const response = await fetch('http://localhost:8080/api/auth/islogged');
        const data = await response.json();
        console.log('Test API: API response:', data);

        return Response.json({ status: response.status, data });

    } catch (error) {
        console.error('Test API: Error:', error);
        return Response.json({ error: 'Failed to fetch auth status' });

    }
}