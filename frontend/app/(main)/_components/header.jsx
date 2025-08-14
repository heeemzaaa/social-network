import { useEffect, useState } from 'react'
import Button from '../../_components/button'
import NotificationsPopover from './notifications/NotificationsContainer'
import {
  HiBell,
} from "react-icons/hi2";
import { useModal } from '../_context/ModalContext';
import { useUserContext } from '../_context/userContext';
const API_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080'

import { useNotification } from '../_context/NotificationContext';

export default function Header() {
  const { openModal } = useModal()
  const { authenticatedUser, hasNewNotification, setHasNewNotification } = useUserContext()
  const { showNotification } = useNotification()


  // Fetch notification seen status
  useEffect(() => {
    const LoadPosts = async () => {
      const getRequest = {
        method: "GET",
        credentials: "include"
      }
      
      try {
        let res = await fetch(`${API_URL}/api/notifications/`, getRequest)
        let response = await res.json()
        if (response?.Status === true) {
          // setHasNewNotification(true)
          showNotification({ Content: "You have new notifications", Status: "info" })
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
        <h2 className='font-semibold text-2xl '>
          {authenticatedUser && `Welcome ${authenticatedUser.fullName}!`}
        </h2>
      </div>

      <Button variant='btn-icon' className='flex gap-2 ' onClick={()=> openModal(<NotificationsPopover />)}>
        <div className='relative' style={{height:"24px"}} >
              <HiBell size={24} />
              {/* {hasNewNotification && (
                <span className="notification-badge"></span>
              )} */}
        </div>
      </Button>
    </header>
  )
}