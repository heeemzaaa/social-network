import { useCallback, useEffect, useState, useRef } from "react";
import Button from "@/app/_components/button";
import GroupEventCard from "./groupEventCard";
import { useModal } from "../../_context/ModalContext";

export default function GroupEventCardList({ groupId }) {
    const [data, setData] = useState([]);
    const [page, setPage] = useState(0);
    const [isLoading, setIsLoading] = useState(false);
    const [hasMore, setHasMore] = useState(true);
    const [error, setError] = useState(null);
    const abortControllerRef = useRef(null);

    const { getModalData, setModalData } = useModal()

    useEffect(() => {
        let data = getModalData()
        console.log("Modal Data:", data)
        if (data?.type === "groupEvent") {
            setData(prev => [data, ...prev])
        }
    }, [setModalData])

    const getUrl = useCallback(
        (page) => {
            const params = new URLSearchParams({
                offset: page * 20,
            });
            return `http://localhost:8080/api/groups/${groupId}/events/?${params.toString()}`;
        },
        [groupId]
    );

    // Fetch data function
    const fetchData = useCallback (
        async (currentPage) => {
            if (isLoading || !hasMore) return;
            setIsLoading(true);
            abortControllerRef.current = new AbortController();
            const signal = abortControllerRef.current.signal;
            try {
                const url = getUrl(currentPage);
                console.log("url: ", url)
                const response = await fetch(url, { credentials: "include", signal });
                const result = await response.json();
                if (!response.ok) {
                    console.log(result)
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
                console.log(result)

                if (result.length === 0) {
                    setHasMore(false); // No more data to fetch
                } else {
                    console.log(data)
                    if (result.length < 20) setHasMore(false);
                    setData((prevData) => [...prevData, ...result]); // Append new data
                }
            } catch (err) {
                if (err.name === "AbortError") {
                    return; // Ignore AbortError
                }
                console.error(err);
                setError(err.message);
            } finally {
                setIsLoading(false);
            }
        },
        [getUrl,groupId]
    );

    // Reset data and fetch initial page when groupId changes
    useEffect(() => {
        setData([]);
        setPage(0);
        setHasMore(true);
        setError(null);
        fetchData(0);

        return () => {
            if (abortControllerRef.current) {
                abortControllerRef.current.abort();
                abortControllerRef.current = null;
            }
        };
    }, [groupId, fetchData]);

    // Fetch data when page changes
    useEffect(() => {
        if (page > 0) {
            fetchData(page);
        }
    }, [page, fetchData]);

    // Load more handler
    const loadMore = () => {
        setPage((prevPage) => prevPage + 1);
    };

    if (error) return <p className="text-danger text-center">Error: {error}</p>;
    if (data.length === 0 && !isLoading) return (
        <img
            className="w-half mx-auto"
            src="/no-data-animate.svg"
            alt="No data"
        />
    );

    return (
        <div className="list-container flex flex-wrap gap-2 align-center justify-center overflow-y-auto h-full">
            {data.map((event, index) => (
                <GroupEventCard {...event} key={index} />
            ))}
            {isLoading && <p className="text-center w-full">Loading...</p>}
            {hasMore && !isLoading && (
                <div className="w-full" style={{ textAlign: "center" }}>
                    <Button variant="btn-tertiary" onClick={loadMore}>
                        Load More...
                    </Button>
                </div>
            )}
        </div>
    );
}