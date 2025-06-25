import React from 'react'
import Header from './_components/header'

export default function MainLayout({children}) {
  return (
    <>
      <Header />
      {children}
    </>
  )
}
