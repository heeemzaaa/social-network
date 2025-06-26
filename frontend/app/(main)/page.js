import InfosDiv from "./_components/user_info";

export default function Home() {
  const userInfos =  {
    firstName: "Hamza",
    lastName: "Elkhawlani",
    email: "hamza@gmail.com",
    dateOfBirth: "20-09-2000",
    nickname: "heeemzaaa",
    followers: 153,
    following: 147,
    posts: 52,
    groups: 18
  }
  return (
    <>
      <InfosDiv {...userInfos}/>
    </>
  );
}


