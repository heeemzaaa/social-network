// "use server"
// import { cookies } from "next/headers"

// export async function getUserId() {
//     try {
//         console.log("********************************")
//        const sessionCookie = await cookies().get("session")?.value;
//         const res = await fetch("http://localhost:8080/api/auth/islogged", {
//             method: "GET",
//             credentials: 'include',
//             headers: sessionCookie ? { Cookie: `session=${sessionCookie}` } : {}
//         });
//         if (!res.ok) {
//             console.error("Failed to fetch user ID. Status:", res.status);
//             return null;
//         }
        
//         const data = await res.json();
//         console.log("User data:", data);

//         return data?.UserId || null;
//     } catch (error) {
//         console.error("Error in getting user ID:", error);
//         return null;
//     }
// }
