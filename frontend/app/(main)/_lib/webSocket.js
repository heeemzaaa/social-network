"use client";

import { useState, useEffect, useRef } from "react";
import { UserContext } from "../_context/userContext";
import { userList } from "@/lib/global";

export default function UserProvider({ children }) {
  const [users, setUsers] = useState([]);
  const hasConnected = useRef(false);

  useEffect(() => {
    if (hasConnected.current) return;
    hasConnected.current = true;

    const init = async () => {
      try {
        const socket = new WebSocket("ws://localhost:8080/ws/chat/");
        const response = await fetch("http://localhost:8080/api/get-users/" , {
          cache: "no-store", // Ensure fresh data
          headers: {
            "Content-Type": "application/json",
          },
          credentials: "include", // Include cookies for authentication
          });
        
        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`);
        }

        const usersList = await response.json();
        console.log("🔌 Initializing WebSocket connection. Users list:", usersList);

        socket.onopen = () => {
          console.log("✅ WebSocket connected");

          const testMessage = {
            content: "Just testing",
            receiver_id: "dc079e8c-0705-4969-b6a4-3fd5cc4d7e04",
            type: "private",
          };

          socket.send(JSON.stringify(testMessage));
        };

        socket.onmessage = (event) => {
          try {
            const data = JSON.parse(event.data);
            if (data.type === "online" && Array.isArray(data.data)) {
              const newUsers = data.data.map((u) => ({
                ...userList,
                userID: u.id,
                username: u.firstname + " " + u.lastname,
                online: true,
              }));
              setUsers(newUsers);
              console.log("🟢 Online users updated:", newUsers);
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
          console.log("🔌 WebSocket closed");
        };

        // Clean up
        return () => socket.close();
      } catch (err) {
        console.error("❌ Error in init:", err);
      }
    };

    init();
  }, []);

  return (
    <UserContext.Provider value={{ users }}>
      {children}
    </UserContext.Provider>
  );
}
