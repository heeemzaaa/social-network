import React, { useEffect, useState } from 'react'
import Avatar from '../../_components/avatar';
import Button from '@/app/_components/button';
import { useActionState } from 'react';
import { inviteUserAction, CancelInvitationAction } from '@/app/_actions/group';

export default function UserCard({ user, groupId }) {
    const [inviteActionState, setInviteAction] = useActionState(inviteUserAction, {});
    const [cancelActionState, setCancelAction] = useActionState(CancelInvitationAction, {});
    const [inviteState, setInviteState] = useState(user.invited)
    // console.log("staaaaaaaaaate", inviteState);
    // inviteState =
    //     inviteActionState?.message === "success" ? 1 :
    //         cancelState?.message === "success" ? 0 :
    //             user.invited ? 1 : 0;


    // console.log("invite staaaaaaaaaaate: ", inviteState)


    // Track success and toggle invite state
    useEffect(() => {
        if (inviteActionState?.message === 'done') {
            setInviteState(1); // user is now invited
        } else if (cancelActionState?.message === 'done') {
            setInviteState(0); // invitation cancelled
        }
    }, [inviteActionState, cancelActionState]);

    return (
        <form action={inviteState === 0 ? setInviteAction : setCancelAction}>
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
                        // style={{ marginRight: '12px' }/**/}
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
                    {inviteState == 0 ?
                        <Button type={"submit"} onClick={() => setInviteState(1)} >Invite</Button> :
                        <Button type={"submit"} onClick={() => setInviteState(0)} >Cancel</Button>}
                </div>
            </section>
        </form>
    );
};
