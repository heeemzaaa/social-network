"use client";
import { useEffect, useState } from "react";
import PostCard from "./postCard";
import { useModal } from "../../_context/ModalContext";

export default function PostCardList() {
    const [posts, setPosts] = useState([])
    const { getModalData, setModalData } = useModal()
    // run every time the modaldata changes
    useEffect(() => {
        let postData = getModalData()
        if (postData?.type !== 'post') return;

        setPosts((prev) => {
            console.log("prev: ", prev)
            if (!prev) {
                return [postData]
            } else {
                return [postData, ...prev]
            }
        })
        // clear modal data 
        setModalData(null)
    }, [setModalData])
    // execute in first rendring
    useEffect(() => {
        async function fetchPosts() {
            console.log("fetch posts here.");
            try {
                const resp = await fetch("http://localhost:8080/api/posts", {
                    method: "GET",
                    credentials: "include",
                });

                if (!resp.ok) {
                    console.log("error fetching posts 1");
                    return;
                }
                const data = await resp.json();
                console.log(data)
                setPosts(data);
            } catch (error) {
                console.log("error fetching posts", error);
            }
        }

        fetchPosts();
    }, []);

    return (
        <div className="list-container " style={{ overflowY: "auto" }}>
            {posts?.map((post) => (
                <PostCard key={post.id} {...post} />
            ))}
        </div>
    );
}
