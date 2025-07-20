'use client'
import React from 'react'
import Button from '../../_components/button'
import FollowersContainer from './followersContainer'
import {useModal} from '../_context/ModalContext'

export default function Followers() {
    const {openModal} = useModal()
  return (
    <Button onClick={() => openModal(<FollowersContainer />)}>Followers</Button>
  )
}
