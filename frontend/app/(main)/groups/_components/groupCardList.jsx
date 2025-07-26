import { useCallback, useEffect, useState, useRef } from "react"
import GroupCard from "./groupCard"
import { useModal } from "../../_context/ModalContext"

export default function GroupCardList({ filter }) {
    const [data, setData] = useState([])
    const [page, setPage] = useState(0)
    const [isLoading, setIsLoading] = useState(false)
    const [hasMore, setHasMore] = useState(true)
    const [error, setError] = useState(null)
    const abortControllerRef = useRef(null)
    const observerRef = useRef(null)
    const loadMoreRef = useRef(null)

    const { getModalData } = useModal()

    useEffect(() => {
        const data = getModalData()
        if (data?.type === "groupCard" && filter === "owned") {
            setData((prev) => [data, ...prev])
        }
    }, [getModalData])

    const fetchData = useCallback(
        async (id) => {
            console.log("fetch created groups: ", id)
            if (isLoading || !hasMore) return
            setIsLoading(true)
            abortControllerRef.current = new AbortController()
            const signal = abortControllerRef.current.signal
            try {
                const response = await fetch(`http://localhost:8080/api/groups?filter=${filter}&offset=${id}`,
                    { credentials: "include", signal })
                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`)
                }
                const result = await response.json()
                if (result.length === 0) {
                    console.log("!!! you reached the end")
                    setHasMore(false)
                } else {
                    if (result.length < 6) setHasMore(false)
                    console.log("!!! you reached the end")
                    setData((prevData) => [...prevData, ...result])
                }
            } catch (err) {
                if (err.name === "AbortError") return
                setError(err.message)
            } finally {
                setIsLoading(false)
            }
        },
        [filter]
    )

    useEffect(() => {
        console.log("initial fetch", [console.log(data)])
        setData([])
        setPage(0)
        setHasMore(true)
        setError(null)
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
            console.log("previous fetched data: ", data)
            console.log("page: ", page)
            let id = data[data.length - 1]?.group_id 
            fetchData(id)
        }
    }, [page])

    if (error) return <p className="text-danger text-center">Error: {error}</p>

    return (
        <div className="list-container flex flex-wrap gap-4 justify-center items-start overflow-y-auto">
            {data.map((item, index) => (
                <GroupCard key={item.group_id} type={filter} {...item} />
            ))}
            {data.length === 0 && (
                <img
                    className="w-half mx-auto"
                    src="/no-data-animate.svg"
                    alt="No data"
                />
            )}
            {isLoading && <p className="text-center w-full">Loading...</p>}
            {hasMore && !isLoading && (
                <div ref={loadMoreRef} className="w-full" style={{ height: "20px" }}></div>
            )}
        </div>
    )
}