import React from 'react'

export default function Logo() {

    return (
        <div className='logo flex-col justify-center align-center'>
            <img src='/social_network_logo.svg' style={{ width: '80px', height: '80px'}} />
            <h2 style={{color: 'var(--color-primary)', fontSize: 'var(--font-size-2xl)', fontWeight: 'var(--font-weight-medium)'}}>EmiTalk</h2>
        </div>
    )
}
