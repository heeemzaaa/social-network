import React from 'react'
import InfiniteList from '../../_components/infiniteList';

export default function GroupEventCardList() {
    const getUrl = (page) => {
        const params = new URLSearchParams({
            offset: page * 20,
        });
        return `http://localhost:8080/api/groups/${gorupdId}/?${params.toString()}`;
    }

    return (
        <InfiniteList
            getUrl={getUrl}
            itemsPerPage={20}
            renderItem={({ item, index }) => <GroupCard key={index} {...item} />}
        />
    );
}
