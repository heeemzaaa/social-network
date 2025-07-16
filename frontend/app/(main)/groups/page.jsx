'use client'

// group page

import Button from '@/app/_components/button'
import Tabs from '../_components/tab/tabs';
import Tab from '../_components/tab/tab';
import TabContent from '../_components/tab/tabContent';
import { useModal } from '../_context/ModalContext';
import GroupCardList from './_components/groupCardList';
import { HiMiniUserGroup } from 'react-icons/hi2';
import AddGroupForm from './_components/createGroupForm';
import { useState } from 'react';
import CreateGroupForm from './_components/createGroupForm';


let tabs = [
  {
    label: "Your Groups",
    type: "owned",
  },
  {
    label: "Joined Groups",
    type: "joined",
  },
  {
    label: "Available Groups",
    type: "available",
  }
]


export default function Groups() {
  let [activeTab, setActiveTab] = useState("owned")


  const { openModal } = useModal()
  return (
    <main className='flex-col flex-start border-red align-end'>
      <Button className={'justify-start'} onClick={() => { openModal(<CreateGroupForm />) }}>
        <HiMiniUserGroup size={"24px"} />
        <span>Create New Group</span>
      </Button>
      <Tabs className={''}>
        <Tab label="Your Groups" />
        <Tab label="Joined Groups" />
        <Tab label="Groups" />
        <TabContent>
          <GroupCardList key={"owned"} filter={"owned"} />
        </TabContent>
        <TabContent>
          <GroupCardList key={"joined"} filter={"joined"} />
        </TabContent>
        <TabContent>
          <GroupCardList key={"available"} filter={"available"} />
        </TabContent>
      </Tabs> 
    </main>
  )
}


