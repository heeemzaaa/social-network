'use client'

import React, { useCallback, useEffect, useRef, useState } from 'react'
import PostCard from '@/app/(main)/_components/posts/postCard'
import { useModal } from '@/app/(main)/_context/ModalContext'
import Error from '@/app/(main)/_components/error'
import Loader from '@/app/(main)/_components/loader'
import { useRouter } from 'next/navigation'

export default function UserPosts({ profileId, access, changed }) {
    const [posts, setPosts] = useState([])
    const [page, setPage] = useState(0)
    const [isLoading, setIsLoading] = useState(true)
    const [hasMore, setHasMore] = useState(true)
    const [error, setError] = useState(null)
    const observerRef = useRef(null)
    const loadMoreRef = useRef(null)
    const { setModalData, getModalData } = useModal()
    const router = useRouter()


    // add the new created post to posts state
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

    // post fetch function
    const fetchData = useCallback(
        async (id) => {
            if ((isLoading || !hasMore) && posts.length !== 0) return
            try {
                const response = await fetch(`http://localhost:8080/api/profile/${profileId}/data/posts${id != 0 ? "?last=" + id : ""}`, {
                    credentials: 'include',
                    headers: {
                        'Content-Type': 'application/json'
                    }
                })
                const result = await response.json() || []
                if (!response.ok) {
                    if (result.status === 401) {
                        router.push("/login")
                        return
                    }
                    if (result.status === 400 || result.status === 404 || result.status === 500) {
                        setError(result.error)
                        return
                    }
                }
                if (result.length === 0) {
                    setHasMore(false)
                } else {
                    if (result.length < 10) setHasMore(false)
                    setPosts((prevData) => [...prevData, ...result])
                }
                setIsLoading(false)
            } catch (err) {
                setError(err.message)
                setIsLoading(false)
            }
        },
        [profileId]
    )

    // first fetch
    useEffect(() => {
        setPosts([])
        setPage(0)
        setHasMore(true)
        setError(null)
        setIsLoading(true)
        fetchData(0)
    }, [profileId, changed])


    // increment page count when the intersection observer is triggered
    useEffect(() => {
        if (!hasMore || isLoading || posts.length === 0) return;
        observerRef.current = new IntersectionObserver(
            (entries) => {
                if (entries[0].isIntersecting) {
                    setPage((prevPage) => prevPage + 1)
                }
            },
            { threshold: 0.1 }
        )

        if (loadMoreRef.current) {
            observerRef.current.observe(loadMoreRef.current)
        }

        return () => {
            if (observerRef.current) {
                observerRef.current.disconnect()
            }
        }
    }, [hasMore, isLoading])

    // fetch new posts on page change
    useEffect(() => {
        if (page > 0) {
            let id = posts[posts.length - 1]?.id
            setIsLoading(true)
            fetchData(id)
        }
    }, [page])


    if (error) return <Error error={error} />
    if (isLoading && posts.length === 0) return <Loader />

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
            {isLoading && hasMore && <div className="w-full"> <Loader /></div>}
            {hasMore && !isLoading && (
                <div ref={loadMoreRef} className="w-full" style={{ height: "20px" }}></div>
            )}
        </section>
    )
}
