
// frontend/components/Notification.jsx
import React from "react";

export default function Notification({ notif, isPopup = false }) {
  const { Type, Status } = notif;

  // Build dynamic className
  const notifClass = `notification-card ${Type} ${Status || ""} ${isPopup ? "popup" : ""}`;

  return (
    <div className={notifClass}>
      <p className="notif-content">{notif.Content}</p>
      <small className="notif-type">Type: {Type}</small>

      {Type === "group-invitation" && (
        <div className="action-buttons">
          <button className="accept-btn">Accept</button>
          <button className="reject-btn">Decline</button>
        </div>
      )}
    </div>
  );
}
