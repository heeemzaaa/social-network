import React, { use, useEffect, useState } from 'react'
import Avatar from '../../_components/avatar';
import Button from '@/app/_components/button';
import { useActionState } from 'react';
import { inviteUserAction, CancelInvitationAction } from '@/app/_actions/group';


export default function UserCard({ user, groupId }) {
    const [inviteActionState, inviteAction] = useActionState(inviteUserAction, {});
    const [cancelState, cancelAction] = useActionState(CancelInvitationAction, {});
    const [inviteState, setInviteState] = useState(user.invited)
    console.log("staaaaaaaaaate", inviteState);


    useEffect(()=>{
        if (inviteActionState?.message) setInviteState(1)
        if (cancelState?.message) setInviteState(0)
    },[inviteActionState, cancelState])

    console.log("invite staaaaaaaaaaate: ", inviteState)
    return (    
        <form
            action={inviteState == 0 ? inviteAction : cancelAction}>
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
                    <input
                        name='user_id'
                        defaultValue={user.id}
                        style={{ marginRight: '12px' }}
                        hidden
                    />
                    <input type="hidden" name="groupId" value={groupId} />

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
                    {inviteState == 0 ? <Button type={"submit"} >Invite</Button> : <Button type={"submit"} >Cancel</Button>}
                </div>
            </section>
        </form>
    );
};
