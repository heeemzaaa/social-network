import { useEffect, useState } from "react";
import GroupCard from "./groupCard";

export default function GroupCardList({ filter }) {
    console.log(filter)
    let [data, setData] = useState(null)
    let [offset, setOffset] = useState(0)
    let [isLoading, setIsLoading] = useState(true)
    let [error, setError] = useState(null)

    useEffect(() => {
        console.log("inside use effect");
        const fetchData = async () => {
            try {
                const response = await fetch(`http://localhost:8080/api/groups/?filter=${filter}&offset=${offset}`,{
                    credentials: "include",
                });
                const result = await response.json();
                if (!response.ok) {
                    console.log(result)
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
                console.log("groups: ",  result)
                setData(result);
            } catch (err) {
                console.error(err)
                setError(JSON.stringify(err));
            } finally {
                setIsLoading(false);
            }
        };
        fetchData();
    }, [filter, offset])

    if (isLoading) return (<p>Loading...</p>);
    if (error) return (<p>Error: {error}</p>);
    return (
        <div className="grp-cards-list flex flex-wrap  gap-4 justify-center overflow-y-auto">
            {
                data.length > 0 ?
                    data.map((grp, index) => <GroupCard key={index} {...grp} />)
                    : <img className="w-half" src="no-data-animate.svg" alt="" />
            }
        </div>
    )
}
