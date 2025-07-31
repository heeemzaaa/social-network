"use client"
import { HiMiniHome, HiMiniUser, HiMiniUserGroup, HiChatBubbleOvalLeft } from "react-icons/hi2";
import { TbLogout } from "react-icons/tb";
import Button from '@/app/_components/button'
import Link from 'next/link'
import { usePathname } from 'next/navigation'
import "./components.css"
import Logo from "@/app/_components/logo";
import { logout } from "@/app/_actions/user";
import { useEffect, useState } from "react";

export default function Navigation() {
  const currentPath = usePathname()

  const [id, setId] = useState(null)

  useEffect(() => {
    async function GetUserInfo() {
      let res = await fetch("http://localhost:8080/api/loggedin", { credentials: 'include' })
      let data = await res.json()
      setId(data.id)
    }
    GetUserInfo()
  }, []);


  const routes = [
    {
      page: "Home",
      href: "/",
      icon: <HiMiniHome size="24" />
    },
    {
      page: "Profile",
      href: `/profile/${id}`,
      icon: <HiMiniUser size="24" />
    },
    {
      page: "Groups",
      href: "/groups",
      icon: <HiMiniUserGroup size="24" />
    },
    {
      page: "Chat",
      href: "/chat",
      icon: <HiChatBubbleOvalLeft size="24" />
    }
  ]

  return (
    <aside>
      <Logo />
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
      <form action={logout}>
        <Button className={"logout-btn"} style={{ gap: "6px" }} variant='btn-danger' type="submit">
          <span>
            Log-out
          </span>
          <TbLogout size={20} />
        </Button>
      </form>
    </aside>
  )
}
