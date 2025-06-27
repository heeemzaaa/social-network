import "./components.css"

export default function InfosDiv({ firstName, lastName, email, dateOfBirth, nickname, followers, following, posts, groups }) {
  return (
    <section className="profileLeftSection h-full">
      <div className="ProfileContainer p2" style={{ backgroundImage: `url(https://www.portraitprofessionnel.fr/wp-content/uploads/2020/02/portrait-professionnel-corporate-4.jpg)` }} >
        <div className="ProfileData p2 flex-col gap-1">
          <p><span className="font-bold">First Name:</span> {firstName}</p>
          <p><span className="font-bold">Last Name:</span> {lastName}</p>
          <p><span className="font-bold">Email:</span> {email}</p>
          <p><span className="font-bold">Date of Birth:</span> {dateOfBirth}</p>
          <p><span className="font-bold">Nickname:</span> {nickname}</p>
        </div>
      </div>

      <div className="UserFollowers p2">
        <div className="followers p2">
          <p className="font-bold">Followers</p><p>{followers}</p>
        </div>

        <div className="following p2">
          <p className="font-bold">Following</p><p>{following}</p>
        </div>

        <div className="posts p2">
          <p className="font-bold">Posts</p><p>{posts}</p>
        </div>

        <div className="groups p2">
          <p className="font-bold">Groups</p><p>{groups}</p>
        </div>
      </div>
    </section>
  )
}