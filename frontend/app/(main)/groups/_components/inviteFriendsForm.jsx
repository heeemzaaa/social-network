"use client"

import { useState, useEffect } from 'react';
import { useActionState } from 'react';
import Button from '@/app/_components/button';
import { inviteUsersAction } from '@/app/_actions/group';
import UserCard from './userCard';
import { useModal } from '../../_context/ModalContext';

// InviteFriendForm component
const InviteFriendForm = ({ groupId }) => {
    const [state, formAction, isPending] = useActionState(inviteUsersAction, {});
    const [selectedUsers, setSelectedUsers] = useState([]);
    const [followers, setFollowers] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const { closeModal } = useModal()

    useEffect(() => {
        async function handleGetFollowers() {
            try {
                const res = await fetch(`http://localhost:8080/api/profile/d7704b4f-6428-441b-93c8-3601ba3242fc/connections/followers`, {
                    credentials: "include",
                })

                if (res.ok) {
                    const result = await res.json()
                    console.log(result);
                    setFollowers(result)

                }

            } catch (err) {
                console.error("Failed to fetch followers", err)
            } finally {
                setLoading(false)
            }
        }

        handleGetFollowers()
    }, [])

    useEffect(() => {
        if (!state.message) return;
        closeModal()
    }, [state])

    // Handle user selection
    const handleSelect = (userId) => {
        setSelectedUsers((prev) =>
            prev.includes(userId) ? prev.filter((id) => id !== userId) : [...prev, userId]
        );
    };
     
    if (loading) {
        
    }
    return (
        <form
            action={formAction}
            style={{ padding: '24px', maxWidth: '600px', margin: '0 auto' }}
        >
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
                        isSelected={selectedUsers.includes(user.id)}
                        onSelect={handleSelect}
                    />
                ))}
            </div>
            <input type="hidden" name="groupId" value={groupId} />
            <div style={{ display: 'flex', justifyContent: 'flex-end', gap: '8px' }}>
                <Button style={{ width: "100%" }}>
                    {isPending ? 'Inviting...' : `Invite ${selectedUsers.length > 0 ? `(${selectedUsers.length})` : ''}`}
                </Button>
            </div>
            {state.error && <p style={{ color: '#dc2626', fontSize: '16px' }}>{state.error}</p>}
        </form>
    );
};

export default InviteFriendForm;