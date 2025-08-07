import React, { useEffect, useState } from 'react'
import ConnectionCard from './followerCard'


export default function UsersContainer({ type, userID }) {
  const [data, setData] = useState([])

  useEffect(() => {
    async function handleGetConnections() {
      try {
        const res = await fetch(`http://localhost:8080/api/profile/${userID}/connections/${type}`, {
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
        return <ConnectionCard key={user.id} {...user} />
      })}
    </div>
  )
}
