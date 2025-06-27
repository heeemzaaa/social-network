import React from 'react'
import InfosDiv from '../../_components/user_info';
import AboutUser from '../../_components/about_user';

export default async function Profile({ params }) {
  let { id } = await params

  const userInfos = {
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

  const aboutUser = {
    aboutMe: `Hi there! 👋 I'm someone who values meaningful connections,
              good conversations, and constant learning. I enjoy discovering new ideas,
              sharing moments with friends, and exploring what the world has to offer—online and off.
              Whether it's tech, travel, creativity, or just everyday thoughts,
              I'm here to be real and connect with others. Let's grow, laugh, and learn together!`
  }
  return (
    <main className='profile_page_section flex h-full p2 gap-4'>
      <InfosDiv {...userInfos} />
      <AboutUser {...aboutUser} />
    </main>
  );
}
