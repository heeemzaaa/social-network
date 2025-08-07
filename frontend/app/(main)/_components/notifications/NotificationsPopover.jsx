"use client";
import { useEffect, useState, useRef } from "react";
import { useNotification } from "../../_context/NotificationContext"; // Adjust path to match your project

export default function NotificationsPopover() {
  const [notifications, setNotifications] = useState([]);
  const [page, setPage] = useState(0);
  const [hasMore, setHasMore] = useState(true);
  const containerRef = useRef();

  const { showNotification } = useNotification();

  const notificationContent = (notification) => {
    // let content
    if (notification.Status == "later") {
      switch (notification.Type) {
        case "follow-private":
          return `${notification.SenderFullName} send a follow request`
        case "group-join":
          return `${notification.SenderFullName} would like to join ${notification.GroupName} group`
        case "group-invitation":
          return `${notification.SenderFullName} send a request to join ${notification.GroupName} group`
        case "group-event":
          return `${notification.SenderFullName} create event at ${notification.GroupName} group`
      }
    } else if (notification.Status == "accept") {
      switch (notification.Type) {
        case "follow-private":
          return `${notification.SenderFullName} follow you`
        case "group-join":
          return `${notification.SenderFullName} join your ${notification.GroupName} group`
        case "group-invitation":
          return `you are now a member of ${notification.GroupName} group`
        case "group-event":
          return `don't forget to go to ${notification.SenderFullName} event at ${notification.GroupName}` 
      }
    } else if (notification.Status == "reject") {
      switch (notification.Type) {
        case "follow-private":
          return `you rejected ${notification.SenderFullName} follow request`
        case "group-join":
          return `you refused ${notification.SenderFullName} to join your ${notification.GroupName} group`
        case "group-invitation":
          return `you rejected ${notification.SenderFullName} request to join ${notification.GroupName} group`
        case "group-event":
          return `you refused to go to ${notification.SenderFullName} event at ${notification.GroupName}` 
      }
    }
    return "content information !!"
  }

  // Handle accept/reject notification
  const handleNotificationAction = async (notification, status) => {

    if (!notification.Id) {
      console.error("No notification ID found. Available fields:", Object.keys(notification));
      return;
    }

    try {
      const postRequest = {
        method: "POST",
        credentials: "include",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify({
          NotifId: notification.Id,
          Status: status, // "accept" or "reject"
          Type: notification.Type,
          GroupId: notification.GroupId,
        })
      };

      let response = await fetch("http://localhost:8080/api/notifications/update/", postRequest);
      let data = await response.json();

      // Show popup with response message
      showNotification({
        Type: "response",
        Content: `Notification ${status}ed successfully: ${data?.Message}`,
        // Status: status === "accept" ? "success" : "error"
      });

      // Update local state
      setNotifications(prev =>
        prev.map(notif =>
          notif.Id === notification.Id ? { ...notif, Status: status } : notif
        )
      );

    } catch (error) {
      console.error(`Error ${status}ing notification:`, error);

      showNotification({
        Type: "error",
        Content: `Failed to ${status} notification: ${error.message}`,
        Status: "error"
      });
    }
  };

  // Load notifications
  useEffect(() => {
    loadNotifications(page);
  }, [page]);

  const loadNotifications = async (data_length) => {
    try {
      const res = await fetch(`http://localhost:8080/api/notifications?Count=${data_length}`, {
        method: "GET",
        credentials: "include"
      });
      const data = await res.json();

      if (data.length === 0) {
        setHasMore(false);
      } else {
        // Filter out duplicates based on ID
        setNotifications((prev) => {
          const existingIds = new Set(prev.map(notif => notif.Id));
          const newNotifications = data.filter(notif => !existingIds.has(notif.Id));
          return [...prev, ...newNotifications];
        });
      }
    } catch (error) {
      console.error("Error fetching notifications:", error);
      setHasMore(false);
    }
  };

  // Scroll handler
  const handleScroll = () => {
    if (!containerRef.current || !hasMore) return;

    const { scrollTop, scrollHeight, clientHeight } = containerRef.current;

    if (scrollTop + clientHeight >= scrollHeight - 10) {
      setPage((prev) => prev + 10);
    }
  };

  return (
    <div
      ref={containerRef}
      onScroll={handleScroll}
      style={{ maxHeight: "350px", overflowY: "auto", width: "300px" }}
      className="bg-white shadow p-2 rounded"
    >
      {notifications.length === 0 && <p>No notifications</p>}

      {notifications.map((notif) => (
        <div
          key={notif.Id}
          className={`notification-card ${notif.Type} ${notif.Status} ${notif.Seen ? "seen" : "unseen"}`}
        >
          <p>{notificationContent(notif)}</p>

          {notif.Status === "later" && notif.Type !== "follow-public" && (
            <div className="action-buttons">
              <button
                className="accept-btn"
                onClick={() => handleNotificationAction(notif, "accept")}
              >
                ✔
              </button>
              <button
                className="reject-btn"
                onClick={() => handleNotificationAction(notif, "reject")}
              >
                ✘
              </button>
            </div>
          )}

          {notif.Status === "accept" && <div className="green-dote" />}
          {notif.Status === "reject" && <div className="red-dote" />}
        </div>
      ))}

      {hasMore && <p className="text-center text-gray-400 text-xs">Loading more...</p>}
      {!hasMore && notifications.length > 0 && (
        <p className="text-center text-gray-400 text-xs">No more notifications</p>
      )}
    </div>
  );
}
