import { useEffect, useState } from "react";
import GroupCard from "./groupCard";

let groups = [
    // {
    //     title: "Hiking Enthusiasts",
    //     description: "A group for people who love exploring nature trails and going on hikes.",
    //     membersCount: 128
    // },
    // {
    //     title: "JavaScript Developers",
    //     description: "A place to discuss JavaScript, share resources, and collaborate on projects.",
    //     membersCount: 342
    // },
    // {
    //     title: "Bookworms United",
    //     description: "For those who love reading and sharing book recommendations.",
    //     membersCount: 89
    // },
    // {
    //     title: "Fitness & Wellness",
    //     description: "A community focused on health, fitness routines, and mindfulness.",
    //     membersCount: 210
    // },
    // {
    //     title: "Digital Nomads",
    //     description: "Connect with remote workers traveling the world and sharing tips.",
    //     membersCount: 154
    // },
    // {
    //     title: "Photography Pros",
    //     description: "Share your best shots, learn techniques, and review gear with fellow photographers.",
    //     membersCount: 277
    // },
    // {
    //     title: "Indie Game Developers",
    //     description: "A supportive space for indie game devs to share progress, feedback, and ideas.",
    //     membersCount: 196
    // },
    // {
    //     title: "Startup Founders Hub",
    //     description: "Connect with fellow entrepreneurs, share startup stories, and seek advice.",
    //     membersCount: 311
    // },
    // {
    //     title: "Sustainable Living",
    //     description: "Tips and discussions around eco-friendly habits and sustainable practices.",
    //     membersCount: 134
    // },
    // {
    //     title: "Language Learners",
    //     description: "Practice and exchange languages with people from around the world.",
    //     membersCount: 402
    // }
];


export default function GroupCardList({ filter }) {
    let [data, setData] = useState(null)
    let [offset, setOffset] = useState(0)
    let [isLoading, setIsLoading] = useState(true)
    let [error, setError] = useState(null)

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await fetch(`http://localhost:8080/api/groups/?filter=${filter}&offset=${offset}`,{
                    credentials: "include",
                });
                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
                const result = await response.json();
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
    }, [])

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
