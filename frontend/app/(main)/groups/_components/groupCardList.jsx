'use client'
import { useCallback, useEffect, useState, useRef } from "react"
import GroupCard from "./groupCard"
import { useModal } from "../../_context/ModalContext"
import Loader from "../../_components/loader"
import Error from "../../_components/error"
import { redirect, useRouter } from "next/navigation"


const API_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080'

export default function GroupCardList({ filter }) {
    const [data, setData] = useState([])
    const [page, setPage] = useState(0)
    const [isLoading, setIsLoading] = useState(true)
    const [hasMore, setHasMore] = useState(true)
    const [error, setError] = useState(null)
    const observerRef = useRef(null)
    const loadMoreRef = useRef(null)
    const router = useRouter()
    const { getModalData, setModalData } = useModal()

    useEffect(() => {
        const data = getModalData()
        if (data?.type === "groupCard" && filter == "owned") {
            setData((prev) => [data, ...prev])
            setModalData(null)
        }
    }, [setModalData])

    const fetchData = useCallback(
        async (id) => {
            if ((isLoading || !hasMore) && data.length !== 0) return
            try {
                const response = await fetch(`http://localhost:8080/api/groups/?filter=${filter}&offset=${id}`, {
                    credentials: 'include',
                    headers: {
                        'Content-Type': 'application/json'
                    }
                })

                const result = await response.json() || []
                if (!response.ok) {
                    if (result.status === 401) {
                        router.push("/login")
                    }

                    setError(result.error || `Failed to fetch groups`)
                    return
                }
                if (result.length === 0) {
                    setHasMore(false)
                } else {
                    console.log(result)
                    if (result.length < 3) setHasMore(false)
                    setData((prevData) => [...prevData, ...result])
                }
                setIsLoading(false)
            } catch (err) {
                setError(err.message)
                setIsLoading(false)
            }
        },
        [filter]
    )

    useEffect(() => {
        setData([])
        setPage(0)
        setHasMore(true)
        setError(null)
        setIsLoading(true)
        fetchData(0)
    }, [filter])

    useEffect(() => {
        if (!hasMore || isLoading || data.length === 0) return;
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
            let id = data[data.length - 1]?.group_id
            setIsLoading(true)
            fetchData(id)
        }
    }, [page])


    if (error) return <Error error={error} />
    if (isLoading && data.length === 0) return <Loader />

    return (
        <div className="list-container flex flex-wrap gap-4 justify-center items-start overflow-y-auto">
            {data.length === 0
                ? <img
                    className="w-half mx-auto"
                    src="/no-data-animate.svg"
                    alt="No data"
                />
                : data.map((item, index) => (
                    <GroupCard key={item.group_id} type={filter} {...item} />
                ))}
            {isLoading && hasMore && <div className="w-full"> <Loader /></div>}
            {hasMore && !isLoading && (
                <div ref={loadMoreRef} className="w-full" style={{ height: "20px" }}></div>
            )}
        </div>
    )
}