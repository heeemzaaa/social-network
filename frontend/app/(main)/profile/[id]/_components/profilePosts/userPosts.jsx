'use client'

import React, { useEffect, useState } from 'react'
import PostCard from '@/app/(main)/_components/posts/postCard'
import { useModal } from '@/app/(main)/_context/ModalContext'

export default function UserPosts({ id, access, changed }) {
    const [posts, setPosts] = useState([])
    const {setModalData, getModalData} = useModal()

     useEffect(() => {
        let postData = getModalData()
        if (postData?.type !== 'post') return;

        setPosts((prev) => {
            if (!prev) {
                return [postData]
            } else {
                return [postData, ...prev]
            }
        })
        setModalData(null)
    }, [setModalData])
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
    }, [id, changed])

    if (access === false) {
        return (
            <section className='posts_container w-full h-full flex-col justify-center align-center'>
                <img src="/forbidden-posts.svg" style={{ height: '100%' }} />
                <p className='text-2xl'>You must follow to see the posts</p>
            </section>
        )
    }

    return (
        <section className='posts_container scrollable-section w-full h-full'>
            {posts.length === 0 ? (
                <img src="/no-posts.svg" className='w-full h-full' />
            ) : (
                posts.map((post) => {
                    return <PostCard {...post} key={post.id} />
                })
            )}
        </section>
    )
}
