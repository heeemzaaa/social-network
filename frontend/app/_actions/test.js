// actions/test.js
"use server";

export async function fetchTestData() {
    try {
        const response = await fetch("http://localhost:8080/api/test", {
            cache: "no-store", // Ensure fresh data
        });
        if (!response.ok) {
            throw new Error(`API request failed with status ${response.status}`);
        }
        return response.json();
    } catch (error) {
        console.error("Error fetching data:", error);
        throw error;
    }
}