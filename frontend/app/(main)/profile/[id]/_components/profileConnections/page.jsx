'use client'
import Button from '../../_components/button'
import {useModal} from '../_context/ModalContext'
import FollowersContainer from './followersContainer'

export default function Followers() {
    const {openModal} = useModal()
  return (
    <Button onClick={() => openModal(<FollowersContainer />)}>Followers</Button>
  )
}
