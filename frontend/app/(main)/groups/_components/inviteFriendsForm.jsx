"use client"

import { useState, useEffect, use } from 'react';


import UserCard from './userCard';
import { useModal } from '../../_context/ModalContext';

// InviteFriendForm component
const InviteFriendForm = ({ groupId }) => {
    const [followers, setFollowers] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

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
                        setError(result.error || "unexpected error while fetching followers")
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


    
    if (loading) return <p style={{ color: '#374151', fontSize: '16px' }}>Loading followers...</p>
    if (error) return  <p style={{ color: '#dc2626', fontSize: '16px', padding:"1rem" }}>{error}</p>

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