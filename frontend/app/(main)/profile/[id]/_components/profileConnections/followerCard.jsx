import React from 'react'
import "./followers.css"
import { useRouter } from 'next/navigation'

export default function FollowerCard({
    id,
    avatar,
    firstname,
    lastname,
    nickname
}) {
    const router = useRouter()
    function goToProfile(id) {
        router.push(`/profile/${id}`)
    }

    return (
        <section className='user_card p2 flex justify-start rounded-lg shadow-md m1' onClick={()=> goToProfile(id)}>
            <div className='name_image flex align-center gap-1'>
                <img className='no_image' src={avatar ? `http://localhost:8080/static/${avatar}` : "/no-profile.png"} />
                <div className='flex-col justify-start'>
                    <p className='name'>{firstname + " " + lastname}</p>
                    <p className='nickname'>{nickname && "@" + nickname}</p>
                </div>
            </div>
        </section>
    )
}
