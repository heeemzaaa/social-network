import React, { useEffect, useState } from 'react'
import Avatar from '../../_components/avatar';
import Button from '@/app/_components/button';

import { useNotification } from "../../_context/NotificationContext";
import { useUserContext } from "../../_context/userContext";
import { useRouter } from 'next/navigation';

export default function UserCard({ user, groupId }) {
    const { sendSocketMessage, authenticatedUser } = useUserContext();
    const [inviteState, setInviteState] = useState(user.invited)
    const { showNotification } = useNotification()
    const [error, setError] = useState(null)
    const router = useRouter()
    // let's create here the function that toggles the state of the button with the same
    // way as hamza 
    async function handleInviteCancelButtons() {
        let endpoint = `http://localhost:8080/api/groups/${groupId}/invitations/`
        let method = inviteState === 0 ? 'POST' : 'DELETE'
        try {
            const res = await fetch(endpoint, {
                method: method,
                headers: { 'Content-Type': 'application/json' },
                credentials: 'include',
                body: JSON.stringify({ 'id': user.id }),
            })


            const data = await res.json();
            if (!res.ok) {
                if (data.status === 401) {
                    router.push("/login")
                    return
                }
                showNotification({
                    Content: "Failed to send request",
                    Status: "error"
                });
                setError(data.error)
                return
            }
            console.log(" ==>> ", data);

            // Handle error messages
            if (data.Message === 'ERROR!! already a member!' ||
                data.Message === 'ERROR!! it is not from your followers!' ||
                data.Message === 'ERROR!! Invitation not found') {

                console.warn(data.Message.split("!!")[1])
                showNotification({
                    Content: data.Message.split("!!")[1],
                    Status: "error"
                });

                if (data.Message !== 'ERROR!! Invitation not found') return
                // should remove the user from the list
            }

            if (method === 'POST') {
                sendSocketMessage({
                    type: "notification",
                    Notification: data,
                });
            }

            setInviteState(inviteState === 0 ? 1 : 0);
            console.log("inviteState: ", inviteState)
            showNotification({
                Content: inviteState === 0 ? "Invitation sent!" : "Invitation cancelled)",
                Status: "success"
            });

        } catch (err) {
            console.error("Request error:", err);
            showNotification({
                Content: "Something went wrong. Please try again.",
                Status: "error"
            });
        }
    }

    if (error) {
        return (
            <Error error={error} />
        )
    }
    return (
        <section className='user_card p2 flex justify-start rounded-lg shadow-md m1' >
            <div
                style={{
                    display: 'flex',
                    justifyContent: "space-between",
                    alignItems: 'center',
                    borderRadius: '8px',
                    cursor: 'pointer',
                    transition: 'background-color 0.2s',
                    width: '300px'
                }}
            >

                <div className='flex gap-1'>
                    <Avatar size={42} img={user.avatar} />
                    <div style={
                        {
                            display: "flex",
                            justifyContent: "space-between",
                            gap: '16px'
                        }

                    }>
                        <div className='flex-col justify-center' >
                            <p style={{ color: '#1f2937', fontWeight: '500', fontSize: '16px', marginLeft: "5px" }}>{user.fullname}</p>
                            {user.nickname && <p className='text-sm '>@{user.nickname}</p>}
                        </div>
                    </div>
                </div>
                {
                    inviteState === 0 ? <Button onClick={handleInviteCancelButtons} >
                        Invite
                    </Button> : <Button variant={"btn-danger"} onClick={handleInviteCancelButtons} >
                        Cancel
                    </Button>
                }
            </div>
        </section>
    );
};