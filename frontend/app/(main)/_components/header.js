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
import { CreateGroup } from '../groups/_components/createGroup';


export default function Header() {
  let { openModal } = useModal()
  return (
    <header className='p3 flex justify-between align-center' >
      <div>
        <h2>
          Welcome User!!
        </h2>
      </div>

      <div className='flex gap-2'>
        <Popover trigger={<HiMiniPlusCircle size={24} />}>
          <Button style={"w-full"} variant='btn-tertiary' onClick={() => openModal(<CreatePost />)}>
            <HiMiniPlusSmall size={"30px"} />
            <span>
              Add post
            </span>
          </Button>
          <Button variant='btn-tertiary' onClick={() => openModal(<CreateGroup />)}>
            <HiMiniPlusSmall size={"30px"} />
            <span>
              Add Group
            </span>
          </Button>
        </Popover>

        <Button variant='btn-icon'>
          <HiChatBubbleOvalLeftEllipsis size={24} />
        </Button>

        <Popover trigger={<HiBell size={24} />}>
          <NotificationsPopover />
        </Popover>
      </div>

    </header>
  )
}

