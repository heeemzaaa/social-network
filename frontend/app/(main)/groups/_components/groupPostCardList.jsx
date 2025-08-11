import { useCallback, useEffect, useState, useRef } from "react";
import PostCard from "../../_components/posts/postCard";
import { useModal } from "../../_context/ModalContext";

export default function GroupPostCardList({ groupId, setIsAccessible, isAccessible }) {
    const [data, setData] = useState([]);
    const [page, setPage] = useState(0);
    const [isLoading, setIsLoading] = useState(false);
    const [hasMore, setHasMore] = useState(true);
    const [error, setError] = useState(null);
    const observerRef = useRef(null);
    const loadMoreRef = useRef(null);

    const { getModalData, setModalData } = useModal();

    // Handle modal data for new group posts
    useEffect(() => {
        const modalData = getModalData();
        if (modalData?.type === "groupPost") {
            setData((prev) => [modalData, ...prev]);
            setModalData(null);
        }
    }, [getModalData, setModalData]);

    // Fetch data function
    const fetchData = useCallback(
        async (id) => {
            if (isLoading || !hasMore) return;
            setIsLoading(true);
            try {
                const response = await fetch(
                    `http://localhost:8080/api/groups/${groupId}/posts/?offset=${id}`,
                    { credentials: "include" }
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
                    if (result.length < 20) setHasMore(false); // Adjust based on API page size
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
    }, [hasMore, isLoading, data]);

    // Fetch data when page changes
    useEffect(() => {
        if (page > 0) {
            const id = data[data.length - 1]?.id || 0;
            fetchData(id);
        }
    }, [page, fetchData, data]);

    // Handle forbidden access
    if (isAccessible?.status === 403) {
        return (
            <section className="posts_container w-full h-full flex-col justify-center align-center">
                <img src="/forbidden-posts.svg" style={{ height: "100%" }} />
                <p className="text-xl font-semibold">You must become a member to see the posts</p>
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
        <div className="list-container flex align-start flex-wrap gap-4 justify-center overflow-y-auto">
            {data.map((item) => (
                <PostCard {...item} key={item.id} groupID={groupId || ""} />
            ))}
            {isLoading && <p className="text-center w-full">Loading...</p>}
            {hasMore && !isLoading && (
                <div ref={loadMoreRef} className="w-full" style={{ height: "20px" }}></div>
            )}
        </div>
    );
}