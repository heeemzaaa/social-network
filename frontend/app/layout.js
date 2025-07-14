'use client'
import './global.css'
import { Geist } from 'next/font/google'
import { useRef, useEffect } from 'react'
import TestChat from './(main)/_lib/webSocket'

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
      <body ref={ref}>
		<TestChat />
        {children}
      </body>
    </html>
  );
}
