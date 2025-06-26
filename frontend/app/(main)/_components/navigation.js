"use client"
import Button from '@/app/_components/button'
import Link from 'next/link'
import { usePathname } from 'next/navigation'
import "./components.css"

export default function Navigation() {
  const currentPath = usePathname()
  const routes = [
    {
      page: "Home",
      href: "/"
    },
    {
      page: "Profile",
      href: "/profile/1"
    },
    {
      page: "Groups",
      href: "/groups"
    },
    {
      page: "Chat",
      href: "/chat"
    }
  ]

  return (
    <aside>
      <Link href={"/"}>
        <h2>Social Network</h2>
      </Link>
      <nav>
        {
          routes.map(route => <Link className={`link ${route.href === currentPath ? "link-active" : ""}`} key={route.page} href={route.href}>{route.page}</Link>)
        }
      </nav>
      <Button variant='btn-danger'>
        <span>
          Log-out
        </span>
      </Button>
    </aside>
  )
}
