import Image from 'next/image'
import React from 'react'

export default function Avatar({path}) {
  return (
    <div className='avatar'>
        <Image src={path}></Image>
    </div>
  )
}
