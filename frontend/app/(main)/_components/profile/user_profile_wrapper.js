'use client'

import { useEffect, useState } from "react"
import InfosDiv from "./user_info"
import AboutUser from "./about_user"
import { FaUserEdit, FaLockOpen, FaLock } from "react-icons/fa"
import { RiUserFollowFill, RiUserUnfollowFill } from "react-icons/ri"
import Button from "@/app/_components/button"
import { usePathname } from "next/navigation"
import "./profile.css"



export default function UserProfileWrapper({ id }) {
  const [userInfos, setUserInfos] = useState(null)
  const [loading, setLoading] = useState(true)
  const [mine, setMine] = useState(false)
  const [privacy, setPrivacy] = useState(null)
  const [follows, setFollows] = useState(null)
  const pathname = usePathname()

  useEffect(() => {
    async function fetchUserInfos() {
      try {
        const res = await fetch(`http://localhost:8080/api/profile/${id}`, { credentials: 'include' })
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
          aboutMe: profile.user.about_me || null,
          isMyProfile: profile.is_my_profile,
          isFollower: profile.is_follower,
          visibility: profile.visibility
        }
        // setID(info.id)
        setUserInfos(info)
        setPrivacy(info.visibility)
        setFollows(info.isFollower)
        setMine(info.isMyProfile)

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
      ? `http://localhost:8080/api/profile/${id}/unfollow`
      : `http://localhost:8080/api/profile/${id}/follow`

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
        if (follows) {
          setUserInfos(prev => ({ ...prev, followers: prev.followers - 1 }))
        } else {
          setUserInfos(prev => ({ ...prev, followers: prev.followers + 1 }))

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
      const res = await fetch(`http://localhost:8080/api/profile/${id}/update-privacy`, {
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
    <>
      <InfosDiv userInfos={userInfos}>
        <section className="buttons flex gap-1" style={{marginLeft: 'auto'}}>
          {mine ? (
            <Button variant={'btn-icon glass-bg gap-1'}>
              <>
                <FaUserEdit size={'24px'} color="white" />
                <span style={{ color: 'white' }}>Edit Profile</span>
              </>
            </Button>
          ) : (
            <Button variant={'btn-icon glass-bg gap-1'} onClick={handleToggleFollow}>
              {follows ? (
                <>
                  <RiUserUnfollowFill size={'24px'} color="white" />
                  <span style={{ color: 'white' }}>Unfollow</span>
                </>
              ) : (
                <>
                  <RiUserFollowFill size={'24px'} color="white" />
                  <span style={{ color: 'white' }}>Follow</span>
                </>
              )}
            </Button>
          )}

          {mine && (
            <Button variant={'btn-icon glass-bg gap-1'} onClick={handleTogglePrivacy}>
              {privacy === 'private' ? (
                <>
                  <FaLockOpen size={'24px'} color="white" />
                  <span style={{ color: 'white' }}>to Public</span>
                </>
              ) : (
                <>
                  <FaLock size={'24px'} color="white" />
                  <span style={{ color: 'white' }}>to Private</span>
                </>
              )}
            </Button>
          )}
        </section>
      </InfosDiv>
      {pathname.startsWith("/profile/") && <AboutUser aboutMe={userInfos.aboutMe} />}
    </>
  )
}
