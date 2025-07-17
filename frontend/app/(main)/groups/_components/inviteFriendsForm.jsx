"use client"

import { useState, useEffect } from 'react';
import { useActionState } from 'react';
import Button from '@/app/_components/button';
import { inviteUsersAction } from '@/app/_actions/group';
import UserCard from './userCard';
import { useModal } from '../../_context/ModalContext';

const fetchFollowers = async () => {
    return [
        { id: 1, name: 'Alice Smith', avatar: 'https://i.pravatar.cc/150?img=1' },
        { id: 2, name: 'Bob Johnson', avatar: 'https://i.pravatar.cc/150?img=2' },
        { id: 3, name: 'Charlie Brown', avatar: 'https://i.pravatar.cc/150?img=3' },
        { id: 4, name: 'Diana Lee' },
        { id: 5, name: 'Alice Smith', avatar: 'https://i.pravatar.cc/150?img=1' },
        { id: 6, name: 'Bob Johnson', avatar: 'https://i.pravatar.cc/150?img=2' },
        { id: 7, name: 'Charlie Brown', avatar: 'https://i.pravatar.cc/150?img=3' },
        { id: 8, name: 'Diana Lee', avatar: 'https://i.pravatar.cc/150?img=4' },
        { id: 9, name: 'Alice Smith', avatar: 'https://i.pravatar.cc/150?img=1' },
        { id: 10, name: 'Bob Johnson', avatar: 'https://i.pravatar.cc/150?img=2' },
        { id: 11, name: 'Charlie Brown', avatar: 'https://i.pravatar.cc/150?img=3' },
        { id: 12, name: 'Diana Lee', avatar: 'https://i.pravatar.cc/150?img=4' },
    ];
};

// InviteFriendForm component
const InviteFriendForm = ({ groupId }) => {
    const [state, formAction, isPending] = useActionState(inviteUsersAction, {});
    const [selectedUsers, setSelectedUsers] = useState([]);
    const [followers, setFollowers] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    const { closeModal } = useModal()

    useEffect(() => {
        if (!state.message) return;
        closeModal()
    }, [state])

    // Fetch followers on mount
    useEffect(() => {
        const loadFollowers = async () => {
            try {
                const data = await fetchFollowers();
                setFollowers(data);
                setLoading(false);
            } catch (err) {
                setError('Failed to load followers');
                setLoading(false);
            }
        };
        loadFollowers();
    }, []);

    // Handle user selection
    const handleSelect = (userId) => {
        setSelectedUsers((prev) =>
            prev.includes(userId) ? prev.filter((id) => id !== userId) : [...prev, userId]
        );
    };

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
            {state.message && (
                <p
                    style={{
                        color: state.success ? '#16a34a' : '#dc2626',
                        fontSize: '16px',
                        marginBottom: '16px'
                    }}
                >
                    {state.message}
                </p>
            )}
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
                <Button>
                    {isPending ? 'Inviting...' : `Invite ${selectedUsers.length > 0 ? `(${selectedUsers.length})` : ''}`}
                </Button>
            </div>
        </form>
    );
};

export default InviteFriendForm;