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
import { HiOutlineDocumentPlus } from "react-icons/hi2"
import { useModal } from "../../_context/ModalContext"
import CreatePost from "../../_components/posts/createPost"
import { createPostAction } from "@/app/_actions/posts"
import Loader from "../../_components/loader"
import { useRouter } from "next/navigation"
import Error from "../../_components/error"
import { useNotification } from "../../_context/NotificationContext"

export default function Page({ params }) {
  const [userInfos, setUserInfos] = useState(null)
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState(null)
  const [isFollower, setIsFollower] = useState(null)
  const [changed, setChanged] = useState(false)
  const { openModal } = useModal()
  const { showNotification } = useNotification()
  const router = useRouter()

  const resolvedParams = React.use(params);
  const id = resolvedParams.id;

  useEffect(() => {
    async function fetchUserInfo() {
      try {
        const res = await fetch(`http://localhost:8080/api/profile/${id}/info`, { credentials: 'include' })

        const profile = await res.json()

        if (!res.ok) {
          if (profile.status === 401) {
            router.push("/login")
            return
          }

          if (profile.status === 400 || profile.status === 404 || profile.status === 500) {
            setError(profile.error)
            return
          }
        }

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
        setError(err)
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

      const updated = await res.json()

      if (!res.ok) {
        if (updated.status === 401) {
          router.push("/login")
          return
        }
        if (updated.status === 400 || updated.status === 404 || updated.status === 500) {
          setError(updated.error)
          return
        }
      }

      setChanged(!changed)
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
      
      if (!updated.is_follower && !updated.is_requested && isFollower) showNotification({ Content: 'You Unfollowed this profile successfully !', Status: 'success' })
      setIsFollower(updated.is_follower)
    } catch (err) {
      setError(err)
      console.error("Error:", err)
    }
  }


  async function handleTogglePrivacy() {
    const newPrivacy = userInfos.visibility === 'private' ? 'public' : 'private'
    try {
      const res = await fetch(`http://localhost:8080/api/profile/${id}/edit/update-privacy`, {
        method: 'PATCH',
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include',
        body: JSON.stringify({
          wanted_status: newPrivacy,
        }),
      })


      const profile = await res.json()

      if (!res.ok) {
        if (profile.status === 401) {
          router.push("/login")
          return
        }
        if (profile.status === 400 || profile.status === 404 || profile.status === 500) {
          setError(profile.error)
          return
        }
      }
      
      setUserInfos(prev => ({
        ...prev,
        visibility: newPrivacy,
        followers: profile.followers_count || prev.followers
      }))

      showNotification({Content: `Your profile is now ${newPrivacy}`, Status: 'success'})
    } catch (err) {
      setError(err)
      console.error("Error:", err)
    }
  }

  if (loading) return (
    <main className='profile_page_section flex h-full p4 gap-4'>
      <Loader />
    </main>
  )

  if (!userInfos || error) return (
    <main className='profile_page_section  h-full p4 gap-4'>
      <Error error={error} />
    </main>
  )

  return (
    <main className='profile_page_section flex h-full p4 gap-4'>
      <InfosDiv userInfos={userInfos}>
        <section className="buttons flex gap-1" style={{ marginLeft: 'auto' }}>
          {!userInfos.isMyProfile && (
            <Button variant="btn-primary glass-bg gap-1" onClick={() => handleToggleFollow()}>
              {userInfos.isRequested ? (
                <>
                  <MdPending size="24px" color="white" />
                  <span style={{ color: 'white', fontSize: 'var(--font-size-lg)', fontWeight: 'var(--font-weight-medium)' }}>Pending</span>
                </>
              ) : userInfos.isFollower ? (
                <>
                  <RiUserUnfollowFill size="24px" color="white" />
                  <span style={{ color: 'white', fontSize: 'var(--font-size-lg)', fontWeight: 'var(--font-weight-medium)' }}>Unfollow</span>
                </>
              ) : (
                <>
                  <RiUserFollowFill size="24px" color="white" />
                  <span style={{ color: 'white', fontSize: 'var(--font-size-lg)', fontWeight: 'var(--font-weight-medium)' }}>Follow</span>
                </>
              )}
            </Button>
          )}

          {userInfos.isMyProfile && (
            <Button variant="btn-icon privacy glass-bg gap-1" onClick={handleTogglePrivacy} style={{ backgroundColor: userInfos.visibility === 'private' ? 'var(--color-red)' : 'var(--color-green)' }}
            >
              {userInfos.visibility === 'private' ? (
                <>
                  <FaLock size="20px" color="white" />
                  <span style={{ color: 'white' }}>Private</span>
                </>
              ) : (
                <>
                  <FaLockOpen size="20px" color="white" />
                  <span style={{ color: 'white' }}>Public</span>
                </>
              )}
            </Button>
          )}
        </section>
      </InfosDiv>

      <div className="data-container flex-col w-full align-center gap-4">
        {(userInfos.access && userInfos.aboutMe) && <AboutUser aboutMe={userInfos.aboutMe} />}
        <div style={{ zIndex: "1", background: "var(--color-secondary)", alignSelf: "stretch", position: "sticky", top: "0", borderBottom: "solid 1px", paddingBottom: ".5rem", marginBlock: ".5rem" }} >
          <Button style={{ marginLeft: "auto" }} onClick={() => openModal(<CreatePost postAction={createPostAction} />)}>
            <HiOutlineDocumentPlus size={24} />
            <span className="text-lg font-medium">Add New Post</span>
          </Button>
        </div>
        {<UserPosts profileId={userInfos.id} access={userInfos.access} changed={changed} />}
      </div>
    </main>
  )
}
