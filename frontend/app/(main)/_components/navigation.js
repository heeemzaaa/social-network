"use client"
import { HiMiniHome,HiMiniUser,HiMiniUserGroup,HiChatBubbleOvalLeft } from "react-icons/hi2";
import Button from '@/app/_components/button'
import Link from 'next/link'
import { usePathname } from 'next/navigation'
import "./components.css"
import Logo from "@/app/_components/logo";

export default function Navigation() {
  const currentPath = usePathname()
  const routes = [
    {
      page: "Home",
      href: "/",
      icon: <HiMiniHome  size="24"/>
    },
    {
      page: "Profile",
      href: "/profile/1",
      icon: <HiMiniUser size="24"/>
    },
    {
      page: "Groups",
      href: "/groups",
      icon: <HiMiniUserGroup size="24"/>
    },
    {
      page: "Chat",
      href: "/chat",
      icon: <HiChatBubbleOvalLeft size="24"/>
    }
  ]

  return (
    <aside>
      <Logo/>
      <nav>
        {
          routes.map(route => <Link className={`link ${route.href === currentPath ? "link-active" : ""}`} key={route.page} href={route.href}>
            {route.icon}
            <span>
              {route.page}
            </span>
          </Link>)
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
