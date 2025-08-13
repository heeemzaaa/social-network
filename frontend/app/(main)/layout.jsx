"use client"

import Header from './_components/header'
import Navigation from './_components/navigation'
import UserProvider from './_lib/webSocket';
import { ModalProvider } from './_context/ModalContext'
import { NotificationProvider } from './_context/NotificationContext';

export default function MainLayout({ children }) {
  return (
      <NotificationProvider>
    <UserProvider>
        <ModalProvider>
          <Header />
          <Navigation />
          {children}
        </ModalProvider>
    </UserProvider>
      </NotificationProvider>
  )
}
