// app/api/singleNotif/route.js

export async function GET() {
    try {
        const response = await fetch("http://localhost:8080/api/notification")
    } catch (error) {
        console.error("Error in API single notifications:", error.message);
    }
}