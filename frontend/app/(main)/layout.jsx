"use client"

import Header from './_components/header'
import Navigation from './_components/navigation'
import { ModalProvider } from './_context/ModalContext'
import { NotificationProvider } from './_context/NotificationContext'; // ✅ import it
import UserProvider from './_lib/webSocket';

export default function MainLayout({ children }) {
  return (
    <UserProvider>
      <ModalProvider>
        <NotificationProvider> {/* ✅ wrap everything */}
          <Header />
          <Navigation />
          {children}
        </NotificationProvider>
      </ModalProvider>
    </UserProvider>
  )
}
