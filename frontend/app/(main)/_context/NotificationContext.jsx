// _context/NotificationContext.js
"use client"

import React, { createContext, useContext, useState } from "react";

const NotificationContext = createContext();

export function useNotification() {
  return useContext(NotificationContext);
}

export function NotificationProvider({ children }) {
  const [notification, setNotification] = useState(null);

  const showNotification = ({ Content, Status }) => {
    setNotification({ Content, Status });
    setTimeout(() => setNotification(null), 4000);
  };

  const onClose = () => setNotification(null);

  return (
    <NotificationContext.Provider value={{ showNotification }}>
      {children}

      {notification && (
        <div className={`toast-popup ${notification.Status}`}>
          <div className="toast-content">
            <div className="toast-message">
              {notification.Content}
            </div>
            <button className="toast-close" onClick={onClose}>x</button>
          </div>
        </div>
      )}
    </NotificationContext.Provider>
  );
}
