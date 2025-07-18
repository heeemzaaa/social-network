'use client'

import React, { useEffect, useState } from "react"
import InfosDiv from "./_components/userInfo"
import AboutUser from "./_components/abouUser"
import UserPosts from "./_components/userPosts"
import { FaLockOpen, FaLock } from "react-icons/fa"
import { RiUserFollowFill, RiUserUnfollowFill } from "react-icons/ri"
import { MdPending } from "react-icons/md";
import Button from "@/app/_components/button"
import "./_components/profile.css"



export default function Page({ params }) {
  const [userInfos, setUserInfos] = useState(null)
  const [loading, setLoading] = useState(true)
  const [mine, setMine] = useState(false)
  const [privacy, setPrivacy] = useState(null)
  const [follows, setFollows] = useState(null)
  const [requested, setRequested] = useState(null)
  const [access, setAccess] = useState(null)


  const resolvedParams = React.use(params)
  const id = resolvedParams.id


  useEffect(() => {
    async function fetchUserInfos() {
      try {
        const res = await fetch(`http://localhost:8080/api/profile/${id}/info`, { credentials: 'include' })
        const profile = await res.json()

        const info = {
          id: profile.user.id,
          firstName: profile.user.firstname,
          lastName: profile.user.lastname,
          email: profile.user.email,
          dateOfBirth: profile.user.birthdate,
          nickname: profile.user.nickname || null,
          img: profile.user.avatar || null,
          followers: profile.followers_count || 0,
          following: profile.following_count || 0,
          posts: profile.posts_count || 0,
          groups: profile.groups_count || 0,
          aboutMe: profile.user.about_me,
          isMyProfile: profile.is_my_profile,
          isFollower: profile.is_follower,
          isRequested: profile.is_requested,
          visibility: profile.visibility,
          access: profile.access
        }
        setUserInfos(info)
        setPrivacy(info.visibility)
        setFollows(info.isFollower)
        setMine(info.isMyProfile)
        setRequested(info.isRequested)
        setAccess(info.access)
        


      } catch (err) {
        console.error("Error fetching user profile:", err)
      } finally {
        setLoading(false)
      }
    }

    fetchUserInfos()
  }, [id])

  async function handleToggleFollow() {
    const endpoint = follows
      ? `http://localhost:8080/api/profile/${id}/actions/unfollow`
      : `http://localhost:8080/api/profile/${id}/actions/follow`

    try {
      const res = await fetch(endpoint, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        credentials: 'include',
        body: JSON.stringify({ profile_id: id }),
      })

      if (res.ok) {
        setFollows(!follows)
        if (privacy == 'public') {
          if (follows) {
            setUserInfos(prev => ({ ...prev, followers: prev.followers - 1 }))
          } else {
            setUserInfos(prev => ({ ...prev, followers: prev.followers + 1 }))
          }
        } else {
          if (follows) {
            setAccess(false)
            setUserInfos(prev => ({ ...prev, email: "" }))
            setUserInfos(prev => ({ ...prev, dateOfBirth: "" }))
            setUserInfos(prev => ({ ...prev, aboutMe: "" }))
            setUserInfos(prev => ({ ...prev, followers: prev.followers - 1 }))
          } else {
            let isRequestedNow = !requested
            setRequested(isRequestedNow)
            setUserInfos(prev => ({ ...prev, isRequested: isRequestedNow }))
          }
        }

      } else {
        console.error("Failed to follow/unfollow")
      }
    } catch (err) {
      console.error("Error:", err)
    }
  }


  async function handleTogglePrivacy() {
    const newPrivacy = privacy === 'private' ? 'public' : 'private'

    try {
      const res = await fetch(`http://localhost:8080/api/profile/${id}/edit/update-privacy`, {
        method: 'PATCH',
        headers: {
          'Content-Type': 'application/json',
        },
        credentials: 'include',
        body: JSON.stringify({
          profile_id: id,
          wanted_status: newPrivacy,
        }),
      })

      if (res.ok) {
        setPrivacy(newPrivacy)
        setUserInfos(prev => ({ ...prev, privacy: newPrivacy }))
      } else {
        console.error("Failed to update privacy status")
      }
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
          {!mine && (
            <Button variant={'btn-primary glass-bg gap-1'} onClick={handleToggleFollow} disabled={requested}>
              {requested ? (
                <>
                  <MdPending size="24px" color="white" />
                  <span style={{ color: 'white' }}>Pending</span>
                </>
              ) : follows ? (
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

          {mine && (
            <Button variant={'btn-icon glass-bg gap-1'} onClick={handleTogglePrivacy}>
              {privacy === 'private' ? (
                <>
                  <FaLock size={'24px'} color="white" />
                  <span style={{ color: 'white' }}>Private</span>
                </>
              ) : (
                <>
                  <FaLockOpen size={'24px'} color="white" />
                  <span style={{ color: 'white' }}>Public</span>
                </>
              )}
            </Button>
          )}
        </section>
      </InfosDiv>
      <div className="data-container scrollable-section flex-col w-full align-center gap-4 " >
        {userInfos.aboutMe && <AboutUser aboutMe={userInfos.aboutMe} />}
        <div className="flex-col align-center h-full w-full p2">
          <UserPosts id={userInfos.id} />
        </div>
      </div>
    </main>
  )
}
