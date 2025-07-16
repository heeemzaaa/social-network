
export async function fetchMessages(targetId, type) {
  try {
    const response = await fetch(`http://localhost:8080/api/messages?target_id=${targetId}&type=${type}`, {
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
    return messages || [];
  } catch (error) {
    console.error("❌ Error fetching messages:", error);
    return [];
  }
}