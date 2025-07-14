"use client";

import { useEffect, useRef } from "react";
import { userList } from "@/lib/global";

let users = []

export default function TestChat() {
  const hasConnected = useRef(false);

  useEffect(() => {
    if (hasConnected.current) return;
    hasConnected.current = true;

    const socket = new WebSocket("ws://localhost:8080/ws/chat");

    socket.onopen = () => {
      console.log("✅ WebSocket connected");

      const testMessage = {
        content: "Just testing",
        receiver_id: "dc079e8c-0705-4969-b6a4-3fd5cc4d7e04",
        type: "private",
      };

      socket.send(JSON.stringify(testMessage));
      console.log("📤 Sent:", testMessage);
    };

    socket.onmessage = (event) => {
      try {
        const data = JSON.parse(event.data);
		if (data.type === "online" && Array.isArray(data.data)) {
			for (let i = 0; i < data.data.length; i++) {
				let holder = data.data[i]
				const newUser = { ...userList };
				if (data.type === "online") {
					newUser.online = true
				} else {
					newUser.online = false
				}
				newUser.userID = holder.id
				newUser.username = holder.firstname + " " + holder.lastname
				users.push(newUser)
				console.log("*/****: ", data.data[i])
			}
			console.log("list: ", users	)
			console.log("🟢 Online users updated:", data.data[0]);
			console.log("🟢 Oenline users updated:", data.data[1]);
		}
        console.log("📥 Received:", data);
      } catch (err) {
        console.warn("⚠️ Invalid JSON received:", event.data);
      }
    };

    socket.onerror = (err) => {
      console.error("❌ WebSocket error:", err);
    };

    socket.onclose = () => {
      console.log("🔌 WebSocket connection closed");
    };

    return () => {
      socket.close();
    };
  }, []);

  return null; // Since this component doesn't render anything visible
}
