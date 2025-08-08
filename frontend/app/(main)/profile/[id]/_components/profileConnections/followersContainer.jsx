import FollowerCard from './followerCard'
import { useEffect, useState } from 'react'


export default function UsersContainer({ type, userID }) {
  const [data, setData] = useState([])

  useEffect(() => {
    async function handleGetConnections() {
      try {
        const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/profile/${userID}/connections/${type}`, {
          credentials: "include",
        })

        if (res.ok) {
          const result = await res.json()
          setData(result)
        }

      } catch (err) {
        console.error("Failed to fetch followers", err)
      }
    }

    handleGetConnections()
  }, [type])

  return (
    <div className='follow_container p2 gap-1'>
      {data.map((user) => {
        return <FollowerCard key={user.id} {...user} />
      })}
    </div>
  )
}
