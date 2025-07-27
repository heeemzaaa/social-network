// _context/NotificationContext.jsx
"use client"
import React, { createContext, useState, useContext, useCallback } from "react";
import Notification from "../_components/notifications/Notification"; // adjust if needed

const NotificationContext = createContext();

export function useNotification() {
  return useContext(NotificationContext);
}

export function NotificationProvider({ children }) {
  const [notif, setNotif] = useState(null);
  const [confirmationCallback, setConfirmationCallback] = useState(null);

  const showNotification = (notifObj) => {
    setNotif(notifObj);
    setTimeout(() => setNotif(null), 4000); // Auto-dismiss
  };

  const confirmNotification = useCallback((notifObj) => {
    return new Promise((resolve) => {
      setNotif({ ...notifObj, isConfirm: true });
      setConfirmationCallback(() => (result) => {
        resolve(result);
        setNotif(null);
        setConfirmationCallback(null);
      });
    });
  }, []);

  const handleConfirm = (response) => {
    if (confirmationCallback) {
      confirmationCallback(response);
    }
  };

  return (
    <NotificationContext.Provider value={{ showNotification, confirmNotification }}>
      {children}
      {notif && (
        <div className="notification-wrapper">
          <Notification notif={notif} isPopup={true} onConfirm={handleConfirm} />
        </div>
      )}
    </NotificationContext.Provider>
  );
}
