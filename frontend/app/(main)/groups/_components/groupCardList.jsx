import { useEffect } from "react";
import InfiniteList from "../../_components/infiniteList";
import GroupCard from "./groupCard";

export default function GroupCardList({ filter }) {
    const getUrl = (page) => {
        const params = new URLSearchParams({
            filter,
            offset: page * 5,
            limit: 5,
        });
        return `http://localhost:8080/api/groups?${params.toString()}`;
    };

    useEffect(() => {
    }, [filter])

    return (
        <InfiniteList
            getUrl={getUrl}
            itemsPerPage={5}
            renderItem={({ item, index }) => <GroupCard key={index} {...item} />}
        />
    );
}
