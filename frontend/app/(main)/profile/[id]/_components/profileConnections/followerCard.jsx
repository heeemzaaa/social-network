import React from 'react'
import "./followers.css"
import { useRouter } from 'next/navigation'
import { useModal } from '@/app/(main)/_context/ModalContext'


export default function FollowerCard({
    id,
    avatar,
    fullname,
    nickname
}) {
    const { closeModal } = useModal()
    const router = useRouter()
    
    function goToProfile(id) {
        closeModal()
        router.push(`/profile/${id}`)
    }

    return (
        <section className='user_card p2 flex justify-start rounded-lg shadow-md m1' onClick={() => goToProfile(id)}>
            <div className='name_image flex align-center gap-1'>
                <img className='no_image' src={avatar ? `http://localhost:8080/static/${avatar}` : "/no-profile.png"} />
                <div className='flex-col justify-start'>
                    <p className='name text-md font-semibold'>{fullname}</p>
                    <p className='nickname text-sm' style={{opacity: '.8'}}>{nickname && "@" + nickname}</p>
                </div>
            </div>
        </section>
    )
}
