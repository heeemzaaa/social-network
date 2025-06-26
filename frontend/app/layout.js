import './global.css'
import { Geist } from 'next/font/google'
 
const geist = Geist({
  subsets: ['latin'],
})

export default function RootLayout({ children }) {
  console.log(children)

  return (
    <html lang="en" className={geist.className}>
      <body>
        {children}
      </body>
    </html>
  );
}
