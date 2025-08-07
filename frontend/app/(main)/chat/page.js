"use client";

import React, { useState, useEffect, useRef, useCallback } from "react";
import "./chat.css";
import Button from "@/app/_components/button";
import { HiMiniFaceSmile, HiPaperAirplane } from "react-icons/hi2";
import UserList from "../_components/chat/user_list";
import GroupList from "../_components/group_list";
import { useUserContext } from "../_context/userContext";
import { fetchMessages } from "../_components/fetchMessages";
import { SlActionUndo } from "react-icons/sl";

const emojis = ["😀", "😂", "😍", "🔥", "🥺", "👍", "❤️", "🎉"];

export default function Chat() {
  const [users, setUsers] = useState([]);
  const [groups, setGroups] = useState([]);
  const usersBlockRef = useRef(null);
  const chatBlockRef = useRef(null);
  const [chatBodyName, setChatBodyName] = useState("");
  const [chatTarget, setChatTarget] = useState(null);
  const [newMessage, setNewMessage] = useState("");
  const { socket, messages, setMessages, authenticatedUser } = useUserContext();

  const [showEmojiPicker, setShowEmojiPicker] = useState(false);
  const bottomRef = useRef(null);
  const [view, setView] = useState("Users");

  // fetch users
  const fetchUsers = useCallback(async () => {
    try {
      const res = await fetch("http://localhost:8080/api/get-users/", {
        credentials: "include",
      });
      const usersList = await res.json();

      const mapped = usersList.map((user) => ({
        userID: user.id,
        username: user.fullname,
        img: user.img || "/no-profile.png",
      }));
      setUsers(mapped);
    } catch (err) {
      console.error("❌ Error fetching users:", err);
    }
  }, []);

  // fetch groups
  const fetchGroup = useCallback(async () => {
    try {
      const res = await fetch("http://localhost:8080/api/get-groups/", {
        credentials: "include",
      });
      const groupList = await res.json();

      const mappedG = groupList.map((group) => ({
        group_id: group.group_id,
        title: group.title,
        image_path: group.image_path || "/no-profile.png",
      }));

      setGroups(mappedG);
    } catch (err) {
      console.error("❌ Error fetching groups:", err);
    }
  }, []);

  useEffect(() => {
    fetchUsers();
    fetchGroup();
  }, [fetchUsers, fetchGroup]);

  // load targetUser messages
  useEffect(() => {
    if (!chatTarget?.ID || !authenticatedUser) return;

    const loadMessages = async () => {
      const msgs = await fetchMessages(chatTarget.ID, chatTarget.type);
      console.log("messages: ", msgs);
      if (!msgs) return;

      setMessages({
        [chatTarget.ID]: msgs.map((msg) => {
          const isMe = msg.sender_id === authenticatedUser.id;
          return {
            content: msg.content,
            sender: isMe ? "me" : "them",
            createdAt: msg.created_at,
            username: msg.sender_name,
          };
        }),
      });
    };

    loadMessages();
  }, [chatTarget?.ID, chatTarget?.type, authenticatedUser]);

  // handle user click selection
  const handleUserClick = (user) => {
    setChatTarget({
      ID: user.userID,
      type: "private",
    });
    if (window.innerWidth <= 500) {
      chatBlockRef.current.style.display = "flex";
      usersBlockRef.current.style.display = "none";
    }

    setChatBodyName(user.username);
  };

  // handle group click selection
  const handleGroupClick = (group) => {
    setChatTarget({
      ID: group.group_id,
      type: "group",
    });
    if (window.innerWidth <= 500) {
      chatBlockRef.current.style.display = "flex";
      usersBlockRef.current.style.display = "none";
    }
    setChatBodyName(group.title);
  };

  const sendMessage = () => {
    if (!newMessage.trim() || !chatTarget?.ID || socket?.readyState !== 1)
      return;

    const messagePayload = {
      content: newMessage,
      target_id: chatTarget.ID,
      type: chatTarget.type,
    };

    socket.send(JSON.stringify(messagePayload));
    setNewMessage("");
  };

  // handle emojies click selection
  const handleEmojiClick = (emojiData) => {
    setNewMessage((prev) => prev + emojiData);
    setShowEmojiPicker(false);
  };

  useEffect(() => {
    bottomRef.current?.scrollIntoView({ behavior: "smooth" });
  }, [messages]);

  const back = useCallback(() => {
    chatBlockRef.current.style.display = "none";
    usersBlockRef.current.style.display = "flex";
  }, []);

  // for responive
  useEffect(() => {
    const usersBlock = usersBlockRef.current;
    const chatBlock = chatBlockRef.current;

    const handleResize = () => {
      if (chatBlock && usersBlock) {
        if (window.innerWidth >= 500) {
          chatBlock.style.display = "flex";
          usersBlock.style.display = "flex";
        } else {
          chatBlock.style.display = "none";
          usersBlock.style.display = "flex";
        }
      }
    };

    handleResize();
    window.addEventListener("resize", handleResize);
    return () => window.removeEventListener("resize", handleResize);
  }, []);

  return (
    <main className="chat_main_container p4 flex-row">
      <section
        className="user_groups_place h-full flex-col"
        ref={usersBlockRef}
      >
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
            <GroupList groups={groups} onGroupClick={handleGroupClick} />
          )}
        </div>
      </section>

      <section className="chat_place flex-col" ref={chatBlockRef}>
        {chatBodyName ? (
          <div className="chat_header p2">
            <div
              className="goBack cursor-pointer w-8 h-8"
              onClick={back}
              title="Go back"
            >
              <SlActionUndo />
            </div>
            <img src="/no-profile.png" alt="Profile" />
            <p className="text-lg font-semibold">{chatBodyName}</p>
          </div>
        ) : (
          <div className="chat_header p2 text-gray-500 italic">
            No chat selected
          </div>
        )}

        {chatBodyName ? (
          <div className="chat_body">
            {(messages[chatTarget?.ID] || []).map((msg, i) => (
              <div
                key={i}
                className={`message ${
                  msg.sender === "me" ? "sent" : "received"
                }`}
              >
                {msg.username && (
                  <span className="username">{msg.username}</span>
                )}
                {msg.content}
                <span className="timestamp">
                  {new Date(msg.createdAt).toLocaleTimeString([], {
                    hour: "2-digit",
                    minute: "2-digit",
                  })}
                </span>
              </div>
            ))}
            <div ref={bottomRef} />
          </div>
        ) : (
          <div className="chat_body empty">
            <p className="text-gray-500">
              Select a user or group to start chatting!
            </p>
          </div>
        )}

        {chatTarget && (
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
              onKeyDown={(e) => {
                if (e.key === "Enter" && !e.shiftKey) {
                  e.preventDefault();
                  sendMessage();
                }
              }}
              placeholder="Type a message..."
            />
            <HiPaperAirplane
              onClick={sendMessage}
              className="HiPaperAirplane"
              size={"30px"}
              style={{ cursor: "pointer" }}
            />
          </div>
        )}
      </section>
    </main>
  );
}
