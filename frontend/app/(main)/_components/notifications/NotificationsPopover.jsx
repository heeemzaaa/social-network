"use client";
import { useEffect, useState, useRef } from "react";

export default function NotificationsPopover() {
  const [notifications, setNotifications] = useState([]);
  const [page, setPage] = useState(0);
  const [hasMore, setHasMore] = useState(true);
  const containerRef = useRef();

  // Test function
  const test = () => {
    console.log("Test function called!");
    
    // إذا كان container فارغ، زيد notification جديدة
    if (notifications.length === 0) {
      const newNotification = {
        Content: "This is a test notification",
        Type: "test",
        Status: "new",
        id: Date.now() // temporary id
      };
      setNotifications([newNotification]);
      console.log("Added new notification because container was empty");
    } else {
      console.log("Container is not empty, notifications count:", notifications.length);
    }
  };

  // Load notifications
  useEffect(() => {
    loadNotifications(page);
  }, [page]);

  // Add event listener for button clicks
  useEffect(() => {
    const handleButtonClick = (event) => {
      // Check if clicked element is a button
      if (event.target.tagName === 'BUTTON') {
        test();
      }
    };

    // Add event listener to container
    if (containerRef.current) {
      containerRef.current.addEventListener('click', handleButtonClick);
    }

    // Cleanup event listener
    return () => {
      if (containerRef.current) {
        containerRef.current.removeEventListener('click', handleButtonClick);
      }
    };
  }, [notifications.length]); // Re-run when notifications change

  const loadNotifications = async (data_length) => {
    try {
      // Call your Go backend, here you can pass offset or page
      const res = await fetch(`http://localhost:8080/api/notifications?Count=${data_length}`, {
        method: "GET",
        credentials: "include"
      });
      const data = await res.json();

      console.log(data)

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

      {notifications.map((notif, index) => (
        <div
            key={notif.id || index}
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