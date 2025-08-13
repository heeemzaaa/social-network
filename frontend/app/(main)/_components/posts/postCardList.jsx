"use client"
import PostCard from "./postCard"
import { useModal } from "../../_context/ModalContext"
import { useEffect, useState } from "react"
import Loader from "../loader"
import Error from "../error"
import { useRouter } from "next/navigation"

const API_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080'

export default function PostCardList() {
    const [posts, setPosts] = useState([])
    const [loading, setLoading] = useState(true)
    const [error, setError] = useState(null)
    const { getModalData, setModalData } = useModal()
    const router = useRouter()

    useEffect(() => {
        let postData = getModalData()
        if (postData?.type !== 'post') return

        setPosts((prev) => {
            if (!prev) {
                return [postData]
            } else {
                return [postData, ...prev]
            }
        })
        setModalData(null)
    }, [getModalData, setModalData])

    useEffect(() => {
        async function fetchPosts() {
            try {
                setLoading(true)
                setError(null)
                const resp = await fetch(`${API_URL}/api/posts`, {
                    method: "GET",
                    credentials: "include",
                })
                const data = await resp.json()
                if (!resp.ok) {
                    if (resp.status === 401) {
                        router.push("/login")
                        return
                    } 
                    
                    setError(data?.error || "Failed to load posts")
                } else {
                    setPosts(data)
                }
                setLoading(false)
            } catch (error) {
                setError("Failed to load posts. Please try again later.")
                console.error("Error fetching posts:", error)
                setLoading(false)
            }
        }

        fetchPosts()
    }, [])

    if (loading) return <Loader />
    if (error) return <Error error={error} />

    return (
        <div className="list-container" style={{ overflowY: "auto" }}>
            {posts?.length > 0 ? (
                posts?.map((post) => (
                    <PostCard key={post.id} {...post} post={post}/>
                ))
            ) : (
                <div style={{ width: "100%", maxWidth: "500px", margin: "auto", textAlign: "center" }}>
                    <img src="/noFeed.svg" alt="No posts available" />
                    <p className="font-semibold">No posts available for you!</p>
                </div>
            )}
        </div>
    )
}