'use client'
import React, { useEffect, useState } from 'react'

export default function UserPosts({ id, access }) {
    const [posts, setPosts] = useState([])

    useEffect(() => {
        async function getPosts() {
            try {
                const res = await fetch(`http://localhost:8080/api/profile/${id}/data/posts`, { credentials: 'include' })
                if (res.ok) {
                    const data = await res.json()
                    if (data) setPosts(data)
                }
            } catch (err) {
                console.error("Error fetching posts:", err)
            }
        }

        getPosts()
    }, [id])

    if (access === false) {
        return (
            <section className='post_container w-full h-full flex-col justify-center align-center'>
                <img src="/forbidden-posts.svg" style={{ height: '100%' }} />
                <p className='text-2xl'>You must follow to see the posts</p>
            </section>
        )
    }

    return (
        <section className='post_container w-full h-full flex-col justify-center align-center'>
            {posts.length === 0 ? (
                <img src="/no-posts.svg" style={{ height: '90%' }} />
            ) : (
                posts
            )}
        </section>
    )
}
