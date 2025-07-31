import { useCallback, useEffect, useState, useRef } from "react"
import Button from "@/app/_components/button"
import PostCard from "../../_components/posts/postCard"
import { useModal } from "../../_context/ModalContext"

export default function GroupPostCardList({ groupId, setIsAccessible, isAccessible }) {
    const [data, setData] = useState([])
    const [page, setPage] = useState(0)
    const [isLoading, setIsLoading] = useState(false)
    const [hasMore, setHasMore] = useState(true)
    const [error, setError] = useState(null)
    const abortControllerRef = useRef(null)
    const observerRef = useRef(null)
    const loadMoreRef = useRef(null)

    const { getModalData, setModalData } = useModal()

    useEffect(() => {
        const data = getModalData()
        if (data?.type === "groupPost") {
            setData((prev) => [data, ...prev])
            setModalData(null)
        }
    }, [getModalData])

    const getUrl = useCallback(
        (page) => {
            const params = new URLSearchParams({
                offset: page * 20,
            })
            return `http://localhost:8080/api/groups/${groupId}/posts/?${params.toString()}`
        },
        [groupId]
    )

    const fetchData = useCallback(
        async (currentPage) => {
            if (isLoading || !hasMore) return
            setIsLoading(true)
            abortControllerRef.current = new AbortController()
            const signal = abortControllerRef.current.signal
            try {
                const url = getUrl(currentPage)
                const response = await fetch(url, { credentials: "include", signal })
                const result = await response.json()
                if (!response.ok) {
                    if (response.status == 403) setIsAccessible(response)
                    throw new Error(`HTTP error! status: ${response.status}`)
                }
                if (result.length === 0) {
                    setHasMore(false)
                } else {
                    if (result.length < 20) setHasMore(false)
                    setData((prevData) => [...prevData, ...result])
                }
            } catch (err) {
                setError(err.message)
            } finally {
                setIsLoading(false)
            }
        },

        [getUrl]
    )

    useEffect(() => {
        setData([])
        setPage(0)
        setHasMore(true)
        setError(null)
        fetchData(0)
    }, [groupId])

    useEffect(() => {
        if (!hasMore || isLoading) return

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

    useEffect(() => {
        if (page > 0) {
            fetchData(page)
        }
    }, [page])

    if (isAccessible?.status == 403) {
        return (
            <section className='posts_container w-full h-full flex-col justify-center align-center'>
                <img src="/forbidden-posts.svg" style={{ height: '100%' }} />
                <p className='text-xl font-semibold'>You must become a member to see the posts</p>
            </section>
        )
    }

    return (
        <div className="list-container flex align-start flex-wrap gap-4 justify-center overflow-y-auto">
            {data.map((item) => (
                <PostCard {...item} key={item.id} groupId={groupId} />
            ))}
            {data.length === 0 && <img className="w-half mx-auto" src="/no-data-animate.svg" alt="No data" />}
            {isLoading && <p className="text-center w-full">Loading...</p>}
            {hasMore && !isLoading && (
                <div ref={loadMoreRef} className="w-full" style={{ height: "20px" }}></div>
            )}
        </div>
    )
}