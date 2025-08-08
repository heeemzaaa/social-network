import { useCallback, useEffect, useState, useRef, use } from "react";
import Button from "@/app/_components/button";
import GroupEventCard from "./groupEventCard";
import { useModal } from "../../_context/ModalContext";

export default function GroupEventCardList({ groupId, setIsAccessible, isAccessible }) {
    const [data, setData] = useState([]);
    const [page, setPage] = useState(0);
    const [isLoading, setIsLoading] = useState(false);
    const [hasMore, setHasMore] = useState(true);
    const [error, setError] = useState(null);
    const abortControllerRef = useRef(null);
    const observerRef = useRef(null)
    const loadMoreRef = useRef(null)

    const { getModalData, setModalData } = useModal()

    useEffect(() => {
        let modalData = getModalData()
        if (modalData?.type === "groupEvent") {
            setData(prev => [modalData, ...prev])
            setModalData(null)
        }
    }, [setModalData])

    const getUrl = useCallback(
        (page) => {
            const params = new URLSearchParams({
                offset: page * 3,
            });
            console.log("offset here", params);
            return `http://localhost:8080/api/groups/${groupId}/events/?${params.toString()}`;
        },
        [groupId]
    );

    // Fetch data function
    const fetchData = useCallback(
        async (currentPage) => {
            if (isLoading || !hasMore) return;
            setIsLoading(true);
            abortControllerRef.current = new AbortController();
            const signal = abortControllerRef.current.signal;
            try {
                const url = getUrl(currentPage);
                const response = await fetch(url, { credentials: "include", signal });
                const result = await response.json();
                if (!response.ok) {
                    console.log("result inside the response", response);
                    if (response.status == 403) setIsAccessible(response)
                    throw new Error(`HTTP error! status: ${response.status}`);
                }

                if (result.length === 0) {
                    setHasMore(false); // No more data to fetch
                } else {
                    if (result.length < 3) setHasMore(false);
                    setData((prevData) => [...prevData, ...result]); // Append new data
                }
            } catch (err) {
                if (err.name === "AbortError") {
                    return; // Ignore AbortError
                }
                setError(err.message);
            } finally {
                setIsLoading(false);
            }
        },
        [getUrl, groupId]
    );

    // Reset data and fetch initial page when groupId changes
    useEffect(() => {
        setData([]);
        setPage(0);
        setHasMore(true);
        setError(null);
        fetchData(0);
    }, [groupId, fetchData]);

    // here will be adding the useEffect for the infinite scroll 
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

    // Fetch data when page changes
    useEffect(() => {
        if (page > 0) {
            let id = data[data.length-1]?.event_id
            fetchData(id);
        }
    }, [page]);

 


    console.log("data", data)

    if (isAccessible?.status == 403) {
        return (
            <section className='posts_container w-full h-full flex-col justify-center align-center'>
                <img src="/forbidden-posts.svg" style={{ height: '100%' }} />
                <p className='text-xl font-semibold'>You must become a member to see the events</p>
            </section>
        )
    }

    if (data.length === 0 && !isLoading) return (
        <img
            className="w-half mx-auto"
            src="/no-data-animate.svg"
            alt="No data"
        />
    );
            console.log(data)

    data.map((event)=> {
        console.log(event, event.event_id)
    })

    return (
        <div className="list-container flex flex-wrap gap-2 align-center justify-center overflow-y-auto h-full">
            {data.map((event) => (
                <GroupEventCard {...event} key={event.event_id} />
            ))}
            {isLoading && <p className="text-center w-full">Loading...</p>}
            {hasMore && !isLoading && (
                <div className="w-full" style={{ textAlign: "center" }}>
                    <Button variant="btn-tertiary" >
                        Load More...
                    </Button>
                </div>
            )}
        </div>
    );
}