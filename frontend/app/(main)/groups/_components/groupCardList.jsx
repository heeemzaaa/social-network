import { useCallback, useEffect } from "react";
import InfiniteList from "../../_components/infiniteList";
import GroupCard from "./groupCard";

export default function GroupCardList({ filter }) {
    const getUrl = useCallback(
        (page) => {
            console.log("fetching data for: ", filter,Math.random())
            const params = new URLSearchParams({
                filter,
                offset: page *20,
            });
            return `http://localhost:8080/api/groups?${params.toString()}`;
        },
        [filter]
    );

    return (
        <InfiniteList
            getUrl={getUrl}
            itemsPerPage={20}
            renderItem={({ item, index }) => <GroupCard key={index} {...item} />}
        />
    );
}
