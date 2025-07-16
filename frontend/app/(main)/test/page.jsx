"use client"

import Button from '@/app/_components/button';
import AddGroupForm from '../groups/_components/createGroupForm';
import AddEventForm from '../groups/_components/addEventForm';
import { useModal } from '../_context/ModalContext';
import Tabs from '../_components/tab/tabs';
import Tab from '../_components/tab/tab';
import TabContent from '../_components/tab/tabContent';

export default function Page() {
    const { openModal } = useModal()
    return (
        <main className=' '>
            <h2>This is a page to test components :</h2>
            <div className="bts flex gap-1">
                <Button onClick={() => openModal(<AddGroupForm />)} > Add Post </Button>
                <Button onClick={() => openModal(<AddEventForm />)} > Add Event </Button>
            </div>
            <Tabs>
                <Tab label="Tab 1" />
                <Tab label="Tab 2" />
                <TabContent>Content for Tab 1</TabContent>
                <TabContent>Content for Tab 2</TabContent>
            </Tabs>
        </main>
    )
}
