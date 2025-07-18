'use client'
import React, { useEffect, useState } from 'react'

export default function UserPosts({ id }) {
    const [posts, setPosts] = useState([])
    const [access, setAccess] = useState(false)

    useEffect(() => {
        async function getPosts() {
            try {
                let res = await fetch(`http://localhost:8080/api/profile/${id}/data/posts`, { credentials: 'include' })
                if (res.ok) {
                    let data = await res.json()
                    if (data === false) {
                        return
                    }
                    setAccess(true)
                    setPosts(data)
                }
            } catch (err) {
                console.error("Error: ", err)
            }
        }
        getPosts()
    }, [id])


    return (
        <section className='post_container w-full h-full flex-col justify-center align-center'>
            {access == false ? 
            <>
            <img src="/forbidden-posts.svg" style={{height: '100%'}} />
            <p className='text-2xl'>You must follow to see the posts</p>
            </>
             :
              posts.length === 0 ?
                <img src="/no-posts.svg" style={{height: '90%'}}/> : posts}
        </section>
    )
}