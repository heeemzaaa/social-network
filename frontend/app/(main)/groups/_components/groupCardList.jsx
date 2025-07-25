import { useCallback, useEffect, useState, useRef } from "react";
import GroupCard from "./groupCard";
import Button from "@/app/_components/button";
import { useModal } from "../../_context/ModalContext";

export default function GroupCardList({ filter }) {
    const [data, setData] = useState([]);
    const [page, setPage] = useState(0);
    const [isLoading, setIsLoading] = useState(false);
    const [hasMore, setHasMore] = useState(true);
    const [error, setError] = useState(null);
    const abortControllerRef = useRef(null);

    const {getModalData, setModalData} = useModal()

    useEffect(()=>{
        let data = getModalData()
        if (data?.type === "groupCard" && filter === "owned") {
            setData(prev=> [data,...prev])
        }
    },[setModalData])

    const fetchData = useCallback(
        async (currentPage) => {
            if (isLoading || !hasMore) return;
            setIsLoading(true);
            abortControllerRef.current = new AbortController();
            const signal = abortControllerRef.current.signal;
            try {
                const response = await fetch(`http://localhost:8080/api/groups?filter=${filter}&offset=${currentPage * 20}`,
                    { credentials: "include", signal });
                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
                const result = await response.json();
                if (result.length === 0) {
                    setHasMore(false); // No more data to fetch
                } else {
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
        [filter]
    );

    // Reset data and fetch initial page when filter changes
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
    }, [filter, fetchData]);

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
    if (data.length === 0) (
        <img
            className="w-half mx-auto"
            src="/no-data-animate.svg"
            alt="No data"
        />
    )
    return (
        <div className="list-container flex flex-wrap gap-4 justify-center items-start overflow-y-auto">
            {
                data.map((item, index) => <GroupCard key={item.id || index} type={filter} {...item} />)
            }
            {isLoading && <p className="text-center w-full">Loading...</p>}
            {hasMore && !isLoading && (
                <div className="w-full" style={{ textAlign: "" }}>
                    <Button variant={"btn-tertiary"} onClick={loadMore}> Load More... </Button>
                </div>
            )}
        </div>
    );
}