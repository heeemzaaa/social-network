import { useEffect, useState } from 'react'
import Button from '../../_components/button'
import NotificationsPopover from './notifications/NotificationsContainer'
import {
  HiBell,
} from "react-icons/hi2";
import { useModal } from '../_context/ModalContext';
import { useUserContext } from '../_context/userContext';

export default function Header() {
  const { openModal } = useModal()
  const { authenticatedUser, hasNewNotification, setHasNewNotification } = useUserContext()

  // Fetch notification seen status
  useEffect(() => {
    const LoadPosts = async () => {
      const getRequest = {
        method: "GET",
        credentials: "include"
      }
      try {
        let res = await fetch("http://localhost:8080/api/notifications/", getRequest)
        let response = await res.json()
        console.log("fetch is has seen api, response = ", response)
        if (response?.Status === true) {
          setHasNewNotification(true)
        }
      } catch (err) {
        console.error("Failed to fetch notifications", err)
      }
    }

    LoadPosts()
  }, [])

  return (
    <header className='p3 flex justify-between align-center'>
      <div>
        <h2>
          {authenticatedUser && `Welcome ${authenticatedUser.fullName}!`}
        </h2>
      </div>

      <Button variant='btn-icon' className='flex gap-2 ' onClick={()=> openModal(<NotificationsPopover />)}> {/**/}
        <div className='relative' style={{height:"24px"}} >
              <HiBell size={24} />
              {hasNewNotification && (
                <span className="notification-badge"></span>
              )}
        </div>
      </Button>
    </header>
  )
}
