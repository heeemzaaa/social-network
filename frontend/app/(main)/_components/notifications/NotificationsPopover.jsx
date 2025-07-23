"use client";
import { useEffect, useState, useRef } from "react";

export default function NotificationsPopover() {
  const [notifications, setNotifications] = useState([]);
  const [page, setPage] = useState(0);
  const [hasMore, setHasMore] = useState(true);
  const containerRef = useRef(); 
  

  // Handle accept/reject notification
  const handleNotificationAction = async (notification, status) => {
    
    console.log(`${status} notification:`, notification);
    console.log(`Notification ID:`, notification.Id);
    
    if (!notification.Id) {
      console.error('No notification ID found. Available fields:', Object.keys(notification));
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
          Notif_Id: notification.Id,
          Status: status, // "accept" or "reject"
        })
      };

      let response = await fetch("http://localhost:8080/api/notifications/update/", postRequest);
      let data = await response.json();
      console.log(`Notification ${status} response:`, data);

      // Update the notification in the state
      setNotifications(prev => prev.map(notif => notif.Id === notification.Id ? { ...notif, Status: status } : notif));

    } catch (error) {
      console.error(`Error ${status}ing notification:`, error);
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

      console.log(data);

      if (data.length === 0) {
        setHasMore(false);
      } else {
        // Filter out duplicates based on ID
        setNotifications((prev) => {

          const existingIds = new Set(prev.map(notif => notif.Id));
          
          const newNotifications = data.filter(notif => {
            const notifId = notif.Id
            return !existingIds.has(notifId);
          });
          
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

    if (scrollTop + clientHeight >= scrollHeight - 10) setPage((prev) => prev + 10);
  };

  return (
    
    <div ref={containerRef} onScroll={handleScroll} style={{ maxHeight: "350px", overflowY: "auto", width: "300px" }} className="bg-white shadow p-2 rounded">

      {notifications.length === 0 && <p>No notifications</p>}

      {notifications.map((notif) => (

        <div key={notif.Id} className={`notification-card ${notif.Type} ${notif.Status} ${notif.Seen ? "seen" : "unseen"}`}>

          <p>{notif.Content}</p>

          {notif.Status === "later" && notif.Type !== "follow-public" && (
            <div className="action-buttons">
              <button className="accept-btn" onClick={() => handleNotificationAction(notif, "accept")}>
                ✔
              </button>
              <button className="reject-btn" onClick={() => handleNotificationAction(notif, "reject")}>
                ✘
              </button>
            </div>
          )}
          {notif.Status === "accept" && (
            <div className="green-dote">
            </div>
          )}
          {notif.Status === "reject" && (
            <div className="red-dote">
            </div>
          )}

        </div>
      ))}

      {hasMore && <p className="text-center text-gray-400 text-xs">Loading more...</p>}
      {!hasMore && notifications.length > 0 && <p className="text-center text-gray-400 text-xs">No more notifications</p>}
    </div>
  );
}