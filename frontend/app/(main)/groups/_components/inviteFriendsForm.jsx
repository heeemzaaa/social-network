"use client"

import { useState, useEffect, use } from 'react';


import UserCard from './userCard';
import { useModal } from '../../_context/ModalContext';

// InviteFriendForm component
const InviteFriendForm = ({ groupId }) => {
    console.log("grp id ", groupId);
    const [followers, setFollowers] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const { closeModal } = useModal()

    useEffect(() => {
     
        async function handleGetFollowers() {
            try {
                const res = await fetch(`http://localhost:8080/api/groups/${groupId}/invitations/`, {
                    credentials: "include",
                })

                if (res.ok) {
                    const result = await res.json()

                    setFollowers(result)
                    console.log("followers",result);

                }

            } catch (err) {
                console.error("Failed to fetch followers", err)
            } finally {
                setLoading(false)
            }
        }

        handleGetFollowers()
    }, [])





    return (
        <>

            <h2 style={{ fontSize: '24px', fontWeight: '600', color: '#111827', marginBottom: '16px' }}>
                Invite Friends
            </h2>
            {loading && <p style={{ color: '#374151', fontSize: '16px' }}>Loading followers...</p>}
            {error && <p style={{ color: '#dc2626', fontSize: '16px' }}>{error}</p>}

            <div style={{ maxHeight: '400px', overflowY: 'auto', marginBottom: '16px', paddingInline: ".5rem" }}>
                {followers.map((user) => (
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