import React from 'react'

export default async function Profile({params}) {
    let {id } = await params
  return (
    <div>Profile id {id}</div>
  )
}
