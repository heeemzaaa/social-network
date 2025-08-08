"use client";
import PostCard from "./postCard";
import { useModal } from "../../_context/ModalContext";
import {
    useEffect,
    useState
} from "react";

export default function PostCardList() {
    const [posts, setPosts] = useState([])
    const { getModalData, setModalData } = useModal()
    useEffect(() => {
        let postData = getModalData()
        if (postData?.type !== 'post') return;

        setPosts((prev) => {
            if (!prev) {
                return [postData]
            } else {
                return [postData, ...prev]
            }
        })
        setModalData(null)
    }, [setModalData])

    useEffect(() => {
        async function fetchPosts() {
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
