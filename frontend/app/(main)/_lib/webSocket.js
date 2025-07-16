"use client";

import { useState, useEffect, useRef } from "react";
import { UserContext } from "../_context/userContext";

export default function UserProvider({ children }) {
  const [users, setUsers] = useState([]);
  const [messages, setMessages] = useState({});
  const [authenticatedUser, setAuthenticatedUser] = useState(null);
  const socketRef = useRef(null);

  // Step 1: Fetch the logged-in user
  useEffect(() => {
    const fetchLoggedInUser = async () => {
      try {
        const res = await fetch("http://localhost:8080/api/loggedin", {
          credentials: "include",
        });
        const data = await res.json();
        console.log("🚀 Fetched logged-in user:", data);

        if (data.is_logged_in) {
          setAuthenticatedUser({
            id: data.Id,
            username: data.Nickname,
          });
        } else {
          console.warn("🚫 User not logged in");
          setAuthenticatedUser(null); // explicitly clear on no login
        }
      } catch (err) {
        console.error("❌ Error fetching logged-in user:", err);
        setAuthenticatedUser(null);
      }
    };

    fetchLoggedInUser();
  }, []);

  // Step 2: Setup WebSocket and users/messages after user is fetched
  useEffect(() => {
    if (!authenticatedUser) {
      // Clear users and messages when no authenticated user
      setUsers([]);
      setMessages({});
      if (socketRef.current) {
        socketRef.current.close();
        socketRef.current = null;
      }
      return;
    }

    // Initialize WebSocket connection
    socketRef.current = new WebSocket("ws://localhost:8080/ws/chat/");

    socketRef.current.onopen = () => {
      console.log("✅ WebSocket connected");
      // Optional: send initial message here
    };

    socketRef.current.onmessage = (event) => {
      try {
        const data = JSON.parse(event.data);

        // Update online users list
        /*if (data.type === "online" && Array.isArray(data.data)) {
          const newUsers = data.data
            .map((user) => ({
              userID: user.id,
              username: user.firstname + " " + user.lastname,
              online: true,
            }));

          setUsers(newUsers);
        }*/

        // Update messages state
        if (
          typeof data.content === "string" &&
          data.content !== "" &&
          (data.type === "private" || data.type === "group") &&
          (data.sender_id === authenticatedUser.id)
        ) {
          const from = data.sender_id;
          const msg = {
            content: data.content,
            sender: from === authenticatedUser.id ? "me" : "them",
          };

          setMessages((prev) => ({
            ...prev,
            [from]: [...(prev[from] || []), msg],
          }));
        }
      } catch (err) {
        console.error("❌ Error parsing WebSocket message:", err);
      }
    };

    socketRef.current.onerror = (err) => {
      console.error("❌ WebSocket error:", err);
    };

    socketRef.current.onclose = (event) => {
      console.log(`🔌 WebSocket closed. Code: ${event.code}`);
    };

    // Fetch users initially via REST
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
        }));
        setUsers(mapped);
        console.log("Updated users list:", mapped);
      } catch (err) {
        console.error("❌ Error fetching users:", err);
      }
    };

    fetchUsers();

    // Cleanup on component unmount or user change
    return () => {
      socketRef.current?.close();
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
