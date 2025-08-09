"use client";
import { useEffect, useState, useRef } from "react";
import { useNotification } from "../../_context/NotificationContext";
import "./styles.css";

export default function NotificationsPopover() {
  const containerRef = useRef();
  
  const [notifications, setNotifications] = useState([]);
  const [notifId, setNotifId] = useState("0");
  
  const [hasMore, setHasMore] = useState(true);
  const [isLoading, setIsLoading] = useState(false);

  const { showNotification } = useNotification();

  const notificationContent = (notification) => {
    if (notification.Status == "later") {
      switch (notification.Type) {
        case "follow-private":
          return `${notification.SenderFullName} send a follow request`;
        case "group-join":
          return `${notification.SenderFullName} would like to join ${notification.GroupName} group`;
        case "group-invitation":
          return `${notification.SenderFullName} send a request to join ${notification.GroupName} group`;
      }
    } else if (notification.Status == "accept") {
      switch (notification.Type) {
        case "follow-private":
        case "follow-public":
          return `${notification.SenderFullName} follow you`;
        case "group-join":
          return `${notification.SenderFullName} join your ${notification.GroupName} group`;
        case "group-invitation":
          return `you are now a member of ${notification.GroupName} group`;
      }
    } else if (notification.Status == "reject") {
      switch (notification.Type) {
        case "follow-private":
          return `you rejected ${notification.SenderFullName} follow request`;
        case "group-join":
          return `you refused ${notification.SenderFullName} to join your ${notification.GroupName} group`;
        case "group-invitation":
          return `you rejected ${notification.SenderFullName} request to join ${notification.GroupName} group`;
      }
    } else if (notification.Status == "none" && notification.Type == "group-event") {
      return `${notification.SenderFullName} create event at ${notification.GroupName} group`;
    }
    return "content information not found !!";
  };

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
          Type: notification.Type,
          Status: status,
        })
      };

      let response = await fetch("http://localhost:8080/api/notifications/update/", postRequest);

      if (!response.ok) throw new Error("faild to update notification");

      let data = await response.json();
      console.log(`update notification response: Status: ${data.Status}, Data: ${data.Message}`);

      showNotification({ Content: `Notification ${status}ed successfully`, Status: "success"});

      setNotifications(prev => prev.map(notif => notif.Id === notification.Id ? { ...notif, Status: status } : notif));

    } catch (error) {
      console.error(`Error ${status}ing notification:`, error);

      showNotification({ Content: `Failed to ${status} notification: ${error.message}`, Status: "error" });
    }
  };

  useEffect(() => {
    loadNotifications(notifId);
  }, [notifId]);

  const loadNotifications = async (value) => {
    if (isLoading) return;
    setIsLoading(true);

    try {
      const res = await fetch(`http://localhost:8080/api/notifications?Id=${value}`, { method: "GET", credentials: "include" });

      if (!res.ok) throw new Error("faild to update notification");

      const data = await res.json();

      const existingIds = new Set(notifications.map(notif => notif.Id));

      const newNotifications = data.filter(notif => !existingIds.has(notif.Id));

      setNotifications((prev) => [...prev, ...newNotifications]);

      if (data.length < 10) setHasMore(false);

    } catch (error) {
      console.error("Error fetching notifications:", error);
      setHasMore(false);

    } finally {
      setIsLoading(false);
    }
  };

  const handleScroll = () => {
    if (!containerRef.current || !hasMore || isLoading) return;

    const { scrollTop, scrollHeight, clientHeight } = containerRef.current;

    if (scrollTop + clientHeight >= scrollHeight - 10) {
      const lastNotificationId = notifications?.[notifications.length - 1]?.Id || "0";

      console.log("Triggering next page load with ID:", lastNotificationId);

      setNotifId(lastNotificationId);
    }
  };

  return (
    <div
      ref={containerRef}
      onScroll={handleScroll}
      style={{ maxHeight: "350px", overflowY: "auto", width: "300px" }}
      className="bg-white shadow p-2 rounded"
    >
      {notifications.length === 0 && !isLoading && <p>No notifications</p>}

      {notifications.map((notif) => (
        <div key={notif.Id} className={`notification-card ${notif.Seen ? "seen" : "unseen"}`}>
          <p>{notificationContent(notif)}</p>

          {notif.Status === "later" && notif.Type !== "follow-public" && (
            <div className="action-buttons">
              <button className="accept-btn" onClick={() => handleNotificationAction(notif, "accept")}>✔</button>
              <button className="reject-btn" onClick={() => handleNotificationAction(notif, "reject")}>✘</button>
            </div>
          )}
        </div>
      ))}

      {isLoading && <p className="text-center text-gray-400 text-xs">Loading...</p>}
      {!hasMore && notifications.length > 0 && (
        <p className="text-center text-gray-400 text-xs">No more notifications</p>
      )}
    </div>
  );
}
