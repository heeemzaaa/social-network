import "./components.css"

export default function InfosDiv({firstName, lastName, email, dateOfBirth, nickname}) {
  return (
    <div className="ProfileContainer p2" style={{backgroundImage:`url(https://www.portraitprofessionnel.fr/wp-content/uploads/2020/02/portrait-professionnel-corporate-4.jpg)`}} >
        {/* <img className="ProfileImg"></img> */}
        <div className="ProfileData">
              <span></span>First Name: {firstName}
              Last Name: {lastName}
              Email: {email}
              Date of Birth: {dateOfBirth}
              Nickname: {nickname}
        </div>
    </div>
  )
}