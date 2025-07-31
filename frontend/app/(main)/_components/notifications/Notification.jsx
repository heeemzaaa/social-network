// components/Notification.jsx
"use client";
import React from "react";

export default function Notification({ notif, isPopup = false, onConfirm }) {

  const {Type, Status, Content, isConfirm, confirmButtons = { accept: "accept", reject: "reject" }} = notif; // should be add time !!!

  const notifClass = `notification-card ${Type} ${Status || ""} ${isPopup ? "popup" : ""}`;

  return (
    <div className={notifClass}>
      <p className="notif-content">{Content}</p>

      {/* If confirm popup */}
      {isConfirm && (
        <div className="action-buttons">
          <button className="accept-btn" onClick={() => onConfirm(true)}>
            {confirmButtons.accept}
          </button>
          <button className="reject-btn" onClick={() => onConfirm(false)}>
            {confirmButtons.reject}
          </button>
        </div>
      )}
    </div>
  );
}
