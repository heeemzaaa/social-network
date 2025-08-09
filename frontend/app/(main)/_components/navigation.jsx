"use client"
import "./components.css"
import Link from 'next/link'
import Logo from "@/app/_components/logo";
import Button from '@/app/_components/button'
import { logout } from "@/app/_actions/user";
import { TbLogout } from "react-icons/tb";
import { usePathname } from 'next/navigation'
import {
  useEffect,
  useState
} from "react";
import {
  HiMiniHome,
  HiMiniUser,
  HiMiniUserGroup,
  HiChatBubbleOvalLeft
} from "react-icons/hi2";

export default function Navigation() {
  const currentPath = usePathname()

  const [id, setId] = useState(null)

  useEffect(() => {
    async function GetUserInfo() {
      console.log("profile heeree the endpoint", `${process.env.NEXT_PUBLIC_API_URL}/api/loggedin`);
      let res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/loggedin`, { credentials: 'include' })
      let data = await res.json()
      console.log("data feteched", data);
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
