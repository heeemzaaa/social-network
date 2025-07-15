import { useEffect, useState, useRef } from "react";

export default function useInfiniteScroll({
    getUrl,
    initialPage = 0,
    itemsPerPage = 10,
}) {
    const [data, setData] = useState([]);
    const [page, setPage] = useState(initialPage);
    const [isLoading, setIsLoading] = useState(false);
    const [hasMore, setHasMore] = useState(true);
    const [error, setError] = useState(null);
    const sentinelRef = useRef(null);

    // Fetch data function
    const fetchData = async (currentPage) => {
        if (isLoading || !hasMore) return;
        setIsLoading(true);
        try {
            const url = getUrl(currentPage);
            const response = await fetch(url, {
                credentials: "include",
            });
            const result = await response.json();
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            if (result.length === 0) {
                setHasMore(false); // No more data to fetch
            } else {
                setData((prevData) => [...prevData, ...result]); // Append new data
            }
        } catch (err) {
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
    }, [getUrl]);

    // Set up Intersection Observer
    useEffect(() => {
        const observer = new IntersectionObserver(
            (entries) => {
                if (entries[0].isIntersecting && hasMore && !isLoading) {
                    setPage((prevPage) => prevPage + 1);
                }
            },
            { threshold: 0.1 } // Trigger when 10% of sentinel is visible
        );

        if (sentinelRef.current) {
            observer.observe(sentinelRef.current);
        }

        return () => {
            if (sentinelRef.current) {
                observer.unobserve(sentinelRef.current);
            }
        };
    }, [hasMore, isLoading]);

    // Fetch data when page changes
    useEffect(() => {
        if (page > initialPage) {
            fetchData(page);
        }
    }, [page]);

    return { data, isLoading, error, hasMore, sentinelRef };
}