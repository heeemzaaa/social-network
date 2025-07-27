"use client";

import React, { useState, useEffect, useRef } from "react";
import "./chat.css";
import Button from "@/app/_components/button";
import { HiMiniFaceSmile, HiPaperAirplane } from "react-icons/hi2";
import UserList from "../_components/chat/user_list";
import GroupList from "../_components/group_list";
import { useUserContext } from "../_context/userContext";
import { fetchMessages } from "../_components/fetchMessages";

const emojis = ["😀", "😂", "😍", "🔥", "🥺", "👍", "❤️", "🎉"];

export default function Chat() {
  const usersBlock = document.getElementsByClassName("user_groups_place");
  const chatBlock = document.getElementsByClassName("chat_place");
  const [chatBodyName, setChatBodyName] = useState("");
  const [chatTarget, setChatTarget] = useState(null);
  const [newMessage, setNewMessage] = useState("");
  const { users, socket, messages, setMessages, authenticatedUser, groups } =
    useUserContext();
  const [showEmojiPicker, setShowEmojiPicker] = useState(false);
  const bottomRef = useRef(null);
  const [view, setView] = useState("Users");

  useEffect(() => {
    if (!chatTarget?.ID || !authenticatedUser) return;

    const loadMessages = async () => {
      const msgs = await fetchMessages(chatTarget.ID, chatTarget.type);
      if (!msgs) return;

      setMessages((prev) => ({
        ...prev,
        [chatTarget.ID]: msgs.map((msg) => {
          const isMe = msg.sender_id === authenticatedUser.id;
          return {
            content: msg.content,
            sender: isMe ? "me" : "them",
            createdAt: msg.created_at,
            username: msg.sender_name,
          };
        }),
      }));
    };

    loadMessages();
  }, [chatTarget?.ID, chatTarget?.type, authenticatedUser]);

  useEffect(() => {
    if (
      chatTarget?.type === "private" &&
      chatTarget?.ID &&
      !users.some((user) => user.userID === chatTarget.ID)
    ) {
      setChatTarget(null);
      setChatBodyName("");
    }
  }, [users, chatTarget]);

  const handleUserClick = (user) => {
    setChatTarget({
      ID: user.userID,
      type: "private",
    });
    if (window.innerWidth <= 500) {
      chatBlock[0].style.display = "flex";
      usersBlock[0].style.display = "none";
    }

    setChatBodyName(user.username);
  };

  const handleGroupClick = (group) => {
    setChatTarget({
      ID: group.group_id,
      type: "group",
    });
    if (window.innerWidth <= 500) {
      chatBlock[0].style.display = "flex";
      usersBlock[0].style.display = "none";
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

  const handleEmojiClick = (emojiData) => {
    setNewMessage((prev) => prev + emojiData);
  };

  useEffect(() => {
    bottomRef.current?.scrollIntoView({ behavior: "smooth" });
  }, [messages]);

  function back() {
    usersBlock[0].style.display = "flex";
    chatBlock[0].style.display = "none";
  }
  useEffect(() => {
    const handleResize = () => {
      if (window.innerWidth >= 500) {
        chatBlock[0].style.display = "flex";
        usersBlock[0].style.display = "flex";
      } else {
        chatBlock[0].style.display = "none";
        usersBlock[0].style.display = "flex";
      }
    };

    handleResize();

    window.addEventListener("resize", handleResize);

    return () => {
      window.removeEventListener("resize", handleResize);
    };
  }, []);

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
            <GroupList groups={groups} onGroupClick={handleGroupClick} />
          )}
        </div>
      </section>

      <section className="chat_place flex-col">
        {chatBodyName ? (
          <div className="chat_header p2">
            <div
              className="goBack cursor-pointer w-8 h-8"
              onClick={() => back()}
              title="Go back"
            >
              <svg
                fill="#1E201F"
                height="5vh"
                width="5vw"
                viewBox="0 0 206.108 206.108"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  d="M152.774,69.886H30.728l24.97-24.97c3.515-3.515,3.515-9.213,0-12.728c-3.516-3.516-9.213-3.515-12.729,0L2.636,72.523 
        c-3.515,3.515-3.515,9.213,0,12.728l40.333,40.333c1.758,1.758,4.061,2.636,6.364,2.636c2.303,0,4.606-0.879,6.364-2.636 
        c3.515-3.515,3.515-9.213,0-12.728l-24.97-24.97h122.046c19.483,0,35.334,15.851,35.334,35.334s-15.851,35.334-35.334,35.334
        H78.531c-4.971,0-9,4.029-9,9s4.029,9,9,9h74.242c29.408,0,53.334-23.926,53.334-53.334S182.182,69.886,152.774,69.886z"
                />
              </svg>
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
                <div ref={bottomRef} />
              </div>
            ))}
          </div>
        ) : (
          <div className="chat_body empty">
            <p className="text-gray-500">
              NSIT DAK SITE DYAL TSAWR HET HNA TSWIRA BROJOLA HHHH
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
              onKeyDown={(e) =>
                e.key === "Enter" && !e.shiftKey && sendMessage()
              }
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
