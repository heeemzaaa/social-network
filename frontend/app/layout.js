'use client'
import './global.css'
import { Geist } from 'next/font/google'
import { useRef, useEffect } from 'react'

const geist = Geist({
  subsets: ['latin'],
})



export default function RootLayout({ children }) {
  const ref = useRef(null);
  
  useEffect(() => {
    const handleOutSideClick = (event) => {
      if (!ref.current?.contains(event.target)) {
        alert("Outside Clicked.");
        console.log("Outside Clicked. ");
      }
    };
    window.addEventListener("mousedown", handleOutSideClick);
    return () => {
      window.removeEventListener("mousedown", handleOutSideClick);
    };
  }, [ref]);
  
  return (
	   <html lang="en" className={geist.className}>
      <head>
        <title>EmiTalk</title>
        <link rel="icon" href="/logo.svg" />
      </head>
      <body ref={ref}>
        {children}
      </body>
    </html>
  );
}
