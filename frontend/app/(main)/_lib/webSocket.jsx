"use client";

import { useState, useEffect, useRef } from "react";
import { UserContext } from "../_context/userContext";

export default function UserProvider({ children }) {
  const [messages, setMessages] = useState({});
  const [authenticatedUser, setAuthenticatedUser] = useState(null);
  const socketRef = useRef(null);

  useEffect(() => {
    const fetchLoggedInUser = async () => {
      try {
        console.log("wwwwwwwwwwwwwwwwwwwwwwwayli",`${process.env.NEXT_PUBLIC_API_URL}`);
        console.log("data user", `${process.env.NEXT_PUBLIC_API_URL}/api/loggedin`);
        const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/loggedin`, {
          credentials: "include",
        });
        const data = await res.json();
        console.log("✅ Logged in user:", data);
        if (data.is_logged_in) {
          setAuthenticatedUser({
            id: data.id,
            username: data.nickname,
            fullName: data.fullname,
          });
        } else {
          setAuthenticatedUser(null);
          console.warn("🚫 User not logged in");
        }
        console.log("Authenticated user set:", data.id);
      } catch (err) {
        console.error("❌ Error fetching user:", err);
      }
    };
    fetchLoggedInUser();
  }, []);

  useEffect(() => {
    if (!authenticatedUser) return;

    const socket = new WebSocket(`${process.env.NEXT_PUBLIC_WS_URL}/ws/chat/`);
    socketRef.current = socket;

    socket.onopen = () => {
      console.log("🟢 WebSocket connected");
    };

    socket.onmessage = (event) => {
      try {
        const data = JSON.parse(event.data);

        if (
          typeof data.content === "string" &&
          data.content !== "" &&
          (data.type === "private" || data.type === "group")
        ) {
          const isMe = data.sender_id === authenticatedUser.id;

          const chatKey =
            data.type === "group"
              ? data.target_id
              : isMe
                ? data.target_id
                : data.sender_id;

          const newMsg = {
            content: data.content,
            sender: isMe ? "me" : "them",
            createdAt: data.created_at,
            username: data.sender_name || data.receiver_name,
          };

          setMessages((prev) => ({
            ...prev,
            [chatKey]: [...(prev[chatKey] || []), newMsg],
          }));
        }
      } catch (err) {
        console.error("❌ Failed to parse WebSocket message:", err);
      }
    };

    socket.onerror = (err) => {
      console.error("❌ WebSocket error:", err);
    };

    socket.onclose = () => {
      console.log("🔌 WebSocket closed");
    };
    return () => {
      socket.close();
      socketRef.current = null;
    };
  }, [authenticatedUser]);

  return (
    <UserContext.Provider
      value={{
        socket: socketRef.current,
        messages,
        setMessages,
        authenticatedUser,
      }}
    >
      {children}
    </UserContext.Provider>
  );
}
