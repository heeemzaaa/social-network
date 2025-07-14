'use client'

import Button from '@/app/_components/button'
import { CreateGroup } from "./_components/createGroup"
import Tabs from '../_components/tab/tabs';
import Tab from '../_components/tab/tab';
import TabContent from '../_components/tab/tabContent';
import { useModal } from '../_context/ModalContext';
import GroupCardList from './_components/groupCardList';


export default function Groups() {
  const { openModal } = useModal()
  return (
    <main className='flex-col flex-start border-red align-end'>
      <Button className={'justify-start'} onClick={() => { openModal(<CreateGroup />) }}>Add Post</Button>
      <Tabs className={''}>
        <Tab label="Your Groups" />
        <Tab label="Joined Groups" />
        <Tab label="Groups" />
        <TabContent><GroupCardList filter={"owned"}/></TabContent>
        <TabContent> 
          <GroupCardList filter={"joined"}/>
        </TabContent>
        <TabContent>
          <GroupCardList filter={"available"}/>
        </TabContent>
      </Tabs>
    </main>
  )
}
