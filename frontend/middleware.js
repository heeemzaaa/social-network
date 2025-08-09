
// import { NextResponse } from 'next/server';

import { NextResponse } from "next/server";

export async function middleware(request) {
    console.log("heeere", `${process.env.NEXT_PUBLIC_API_URL}`);
    console.log(` the endpoint is : ${process.env.NEXT_PUBLIC_API_URL}/api/loggedin`);
    try {
        // Fetch authentication status from the external API
        const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/loggedin`, {
            headers: {
                Cookie: request.headers.get('cookie'), // Forward client cookies to API
            },
            credentials: 'include', // Required for cross-origin cookie handling
        });

        if (!response.ok) {

            console.error('API request failed with status:', response.status);
            return NextResponse.redirect(new URL('/login', request.url));
        }



        const data = await response.json();
        console.log("response", data);
        const isLoggedIn = data.is_logged_in; // Adjust if API response format differs
        // Handle redirection based on route and authentication status
        console.log("hnaaaaa", request.nextUrl.pathname);
        if (request.nextUrl.pathname === '/login' || request.nextUrl.pathname === '/register') {
            if (isLoggedIn) {
                return NextResponse.redirect(new URL('/', request.url));
            }
        } else {
            if (!isLoggedIn) {
                return NextResponse.redirect(new URL('/login', request.url));
            }
        }
        return NextResponse.next();
    } catch (error) {
        console.error('Middleware error:', error);
        return NextResponse.redirect(new URL('/login', request.url));
    }

}

// Apply middleware to all routes except API and static files
export const config = {
    matcher: '/((?!api|_next/static|_next/image|favicon.ico|social_network_logo).*)',
};