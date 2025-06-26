import Button from '@/app/_components/button'
import Link from 'next/link'

export default function Navigation() {
  return (
    <aside>
      <h2>Social Network</h2>

      <nav className={""}>
        <Link href="/"><span>Home</span></Link>
        <Link href="/profile"><span>Profile</span></Link>
        <Link href="/groups"><span>Groups</span></Link>
        <Link href="/chat"><span>Chat</span></Link>
      </nav>
      <Button variant='btn-primary'>
        <span>
          Log-out
        </span>
      </Button>

    </aside>
  )
}
