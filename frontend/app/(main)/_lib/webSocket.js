"use client";

import { useState, useEffect, useRef } from "react";
import { UserContext } from "../_context/userContext";

export default function UserProvider({ children }) {
  const [users, setUsers] = useState([]);
  const [messages, setMessages] = useState({});
  const [authenticatedUser, setAuthenticatedUser] = useState(null);
  const socketRef = useRef(null);

  // STEP 1: Fetch logged-in user once on mount
  useEffect(() => {
    const fetchLoggedInUser = async () => {
      try {
        const res = await fetch("http://localhost:8080/api/loggedin", {
          credentials: "include",
        });
        const data = await res.json();
        console.log("✅ Logged in user:", data);

        if (data.is_logged_in) {
          setAuthenticatedUser({
            id: data.id,
            username: data.Nickname,
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

  // STEP 2: Setup WebSocket and fetch users after login
  useEffect(() => {
    if (!authenticatedUser) return;

    const socket = new WebSocket("ws://localhost:8080/ws/chat/");
    socketRef.current = socket;

    socket.onopen = () => {
      console.log("🟢 WebSocket connected");
    };

    socket.onmessage = (event) => {
      try {
        const data = JSON.parse(event.data);
        console.log("📨 Message received:", data);

        if (
          typeof data.content === "string" &&
          data.content !== "" &&
          (data.type === "private" || data.type === "group")
        ) {
          const isMe = data.sender_id === authenticatedUser.id;
          const chatPartner = isMe ? data.target_id : data.sender_id;

          const newMsg = {
            content: data.content,
            sender: isMe ? "me" : "them",
            createdAt: data.created_at,
          };

          setMessages((prev) => ({
            ...prev,
            [chatPartner]: [...(prev[chatPartner] || []), newMsg],
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

    // Fetch user list except current user
    const fetchUsers = async () => {
      try {
        const res = await fetch("http://localhost:8080/api/get-users/", {
          credentials: "include",
        });
        const usersList = await res.json();

        const filtered = usersList.filter((u) => u.id !== authenticatedUser.id);
        const mapped = filtered.map((user) => ({
          userID: user.id,
          username: user.firstname + " " + user.lastname,
          img: user.img || "/no-profile.png",
        }));

        setUsers(mapped);
        console.log("👥 Users list updated:", mapped);
      } catch (err) {
        console.error("❌ Error fetching users:", err);
      }
    };

    fetchUsers();

    return () => {
      socket.close();
      socketRef.current = null;
    };
  }, [authenticatedUser]);

  return (
    <UserContext.Provider
      value={{
        users,
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
