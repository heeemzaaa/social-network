import { useCallback, useEffect, useState, useRef } from "react";
import GroupEventCard from "./groupEventCard";
import { useModal } from "../../_context/ModalContext";

export default function GroupEventCardList({ groupId, setIsAccessible, isAccessible }) {
    const [data, setData] = useState([]);
    const [page, setPage] = useState(0);
    const [isLoading, setIsLoading] = useState(false);
    const [hasMore, setHasMore] = useState(true);
    const [error, setError] = useState(null);
    const abortControllerRef = useRef(null);
    const observerRef = useRef(null);
    const loadMoreRef = useRef(null);

    const { getModalData, setModalData } = useModal();

    // Handle modal data for new group events
    useEffect(() => {
        const modalData = getModalData();
        if (modalData?.type === "groupEvent") {
            setData((prev) => [modalData, ...prev]);
            setModalData(null);
        }
    }, [getModalData, setModalData]);

    // Fetch data function
    const fetchData = useCallback(
        async (id) => {
            if (isLoading || !hasMore) return;
            setIsLoading(true);
            abortControllerRef.current = new AbortController();
            const signal = abortControllerRef.current.signal;
            try {
                const response = await fetch(
                    `http://localhost:8080/api/groups/${groupId}/events/?offset=${id}`,
                    { credentials: "include", signal }
                );
                const result = await response.json();
                if (!response.ok) {
                    if (response.status === 403) {
                        setIsAccessible({ status: 403 });
                    }
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
                if (result.length === 0) {
                    setHasMore(false);
                } else {
                    if (result.length < 5) setHasMore(false);
                    setData((prevData) => [...prevData, ...result]);
                }
            } catch (err) {
                if (err.name === "AbortError") return;
                setError(err.message);
            } finally {
                setIsLoading(false);
            }
        },
        [groupId, setIsAccessible]
    );

    // Reset data and fetch initial page when groupId changes
    useEffect(() => {
        setData([]);
        setPage(0);
        setHasMore(true);
        setError(null);
        if (abortControllerRef.current) {
            abortControllerRef.current.abort();
        }
        fetchData(0);
    }, [groupId, fetchData]);

    // Infinite scroll observer
    useEffect(() => {
        if (!hasMore || isLoading || data.length === 0) return;

        observerRef.current = new IntersectionObserver(
            (entries) => {
                if (entries[0].isIntersecting) {
                    setPage((prevPage) => prevPage + 1);
                }
            },
            { threshold: 0.1 }
        );

        if (loadMoreRef.current) {
            observerRef.current.observe(loadMoreRef.current);
        }

        return () => {
            if (observerRef.current) {
                observerRef.current.disconnect();
            }
        };
    }, [hasMore, isLoading, data.length]);

    // Fetch data when page changes
    useEffect(() => {
        if (page > 0) {
            const id = data[data.length - 1]?.event_id || 0;
            fetchData(id);
        }
    }, [page, fetchData]);

    // Handle forbidden access
    if (isAccessible?.status === 403) {
        return (
            <section className="posts_container w-full h-full flex-col justify-center align-center">
                <img src="/forbidden-posts.svg" style={{ height: "100%" }} />
                <p className="text-xl font-semibold">You must become a member to see the events</p>
            </section>
        );
    }

    // Handle no data
    if (data.length === 0 && !isLoading) {
        return (
            <div className="flex justify-center">
                <img
                    className="w-half mx-auto"
                    src="/no-data-animate.svg"
                    alt="No data"
                />
            </div>
        );
    }

    return (
        <div className="list-container flex flex-wrap gap-2 align-center justify-center overflow-y-auto h-full">
            {data.map((event) => (
                <GroupEventCard {...event} key={event.event_id} />
            ))}
            {isLoading && <p className="text-center w-full">Loading...</p>}
            {hasMore && !isLoading && (
                <div ref={loadMoreRef} className="w-full" style={{ height: "20px" }}></div>
            )}
        </div>
    );
}