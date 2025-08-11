import React, { useEffect, useState } from 'react'
import Avatar from '../../_components/avatar';
import Button from '@/app/_components/button';

import { useNotification } from "../../_context/NotificationContext";


export default function UserCard({ user, groupId }) {
    const [inviteState, setInviteState] = useState(user.invited)

    const { showNotification } = useNotification();
    

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

            if (!res.ok) console.error("Failed to send the request")
            const data = await res.json();

            console.log(" ==>> ", data);
            if (data.Message === 'ERROR!! already a member!' || data.Message === 'ERROR!! it is not from your followers!' || data.Message === 'ERROR!! Invitation not found') {
                console.warn(data.Message.split("!!")[1])
                showNotification({ Content: data.Message.split("!!")[1], Status: "error" });

                // should access to the followers list and remove the user from the list

                if (data.Message !== 'ERROR!! Invitation not found') return
            }

            setInviteState(inviteState === 0 ? 1: 0)

            // // should all this part be in the server side? Actions

        } catch (err) {
            console.log(err);
        }
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
                        <div>
                            <p style={{ color: '#1f2937', fontWeight: '500', fontSize: '16px', marginLeft: "5px" }}>{user.fullname}</p>
                            <p className='text-sm '>@{user.nickname}</p>
                        </div>
                    </div>
                </div>
                {
                    inviteState === 0 ? <Button onClick={() => handleInviteCancelButtons()} >
                        Invite
                    </Button> : <Button variant={"btn-danger"} onClick={() => handleInviteCancelButtons()} >
                        Cancel
                    </Button>
                }



            </div>
        </section>
    );
};
