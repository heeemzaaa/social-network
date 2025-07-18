"use client"

import { useModal } from '../_context/ModalContext';
import Avatar from '../_components/avatar';
import PostCardList from './postCardList';

export default function Page() {
    const { openModal } = useModal()
    return (
        <main className=' '>
            <h2>This is a page to test components :</h2>
            {/* <div className="bts flex gap-1">
                <Button onClick={() => openModal(<AddGroupForm />)} > Add Post </Button>
                <Button onClick={() => openModal(<AddEventForm />)} > Add Event </Button>
            </div> */}
            {/* <Tabs>
                <Tab label="Tab 1" />
                <Tab label="Tab 2" />
                <TabContent>Content for Tab 1</TabContent>
                <TabContent>Content for Tab 2</TabContent>
            </Tabs> */}
            {/* <Avatar size={42}/> */}
            <PostCardList/>
        </main>
    )
}
