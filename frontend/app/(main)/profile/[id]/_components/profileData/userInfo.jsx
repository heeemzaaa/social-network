import { useState } from "react"
import "./profile.css"
import { useModal } from "@/app/(main)/_context/ModalContext"
import UsersContainer from "../profileConnections/followersContainer"

export default function InfosDiv({ userInfos, children }) {
  const {openModal} = useModal()
   
  return (
    <section className="profileLeftSection h-full">
      <div
        className="ProfileContainer p2"
        style={{ backgroundImage: `url(${userInfos.img ? `http://localhost:8080/static/${userInfos.img}` : '/no-profile.png'})` }}
      >
        {children}
        <div className="ProfileData p2 flex-col gap-1">
          <p><span className="font-bold">First Name:</span> {userInfos.firstName}</p>
          <p><span className="font-bold">Last Name:</span> {userInfos.lastName}</p>
          {userInfos.access && <p><span className="font-bold">Email:</span> {userInfos.email}</p>}
          {userInfos.access && <p><span className="font-bold">Date of Birth:</span> {userInfos.dateOfBirth}</p>}
          {userInfos.nickname && <p><span className="font-bold">Nickname:</span> {userInfos.nickname}</p>}

        </div>
      </div>

      <div className="UserNumbers p2">
        <div className="followers p2" onClick={() => userInfos.followers !== 0 && userInfos.access && openModal(<UsersContainer type={"followers"} userID={userInfos.id} />)}>
          <p className="font-bold">Followers</p><p>{userInfos.followers}</p>
        </div>
        <div className="following p2" onClick={() => userInfos.following !== 0 && userInfos.access && openModal(<UsersContainer type={"following"} userID={userInfos.id} />)}>
          <p className="font-bold">Following</p><p>{userInfos.following}</p>
        </div>
        <div className="posts p2">
          <p className="font-bold">Posts</p><p>{userInfos.posts}</p>
        </div>
        <div className="groups p2">
          <p className="font-bold">Groups</p><p>{userInfos.groups}</p>
        </div>
      </div>
    </section>

  )
}
