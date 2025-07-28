"use client"

import Header from './_components/header'
import Navigation from './_components/navigation'
import { ModalProvider } from './_context/ModalContext'
import { NotificationProvider } from './_context/NotificationContext'; // ✅ import it

export default function MainLayout({ children }) {
  return (
    <ModalProvider>
      <NotificationProvider> {/* ✅ wrap everything */}
        <Header />
        <Navigation />
        {children}
      </NotificationProvider>
    </ModalProvider>
  )
}
