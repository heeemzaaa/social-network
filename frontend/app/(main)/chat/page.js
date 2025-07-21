"use client";

import React, { useState, useEffect } from "react";
import "./chat.css";
import Button from "@/app/_components/button";
import { HiMiniFaceSmile, HiPaperAirplane } from "react-icons/hi2";
import UserList from "../_components/chat/user_list";
import GroupList from "../_components/group_list";
import { useUserContext } from "../_context/userContext";
import { fetchMessages } from "../_components/fetchMessages";
import { create } from "domain";
const emojis = ["😀", "😂", "😍", "🔥", "🥺", "👍", "❤️", "🎉"];

export default function Chat() {
  const [currentUser, setCurrentUser] = useState({
    username: "",
    ID: "",
    type: "private",
  });
  const [newMessage, setNewMessage] = useState("");
  const { users, socket, messages, setMessages, authenticatedUser } =
    useUserContext();
  const [showEmojiPicker, setShowEmojiPicker] = useState(false);

  const groups = {
    groups: [
      { name: "grp1", img: "" },
      { name: "grp2", img: "" },
      { name: "grp3", img: "" },
      { name: "grp4", img: "" },
      { name: "grp5", img: "" },
      { name: "grp6", img: "" },
      { name: "grp7", img: "" },
    ],
  };

  const [view, setView] = useState("Users");

  // Reset currentUser if they disappear from users list
  useEffect(() => {
    if (
      currentUser.ID &&
      !users.some((user) => user.userID === currentUser.ID)
    ) {
      setCurrentUser({ username: "", ID: "", type: "private" });
    }
  }, [users, currentUser.ID]);

  // Load message history when currentUser changes
  useEffect(() => {
    if (!currentUser.ID || !authenticatedUser) return;

    const loadMessages = async () => {
      const msgs = await fetchMessages(currentUser.ID, currentUser.type);

      if (!msgs) return;

      setMessages((prev) => ({
        ...prev,
        [currentUser.ID]: msgs.map((msg) => {
          const isMe = msg.sender_id === authenticatedUser.id;
          return {
            content: msg.content,
            sender: isMe ? "me" : "them",
            createdAt: msg.created_at,
            username: isMe ? msg.sender_name : msg.receiver_name,
          };
        }),
      }));
    };

    console.log("📤 Message sent:", messages);
    loadMessages();
  }, [currentUser.ID, authenticatedUser]); // Added authenticatedUser as dependency

  const handleUserClick = (user) => {
    setCurrentUser({
      username: user.username,
      ID: user.userID,
      type: "private",
    });
  };

  // Send message to backend via WebSocket and update messages locally optimistically
  const sendMessage = () => {
    if (!newMessage || !currentUser.ID || socket?.readyState !== 1) return;

    const messagePayload = {
      content: newMessage,
      target_id: currentUser.ID,
      type: "private",
    };

    // Send to backend (which will verify, save, then broadcast)
    socket.send(JSON.stringify(messagePayload));

    // Optimistically add message to current chat as "me"

    setNewMessage("");
  };

  const handleEmojiClick = (emojiData) => {
    setNewMessage((prev) => prev + emojiData);
  };

  return (
    <main className="chat_main_container p4 flex-row">
      <section className="user_groups_place h-full flex-col">
        <div className="user_groups_choosing flex-row justify-center align-center">
          <Button
            onClick={() => setView("Users")}
            variant={view === "Users" ? "btn-primary" : "btn-secondary"}
          >
            Users
          </Button>

          <Button
            onClick={() => setView("Groups")}
            variant={view === "Groups" ? "btn-primary" : "btn-secondary"}
            className="p4"
          >
            Groups
          </Button>
        </div>

        <div className="chosing_param">
          {view === "Users" ? (
            <UserList users={users} onUserClick={handleUserClick} />
          ) : (
            <GroupList {...groups} />
          )}
        </div>
      </section>

      <section className="chat_place flex-col">
        <div className="chat_header p2">
          <img src="/no-profile.png" alt="Profile" />
          <p className="text-lg font-semibold">{currentUser.username}</p>
        </div>

        <div className="chat_body">
          {(messages[currentUser.ID] || []).map((msg, i) => (
            <div
              key={i}
              className={`message ${msg.sender === "me" ? "sent" : "received"}`}
            >
              {msg.username && <span className="username">{msg.username}</span>}
              {msg.content}
              <span className="timestamp">
                {new Date(msg.createdAt).toLocaleTimeString([], {
                  hour: "2-digit",
                  minute: "2-digit",
                })}
              </span>
            </div>
          ))}
        </div>

        <div className="chat_footer p2">
          <div style={{ position: "relative" }}>
            <HiMiniFaceSmile
              size={"30px"}
              onClick={() => setShowEmojiPicker((prev) => !prev)}
              style={{ cursor: "pointer" }}
            />

            {showEmojiPicker && (
              <div className="emoji-picker">
                {emojis.map((emoji, index) => (
                  <span
                    key={index}
                    className="emoji"
                    onClick={() => handleEmojiClick(emoji)}
                  >
                    {emoji}
                  </span>
                ))}
              </div>
            )}
          </div>

          <textarea
            value={newMessage}
            onChange={(e) => setNewMessage(e.target.value)}
            onKeyDown={(e) => e.key === "Enter" && !e.shiftKey && sendMessage()}
            placeholder="Type a message..."
          />
          <HiPaperAirplane
            onClick={sendMessage}
            className="HiPaperAirplane"
            size={"30px"}
            style={{ cursor: "pointer" }}
          />
        </div>
      </section>
    </main>
  );
}
