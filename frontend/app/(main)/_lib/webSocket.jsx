"use client";

import { useState, useEffect, useRef } from "react";
import { UserContext } from "../_context/userContext";

import { useNotification } from "../_context/NotificationContext";


export default function UserProvider({ children }) {
  const [messages, setMessages] = useState({});
  const [authenticatedUser, setAuthenticatedUser] = useState(null);
  const socketRef = useRef(null);

  const { showNotification } = useNotification();
  
  const handleNotificationSeen = (data) => data?.content !== "" ? showNotification({Content: data.content, Status: "info"}) : null

  const sendSocketMessage = (data) => {
    if (!socketRef.current) return console.warn("Socket not available");

    if (socketRef.current.readyState !== WebSocket.OPEN) return console.warn("Socket not connected");

    try {
      socketRef.current.send(JSON.stringify(data));
      console.log("✅ Socket message sent:", data);

    } catch (err) {
      console.error("❌ Socket send error:", err);
    }
  };

  useEffect(() => {
    const fetchLoggedInUser = async () => {
      try {
        const res = await fetch(`http://localhost:8080/api/loggedin`, {
          credentials: "include",
        });
        const data = await res.json();
        if (data.is_logged_in) {
          setAuthenticatedUser({
            id: data.id,
            username: data.nickname,
            fullName: data.fullname,
          });
        } else {
          setAuthenticatedUser(null);
        }
      } catch (err) {
        console.error(" Error fetching user:", err);
      }
    };
    fetchLoggedInUser();
  }, []);

  useEffect(() => {
    if (!authenticatedUser) return;

    const socket = new WebSocket(`ws://localhost:8080/ws/chat/`);
    socketRef.current = socket;

    socket.onopen = () => {
      console.log(" WebSocket connected");
    };

    socket.onmessage = (event) => {
      try {
        const data = JSON.parse(event.data);

        if (data.type === "notification") return handleNotificationSeen(data);

        if (typeof data.content === "string" && data.content !== "" && (data.type === "private" || data.type === "group")) {

          const isMe = data.sender_id === authenticatedUser.id;
          const chatKey = data.type === "group" ? data.target_id : isMe ? data.target_id : data.sender_id;

          const newMsg = {
            content: data.content,
            sender: isMe ? "me" : "them",
            createdAt: data.created_at,
            username: data.sender_name || data.receiver_name,
          };

          setMessages((prev) => ({ ...prev, [chatKey]: [...(prev[chatKey] || []), newMsg], }));
        }
      } catch (err) {
        console.error(" Failed to parse WebSocket message:", err);
      }
    };

    socket.onerror = (err) => {
      console.error(" WebSocket error:", err);
    };

    socket.onclose = () => {
      console.log(" WebSocket closed");
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
        sendSocketMessage,
        messages,
        setMessages,
        authenticatedUser,
      }}
    >
      {children}
    </UserContext.Provider>
  );
}
