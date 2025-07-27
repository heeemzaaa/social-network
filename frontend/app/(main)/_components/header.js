import { useEffect, useState } from 'react'
import Button from '../../_components/button'
import NotificationsPopover from './notifications/NotificationsPopover'
import {
  HiBell,
  HiChatBubbleOvalLeftEllipsis,
  HiMiniPlusCircle,
  HiMiniPlusSmall
} from "react-icons/hi2";
import Popover from './popover';
import { useModal } from '../_context/ModalContext';
import CreatePost from './posts/createPost';
import { createPostAction } from '@/app/_actions/posts';
import { useUserContext } from '../_context/userContext';

export default function Header() {
  const {authenticatedUser} = useUserContext()
  console.log('authenticatedUser', authenticatedUser)
  const { openModal } = useModal()
  const [hasNewNotification, setHasNewNotification] = useState(false)

  // Fetch notification seen status
  useEffect(() => {
    const LoadPosts = async () => {
      const getRequest = {
        method: "GET",
        credentials: "include"
      }
      try {
        let res = await fetch("http://localhost:8080/api/notifications/", getRequest)
        let ddd = await res.json()
        if (ddd?.Status === true) {
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
        <h2>Welcome user!!</h2>
      </div>

      <div className='flex gap-2'>
        <Popover trigger={<HiMiniPlusCircle size={24} />}>
          <Button className={"w-full"} variant='btn-tertiary' onClick={() => openModal(<CreatePost postAction={createPostAction} />)}>
            <HiMiniPlusSmall size={"30px"} />
            <span>Add post</span>
          </Button>
          <Button variant='btn-tertiary' onClick={() => openModal(<CreateGroupForm />)}>
            <HiMiniPlusSmall size={"30px"} />
            <span>Add Group</span>
          </Button>
        </Popover>

        <Button variant='btn-icon'>
          <HiChatBubbleOvalLeftEllipsis size={24} />
        </Button>

        <Popover 
          trigger={
            <div className="relative">
              <HiBell size={24} />
              {hasNewNotification && (
                <span className="notification-badge"></span>
              )}
            </div>
          }
        >
          <NotificationsPopover />
        </Popover>
      </div>
    </header>
  )
}
