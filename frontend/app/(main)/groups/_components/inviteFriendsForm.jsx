"use client"

import { useState, useEffect, use } from 'react';


import UserCard from './userCard';
import { useModal } from '../../_context/ModalContext';
import { useRouter } from 'next/navigation';
import Loader from '../../_components/loader';

// InviteFriendForm component
const InviteFriendForm = ({ groupId }) => {
    const [followers, setFollowers] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const router = useRouter()
    useEffect(() => {

        async function handleGetFollowers() {
            try {
                const res = await fetch(`http://localhost:8080/api/groups/${groupId}/invitations/`, {
                    credentials: "include",
                })

                const result = await res.json()
                if (res.ok) {
                        setFollowers(result)
                } else {
                    if (result.status === 401) {
                        router.push("/login")
                        return
                    }
                        setError(result.error || "unexpected error while fetching followers")
                        return
                } 

            } catch (err) {
                setError("Failed to fetch followers. Please try again later.");
                console.error("Failed to fetch followers", err)
            } finally {
                setLoading(false)
            }
        }

        handleGetFollowers()
    }, [])


    
    if (loading) return <Loader />
    if (error) return  <Error error={error}/>

    return (
        <>
            <div style={{ maxHeight: '400px', overflowY: 'auto', marginBottom: '16px', paddingInline: ".5rem" }}>
                {
                    followers.length === 0 
                    ? <span>You currently have no followers available for invitation. Follow others to build your community.</span> 
                    : followers?.map((user) => (
                        <UserCard
                            key={user.id}
                            user={user}
                            groupId={groupId}

                        />
                ))}
            </div>

        </>
    );
};

export default InviteFriendForm;