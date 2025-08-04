"use client";
import React, { createContext, useState, useContext } from "react";

const NotificationContext = createContext();

export function useNotification() {
  return useContext(NotificationContext);
}

// Toast Component - simple w clean
function ToastNotification({ notif, onClose }) {
  const { Status, Content } = notif;
  
  return (
    <div className={`toast-popup ${Status}`}>
      <div className="toast-content">
        <div className="toast-message">
          {Content}
        </div>
        <button className="toast-close" onClick={onClose}>×</button>
      </div>
    </div>
  );
}

export function NotificationProvider({ children }) {
  const [currentNotification, setCurrentNotification] = useState(null);
  const [timeoutId, setTimeoutId] = useState(null);

  const showNotification = (notifObj) => !currentNotification ? displayNotification(notifObj) : null;
  //   if (currentNotification) return

  //   displayNotification(notifObj);
  // };

  const displayNotification = (notifObj) => {
    setCurrentNotification(notifObj);

    // Clear l timeout li 9dim ila kan
    if (timeoutId) {
      clearTimeout(timeoutId);
    }

    // Dir timeout jdid
    const newTimeoutId = setTimeout(() => {
      hideNotification();
    }, 4000);
    
    setTimeoutId(newTimeoutId);
  };

  const hideNotification = () => {
    // Clear timeout
    if (timeoutId) {
      clearTimeout(timeoutId);
      setTimeoutId(null);
    }

    // Add slide-out animation
    const toastElement = document.querySelector('.toast-popup');
    if (toastElement) {
      toastElement.classList.add('slide-out');
      setTimeout(() => {
        setCurrentNotification(null);
        showNextInQueue();
      }, 300);
    } else {
      setCurrentNotification(null);
      showNextInQueue();
    }
  };

  const showNextInQueue = () => {
    // Ma kaynch queue system, just clean up
    setCurrentNotification(null);
  };

  const closeNotification = () => {
    hideNotification();
  };

  return (
    <NotificationContext.Provider value={{ showNotification }}>
      {children}
      
      {/* Render current notification only */}
      {currentNotification && (
        <div 
          style={{ 
            position: 'fixed',
            top: '20px',
            right: '20px',
            zIndex: 9999
          }}
        >
          <ToastNotification 
            notif={currentNotification} 
            onClose={closeNotification}
          />
        </div>
      )}
    </NotificationContext.Provider>
  );
}

// // ✅ Generic Success Messages
// // "Your action was successful."

// // "Success! Everything went as expected."

// // "Done! Your request has been processed."

// // 💬 More Context-Specific (Friendly Style)
// // "🎉 Invitation sent successfully!"

// // "✅ You've successfully joined the group."

// // "Your transaction was completed successfully."

// // "👍 Changes saved."

// // "The invitation has been cancelled."

// // "Your message was delivered."

// // "Profile updated successfully!"