import useInfiniteScroll from "@/app/_hooks/useInfiniteScroll";

export default function InfiniteList({
    getUrl,
    initialPage = 0,
    itemsPerPage = 20,
    renderItem,
    noDataImage = "no-data-animate.svg",
    className = "list-container flex flex-wrap gap-4 justify-center overflow-y-auto",
}) {

    const { data, isLoading, error, hasMore, sentinelRef } = useInfiniteScroll({
        getUrl,
        initialPage,
        itemsPerPage,
    })

    if (error) return <p className="text-danger text-center">Error: {error}</p>;
    return (
        <div className={className}>
            {
                data.length > 0 ? (
                    data.map((item, index) => renderItem({ item, index }))
                ) : (
                    !isLoading && (
                        <img className="w-half" src={noDataImage} alt="No data" />
                    )
                )}
            {isLoading && <p className="text-center">Loading...</p>}
            {hasMore && <div ref={sentinelRef} style={{ height: "20px" }} />}
        </div>
    );
}