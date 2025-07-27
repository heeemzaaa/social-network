"use client"

import Header from './_components/header'
import Navigation from './_components/navigation'
import { ModalProvider } from './_context/ModalContext'
import UserProvider from './_lib/webSocket'



export default function MainLayout({ children }) {
  return (
    <UserProvider>
      <ModalProvider>
        <Header />
        <Navigation />
        {children}
      </ModalProvider>
    </UserProvider>
  )
}
