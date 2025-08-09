'use client'

import "./_components/profileData/profile.css"
import Button from "@/app/_components/button"
import InfosDiv from "./_components/profileData/userInfo"
import AboutUser from "./_components/profileData/abouUser"
import UserPosts from "./_components/profilePosts/userPosts"
import { MdPending } from "react-icons/md"
import { FaLockOpen, FaLock } from "react-icons/fa"
import React, { useEffect, useState } from "react"
import { RiUserFollowFill, RiUserUnfollowFill } from "react-icons/ri"

export default function Page({ params }) {
  const [userInfos, setUserInfos] = useState(null)
  const [loading, setLoading] = useState(true)
  const [isFollower, setIsFollower] = useState(null)
  const resolvedParams = React.use(params);
  const id = resolvedParams.id;

  useEffect(() => {
    async function fetchUserInfo() {
      try {
        const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/profile/${id}/info`, { credentials: 'include' })
        const profile = await res.json()
        const user = profile.user

        setUserInfos({
          id: user.id,
          firstName: user.firstname,
          lastName: user.lastname,
          email: user.email,
          dateOfBirth: user.birthdate,
          nickname: user.nickname || null,
          img: user.avatar || null,
          followers: profile.followers_count || 0,
          following: profile.following_count || 0,
          posts: profile.posts_count || 0,
          groups: profile.groups_count || 0,
          aboutMe: user.about_me,
          isMyProfile: profile.is_my_profile || false,
          isFollower: profile.is_follower || false,
          isRequested: profile.is_requested || false,
          visibility: user.visibility,
          access: profile.access || false,
        })
        setIsFollower(profile.is_follower)
      } catch (err) {
        console.error("Error fetching user profile:", err)
      } finally {
        setLoading(false)
      }
    }
    fetchUserInfo()
  }, [id, isFollower])

  async function handleToggleFollow() {
    let endpoint = ""

    if (userInfos.isRequested) {
      endpoint = `http://localhost:8080/api/profile/${id}/actions/cancel`
    } else if (userInfos.isFollower) {
      endpoint = `http://localhost:8080/api/profile/${id}/actions/unfollow`
    } else {
      endpoint = `http://localhost:8080/api/profile/${id}/actions/follow`
    }

    try {
      const res = await fetch(endpoint, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include',
        body: JSON.stringify({ profile_id: id }),
      })

      if (!res.ok) return console.error("Follow/unfollow/cancel failed")

      const updated = await res.json()

      setUserInfos(prev => ({
        ...prev,
        isFollower: updated.is_follower || false,
        isRequested: updated.is_requested || false,
        access: updated.access || false,
        visibility: updated.user.visibility,
        followers:
          prev.isFollower && !updated.is_follower ? prev.followers - 1 : !prev.isFollower && updated.is_follower
            ? prev.followers + 1
            : prev.followers,
      }))

      if (updated.is_follower) setIsFollower(updated.is_follower)
    } catch (err) {
      console.error("Error:", err)
    }
  }


  async function handleTogglePrivacy() {
    const newPrivacy = userInfos.visibility === 'private' ? 'public' : 'private'
    try {
      console.log("endpoint requested is ", `${process.env.NEXT_PUBLIC_API_URL}/api/profile/${id}/edit/update-privacy`);
      const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/profile/${id}/edit/update-privacy`, {
        method: 'PATCH',
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include',
        body: JSON.stringify({
          profile_id: id,
          wanted_status: newPrivacy,
        }),
      })

      if (!res.ok) return console.error("Failed to update privacy")
      const profile = await res.json()
      setUserInfos(prev => ({
        ...prev,
        visibility: newPrivacy,
        followers: profile.followers_count || prev.followers
      }))
    } catch (err) {
      console.error("Error:", err)
    }
  }

  if (loading) return <p>Loading user info...</p>
  if (!userInfos) return <p>Failed to load user info.</p>

  return (
    <main className='profile_page_section flex h-full p4 gap-4'>
      <InfosDiv userInfos={userInfos}>
        <section className="buttons flex gap-1" style={{ marginLeft: 'auto' }}>
          {!userInfos.isMyProfile && (
            <Button variant="btn-primary glass-bg gap-1" onClick={() => handleToggleFollow()}>
              {userInfos.isRequested ? (
                <>
                  <MdPending size="24px" color="white" />
                  <span style={{ color: 'white' }}>Pending</span>
                </>
              ) : userInfos.isFollower ? (
                <>
                  <RiUserUnfollowFill size="24px" color="white" />
                  <span style={{ color: 'white' }}>Unfollow</span>
                </>
              ) : (
                <>
                  <RiUserFollowFill size="24px" color="white" />
                  <span style={{ color: 'white' }}>Follow</span>
                </>
              )}
            </Button>
          )}

          {userInfos.isMyProfile && (
            <Button variant="btn-icon glass-bg gap-1" onClick={handleTogglePrivacy}>
              {userInfos.visibility === 'private' ? (
                <>
                  <FaLock size="24px" color="white" />
                  <span style={{ color: 'white' }}>Private</span>
                </>
              ) : (
                <>
                  <FaLockOpen size="24px" color="white" />
                  <span style={{ color: 'white' }}>Public</span>
                </>
              )}
            </Button>
          )}
        </section>
      </InfosDiv>

      <div className="data-container flex-col w-full align-center gap-4">
        {(userInfos.access && userInfos.aboutMe) && <AboutUser aboutMe={userInfos.aboutMe} />}
        <UserPosts id={userInfos.id} access={userInfos.access} />
      </div>
    </main>
  )
}
