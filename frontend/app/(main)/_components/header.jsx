import { useEffect, useState } from 'react'
import Button from '../../_components/button'
import NotificationsPopover from './notifications/NotificationsPopover'
import {HiBell} from "react-icons/hi2";

import { useModal } from '../_context/ModalContext';
import { useUserContext } from '../_context/userContext';

export default function Header() {
  const { openModal } = useModal()
  const [hasNewNotification, setHasNewNotification] = useState(false)
  const { authenticatedUser } = useUserContext()

  // Fetch notification seen status
  useEffect(() => {
    const LoadPosts = async () => {
      const getRequest = {
        method: "GET",
        credentials: "include"
      }
      
      try {
        console.log("heere inside the notifications", `${process.env.NEXT_PUBLIC_API_URL}/api/notifications/`);
        let res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/notifications/`, getRequest)
        console.log("response of the notifications", res);
        let response = await res.json()
        console.log("response of notifications", response);
        if (response?.Status === true) {
          setHasNewNotification(true)
        }
      } catch (err) {
        console.log("inside the error", err);
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

      <div className='flex gap-2'>

        <Button variant='btn-icon ' className="relative" onClick={() => openModal(<NotificationsPopover />)}>
          <HiBell size={24} />
          {hasNewNotification && (
            <span className="notification-badge"></span>
          )}
        </Button>


      </div>
    </header>
  )
}
