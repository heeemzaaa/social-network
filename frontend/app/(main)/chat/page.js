'use client'

import React, { use } from 'react'
import "./chat.css"
import { useState } from 'react'
import Button from '@/app/_components/button'
import { HiMiniFaceSmile, HiPaperAirplane } from "react-icons/hi2";
import { RiSendPlaneFill } from "react-icons/ri";
import UserList from '../_components/chat/user_list'
import GroupList from '../_components/group_list'



export default function Chat() {
  const users = {
    users: [
      { username: "Hamza", img: "" },
      { username: "Sara", img: "" },
      { username: "Youssef", img: "" },
      { username: "Lina", img: "" },
      { username: "Mehdi", img: "" },
      { username: "Salma", img: "" },
      { username: "Omar", img: "" },
    ]
  }

  const groups = {
    groups: [
      { name: "grp1", img: "" },
      { name: "grp2", img: "" },
      { name: "grp3", img: "" },
      { name: "grp4", img: "" },
      { name: "grp5", img: "" },
      { name: "grp6", img: "" },
      { name: "grp7", img: "" },
    ]
  }

  const [view, setView] = useState('Users')

  return (
    <main className='chat_main_container p4 flex-row'>

      <section className='user_groups_place h-full flex-col'>
        <div className='user_groups_choosing flex-row justify-center align-center'>
          <Button
            onClick={() => setView('Users')}
            variant={view === 'Users' ? 'btn-primary' : 'btn-secondary'}
          >
            Users
          </Button>

          <Button onClick={() => setView('Groups')}
            variant={view === 'Groups' ? 'btn-primary' : 'btn-secondary'}
            className='p4'
          >
            Groups
          </Button>
        </div>

        <div className='chosing_param'>
          {view === 'Users' ? <UserList {...users} /> : <GroupList {...groups} />}
        </div>
      </section>

      <section className='chat_place flex-col'>
        <div className='chat_header p2'>
          <img src='/no-profile.png'></img>
          <p className='text-lg font-semibold'>Hamza</p>
        </div>
        <div className='chat_body'></div>
        <div className='chat_footer p2'>
          <HiMiniFaceSmile size={'30px'} />
          <textarea></textarea>
          <HiPaperAirplane className='HiPaperAirplane' size={'30px'} />
        </div>
      </section>

    </main>
  )
}
