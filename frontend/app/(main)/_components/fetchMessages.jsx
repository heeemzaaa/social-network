
const API_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080'


export async function fetchMessages(targetId, type) {
  try {
    const response = await fetch(`${API_URL}/api/messages?target_id=${targetId}&type=${type}`, {
      cache: "no-store",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
    });
    if (!response.ok) {
      throw new Error(`HTTP error! Status: ${response.status}`);
    }

    const messages = await response.json();
    if (Array.isArray(messages)) {
      messages.reverse();
    }
    return messages || [];
  } catch (error) {
    console.error("❌ Error fetching messages:", error);
    return [];
  }
}