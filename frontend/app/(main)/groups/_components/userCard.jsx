import React, { useEffect, useState } from 'react'
import Avatar from '../../_components/avatar';
import Button from '@/app/_components/button';
import { useNotification } from '../../_context/NotificationContext';
import { useRouter } from 'next/navigation';


export default function UserCard({ user, groupId }) {
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

            const data = await response.json()
            if (!response.ok) {
                if (data.status === 401) {
                    router.push("/login")
                    return
                }
                setError(data.error)
                return
            }
            if (data.Message === 'ERROR!! You are already a member!') {
                showNotification({ Content: `already a member!`, Status: "error" });
                // should access to the followers list and remove the user from the list

                console.warn(`already a member!`)
                return
            } else if (data.Message === 'Invitation not found') {
                console.warn(`Invitation not found !!`)
                showNotification({ Content: `Invitation not found`, Status: "error" });
            }
            inviteState === 0 ? setInviteState(1) : setInviteState(0)
        } catch (err) {
            setError(err)
            console.log(err);
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
