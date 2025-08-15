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

  const handleNotificationAction = async (notification, status) => {
    if (!notification.id) {
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
          NotifId: notification.id,
          Type: notification.type_notification,
          Status: status,
        })
      };

      let response = await fetch("http://localhost:8080/api/notifications/update/", postRequest);
      let data = await response.json();
      if (!data.ok) {

        if (data.error === "Notification not found" || data.error === "Invitation not found" || data.error === "Already a member!" ) {
          console.warn("Notification not found, remove from the list")
          setNotifications(prev => prev.filter(notif => notif.id !== notification.id))
          showNotification({ Content: "Notification not found, remove from the list", Status: "warn" })
          return
        }
        // showNotification({Content:"failed to update notification", Status:"error"})
      }

      if (data.status === true ) {
        setNotifications(prev => prev.map(notif => notif.id === notification.id ? { ...notif, status: status } : notif ))
        showNotification({ Content: `Notification ${status}ed successfully`, Status: "success"});
      }

    } catch (error) {
      console.error(`Error ${status}ing notification:`, error);
      showNotification({ Content: `Failed to ${status} notification: ${error.error}`, Status: "error" });
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

      if (!res.ok) {
        showNotification({Content:"failed to update notification", Status:"error"})
      }

      const data = await res.json();

      const existingIds = new Set(notifications.map(notif => notif.id));

      const newNotifications = data?.Notifications.filter(notif => !existingIds.has(notif.id));

      setNotifications((prev) => [...prev, ...newNotifications]);

      if (data?.Notifications.length < 10) setHasMore(false);

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
      const lastNotificationId = notifications?.[notifications.length - 1]?.id || "0";

      setNotifId(lastNotificationId);
    }
  };

  return (
    <div
      ref={containerRef}
      onScroll={handleScroll}
      style={{ maxHeight: "350px", overflowY: "auto", width: "300px" }}
      className="notifContainer bg-white shadow p-2 rounded"
    >
      {notifications.length === 0 && !isLoading && <img src="/no-notifications.svg" style={{width: '100%', height: '100%'}}/>}

      {notifications.map((notif) => (
        <div key={notif.id} className={`notification-card ${notif.seen ? "seen" : "unseen"}`}>
          <p>{notif.content || "content information not found !!"}</p>

          {notif.status === "later" && (
            <div className="action-buttons">
              <button className="accept-btn" onClick={() => handleNotificationAction(notif, "accept")}>✔</button>
              <button className="reject-btn" onClick={() => handleNotificationAction(notif, "reject")}>✘</button>
            </div>
          )}
        </div>
      ))}

      {isLoading && <p className="text-center text-gray-400 text-xs">Loading...</p>}
    </div>
  );
}
