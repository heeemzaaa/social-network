import Button from '@/app/_components/button'
import React from 'react'
import Avatar from '../../_components/avatar'
import { HiCalendarDays } from 'react-icons/hi2'
import Tag from '../../_components/tag'


let style = {
    width: "100%",
    maxWidth: "500px",
    height: "max-content"
}

export default function GroupEventCard({
    event_id,
    event_creator,
    title,
    description,
    event_date,
    created_at, 
    going,
}) {
    return (
        
        <div style={style} className="flex-col gap-1  bg-white p2 pi3 rounded-xl shadow-md" key={event_id}>
            <div className='flex align-center gap-2'>
                <Avatar size={42} />
                <div>
                    <p className='font-semibold'>{event_creator.fullname}</p>
                    <span className=''>@nickname</span>
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
                <Button variant='btn-tertiary'> Going</Button>
                <Button variant='btn-danger text-white'>Not Going</Button>
            </div>
        </div>
    )
}

