import React from 'react'
import "./chat.css"

export default function Chat() {
  return (
    <main className='chat_main_container p4'>

        <section className='user_groups_place h-full flex-col'>
            <div className='user_groups_choosing flex-row justify-center align-center'>
                  <div className='users_case p2' ><p>Users</p></div>
                  <div className='groups_case p2'><p>Groups</p></div>
            </div>
        </section>

        <section className='chat_place'>
            <div className='chat_header'></div>
            <div className='chat_body'></div>
            <div className='chat_footer'></div>
        </section>

    </main>
  )
}
