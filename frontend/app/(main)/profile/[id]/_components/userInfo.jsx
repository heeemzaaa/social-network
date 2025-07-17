import { useState } from "react"
import "./profile.css"

export default function InfosDiv({ userInfos, children }) {

  const [followersList, setFollowersList] = useState([])
  const [followingList, setFollowingList] = useState([])

  async function handleGetFollowers() {
    try {
      const res = await fetch(`http://localhost:8080/api/profile/${userInfos.id}/connections/followers`, {
        credentials: "include",
      })
      if (res.ok) {
        const followers = await res.json()
        setFollowersList(followers)

        {
          followersList.length > 0 && (
            <div className="followers-list">
              <h3>Followers</h3>
              <ul>
                {followersList.map((user) => (
                  <li key={user.id}>{user.username}</li>
                ))}
              </ul>
            </div>
          )
        }

      }
    } catch (err) {
      console.error("Failed to fetch followers", err)
    }
  }

  async function handleGetFollowing() {
    try {
      const res = await fetch(`http://localhost:8080/api/profile/${userInfos.id}/connections/following`, {
        credentials: "include",
      })
      if (res.ok) {
        const following = await res.json()
        console.log('following', following)
        setFollowingList(following)
        {
          followingList.length > 0 && (
            <div className="followers-list">
              <h3>Following</h3>
              <ul>
                {followingList.map((user) => (
                  <li key={user.id}>{user.username}</li>
                ))}
              </ul>
            </div>
          )
        }

      }
    } catch (err) {
      console.error("Failed to fetch following", err)
    }
  }


  return (
    <section className="profileLeftSection h-full">
      <div
        className="ProfileContainer p2"
        style={{ backgroundImage: `url(${'/no-profile.png'})` }}
      >
        {children}
        <div className="ProfileData p2 flex-col gap-1">
          <p><span className="font-bold">First Name:</span> {userInfos.firstName}</p>
          <p><span className="font-bold">Last Name:</span> {userInfos.lastName}</p>
          <p><span className="font-bold">Email:</span> {userInfos.email}</p>
          <p><span className="font-bold">Date of Birth:</span> {userInfos.dateOfBirth}</p>
          {userInfos.nickname && <p><span className="font-bold">Nickname:</span> {userInfos.nickname}</p>}

        </div>
      </div>

      <div className="UserNumbers p2">
        <div className="followers p2" onClick={handleGetFollowers}>
          <p className="font-bold">Followers</p><p>{userInfos.followers}</p>
        </div>
        <div className="following p2" onClick={handleGetFollowing}>
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
