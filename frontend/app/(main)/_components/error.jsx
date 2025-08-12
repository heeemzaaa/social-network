import React from 'react'

export default function Error({ error }) {
    return (
        <div style={{ margin: "auto", padding: "2rem" }} className='flex-col align-center gap-1'>
            <img src="/error.svg" alt="" style={{width:"100%", maxWidth:"500px"}} />
            <p className="text-danger text-xl font-semibold" aria-live="assertive">{error}</p>
        </div>
    )
}
