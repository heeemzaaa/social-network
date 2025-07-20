import { useEffect, useState, useRef, useCallback } from "react";

export default function useInfiniteScroll({
    getUrl,
    initialPage = 0,
    itemsPerPage = 20,
}) {
    const [data, setData] = useState([]);
    const [page, setPage] = useState(initialPage);
    const [isLoading, setIsLoading] = useState(false);
    const [hasMore, setHasMore] = useState(true);
    const [error, setError] = useState(null);
    const sentinelRef = useRef(null);
    const abortControllerRef = useRef(null);
    const observerRef = useRef(null);

    // Fetch data function
    const fetchData = async (currentPage) => {
        if (isLoading || !hasMore) return;
        setIsLoading(true);
        abortControllerRef.current = new AbortController();
        const signal = abortControllerRef.current.signal;
        try {
            const url = getUrl(currentPage);
            const response = await fetch(url, { credentials: "include", signal });
            const result = await response.json();
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            if (result.length === 0) {
                setHasMore(false);
            } else {
                setData((prevData) => [...prevData, ...result]);
            }
        } catch (err) {
            if (err.name === "AbortError") {
                return;
            }
            console.error(err);
            setError(err.message);
        } finally {
            setIsLoading(false);
        }
    };

    // Reset data when getUrl changes
    useEffect(() => {
        setData([]);
        setPage(initialPage);
        setHasMore(true);
        setError(null);
        fetchData(initialPage);

        return () => {
            if (abortControllerRef.current) {
                abortControllerRef.current.abort();
                abortControllerRef.current = null;
            }
        };
    }, [getUrl]);

    // Callback to update sentinelRef and re-observe
    const setSentinelRef = useCallback((node) => {
        console.log("Sentinel ref updated:", node);
        // Update sentinelRef
        sentinelRef.current = node;

        // Re-observe if observer exists and node is valid
        if (observerRef.current && node) {
            observerRef.current.observe(node);
        }
    }, []);




    // Set up Intersection Observer
    useEffect(() => {
        // Clean up existing observer
        if (observerRef.current && sentinelRef.current) {
            observerRef.current.unobserve(sentinelRef.current);
        }

        // Create new observer
        observerRef.current = new IntersectionObserver(
            (entries) => {
                if (entries[0].isIntersecting && hasMore && !isLoading) {
                    setPage((prevPage) => prevPage + 1);
                }
            },
            { threshold: 0.1 }
        );

        // Observe the sentinel if it exists
        if (sentinelRef.current) {
            observerRef.current.observe(sentinelRef.current);
        }

        // Cleanup on unmount or when dependencies change
        return () => {
            if (observerRef.current && sentinelRef.current) {
                observerRef.current.unobserve(sentinelRef.current);
            }
            observerRef.current = null;
        };
    }, [hasMore, isLoading, getUrl]);

    // Fetch data when page changes
    useEffect(() => {
        if (page > initialPage) {
            fetchData(page);
        }
    }, [page]);

    return { data, isLoading, error, hasMore, sentinelRef: setSentinelRef };
}