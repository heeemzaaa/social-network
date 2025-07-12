'use client'

import Button from '@/app/_components/button'
import { CreateGroup } from "./_components/createGroup"
import Tabs from '../_components/tab/tabs';
import Tab from '../_components/tab/tab';
import TabContent from '../_components/tab/tabContent';
import { useModal } from '../_context/ModalContext';


export default function Groups() {
  const { openModal } = useModal()
  return (
    <main className='flex-col '>
      <Button className={'self-end'} onClick={() => { openModal(<CreateGroup />) }}>Add Post</Button>
      <Tabs className={'flex-grow'}>
        <Tab label="Your Groups" />
        <Tab label="Joined Groups" />
        <Tab label="Groups" />
        <TabContent>Content for Tab 1</TabContent>
        <TabContent>Content for Tab 2</TabContent>
        <TabContent>Content for Tab 3</TabContent>
      </Tabs>
    </main>
  )
}
