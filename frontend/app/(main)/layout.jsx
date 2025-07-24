"use client"

import Header from './_components/header'
import Navigation from './_components/navigation'
import { ModalProvider } from './_context/ModalContext'

export default function MainLayout({ children }) {
  return (
    <>
      <ModalProvider>
        <Header />
        <Navigation />
        {children}
      </ModalProvider>
    </>
  )
}
