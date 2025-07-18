"use client";
import { useEffect, useState, useRef } from "react";

export default function NotificationsPopover() {
  const [notifications, setNotifications] = useState([]);
  const [page, setPage] = useState(0);
  const [hasMore, setHasMore] = useState(true);
  const containerRef = useRef();

  // Load notifications
  useEffect(() => {
    loadNotifications(page);
  }, [page]);

  const loadNotifications = async (data_length) => {
    try {
      // Call your Go backend, here you can pass offset or page
      const res = await fetch(`http://localhost:8080/api/notifications?Count=${data_length}`, {
        method: "GET",
        credentials: "include"
      });
      const data = await res.json();

      if (data.length === 0) {
        setHasMore(false);
      } else {
        setNotifications((prev) => [...prev, ...data]); 
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
      style={{ maxHeight: "300px", overflowY: "auto", width: "250px" }}
      className="bg-white shadow p-2 rounded"
    >
      {notifications.length === 0 && <p>No notifications</p>}

      {notifications.map((notif) => (
        <div
            className={`notification-card ${notif.Type} ${notif.Status}`}
            >
            <p>{notif.Content}</p>

            {/* Show action buttons if status is later & type is not follow-public */}
            {notif.Status === "later" && notif.Type !== "follow-public" && (
                <div className="action-buttons">
                <button className="accept-btn">Accept</button>
                <button className="reject-btn">Reject</button>
                </div>
            )}
        </div>
      ))}

      {hasMore && <p className="text-center text-gray-400 text-xs">Loading more...</p>}
      {!hasMore && notifications.length > 0 && <p className="text-center text-gray-400 text-xs">No more notifications</p>}
    </div>
  );
}
