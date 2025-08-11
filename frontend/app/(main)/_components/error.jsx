import React from 'react'

export default function Error({ error }) {
    return (
        <div style={{ margin: "auto", padding: "2rem" }} className='flex-col align-center gap-1'>
            <img src="/error.png" alt="" />
            <p className="text-danger font-semibold" aria-live="assertive">{error}</p>
        </div>
    )
}
