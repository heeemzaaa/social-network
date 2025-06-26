import Image from 'next/image'
import React from 'react'

export default function Avatar() {
  return (
    <div className='avatar'>
        <Image src={`avatar.png`}></Image>
    </div>
  )
}
