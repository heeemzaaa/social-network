'use client'
import './global.css'
import { Geist } from 'next/font/google'

const geist = Geist({
  subsets: ['latin'],
})



export default function RootLayout({ children }) {
  return (
    <html lang="en" className={geist.className}>
      <head>
        <title>EmiTalk</title>
        <link rel="icon" href="/logo.svg" />
      </head>
      <body>
        {children}
      </body>
    </html>
  );
}
