import Button from '@/app/_components/button'
import React, { useState } from 'react'
import Avatar from '../../_components/avatar'
import { HiCalendarDays } from 'react-icons/hi2'
import Tag from '../../_components/tag'
import { timeAgo } from "@/app/_utils/time"

import { HiOutlineClock } from "react-icons/hi2"

let style = {
    width: "100%",
    maxWidth: "500px",
    height: "max-content"
}

export default function GroupEventCard({
    group,
    event_id,
    event_creator,
    title,
    description,
    event_date,
    created_at,
    going,
}) {
    const [goingState, setGoingState] = useState(going)
    console.log("goingState", going);
    const endpoint = `http://localhost:8080/api/groups/${group.group_id}/events/${event_id}/`
    async function handleGoingState(actionValue) {
        try {
            const res = await fetch(endpoint, {
                method: "POST",
                headers: { 'Content-Type': 'application/json' },
                credentials: 'include',
                body: JSON.stringify({ 'action': actionValue }),
            })


            console.log(JSON.stringify({ 'action': actionValue }));

            if (!res.ok) return console.error("Failed to send the request")
            let newGoingState = await res.json()
            setGoingState(newGoingState.action)
        } catch (err) {
            console.log(err);
        }
    }
    return (
        <div style={style} className="flex-col gap-1  bg-white p2 pi3 rounded-xl shadow-md" key={event_id}>
            <div className='flex align-center gap-2'>
                <Avatar img={event_creator.avatar} size={42} />
                {/* <img src='/no-profile.png'/> */}
                <div>
                    <p className='font-semibold'>{event_creator.fullname}</p>
                    <span className=''>@{event_creator.nickname}</span>
                    {/* <p className=''>{created_at}</p> */}
                </div>
            </div>
            <hr />
            <div className='flex-col gap-1'>
                <h3>{title}</h3>
                <p>{description}</p>
                <Tag className='flex align-end ' style={{ gap: "5px" }}>
                    <HiCalendarDays size={24} />
                    <span className='text-sm'>{event_date}</span>
                </Tag>
            </div>
            <div className='flex justify-end gap-1'>
                {/* {
                    going == 0 ? variant='btn-danger text-white' :variant='btn-tertiary'
                } */}
                <Button variant={goingState != -1 ? goingState === 0 ? 'btn-tertiary' : 'btn-primary' : 'btn-tertiary'} onClick={() => handleGoingState(1)}> Going</Button>
                <Button variant={goingState != 1 ? goingState === 0 ? 'btn-tertiary' : 'btn-danger' : 'btn-tertiary'} onClick={() => handleGoingState(-1)}>Not Going</Button>


            </div>

            <div style={{ opacity: ".5", gap: "5px", paddingLeft: "3px", marginLeft: "auto" }} className="flex align-end">
                <HiOutlineClock size={24} />
                <span>{timeAgo(created_at)}</span>
            </div>
        </div>
    )
}

