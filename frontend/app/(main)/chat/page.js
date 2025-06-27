'use client'

import React, { use } from 'react'
import "./chat.css"
import { useState } from 'react'
import Button from '@/app/_components/button'
import { HiMiniFaceSmile, HiPaperAirplane } from "react-icons/hi2";
import { RiSendPlaneFill } from "react-icons/ri";



export default function Chat() {

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
      </section>

      <section className='chat_place flex-col'>
        <div className='chat_header p2'>
          <img src='/no-profile.png'></img>
          <p className='text-lg font-semibold'>Hamza</p>
          </div>
        <div className='chat_body'></div>
        <div className='chat_footer p2'>
          <HiMiniFaceSmile  size={'30px'}/>
          <textarea></textarea>
          <HiPaperAirplane className='HiPaperAirplane' size={'30px'}/>
        </div>
      </section>

    </main>
  )
}
