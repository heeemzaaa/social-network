import Button from '@/app/_components/button'
import React from 'react'

export default function groupEventCard({
    going,
    title,
    event_id,
    event_date,
    description,
    evetnt_creator,
}) {
    return (
        <div className="event-card">
            <div>
                <div></div>
            </div>
            <div>
                <Button>Going</Button>
                <Button>Not Going</Button>
            </div>
        </div>
    )
}
