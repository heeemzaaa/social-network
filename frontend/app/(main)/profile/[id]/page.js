import UserProfileWrapper from '../../_components/user_profile_wrapper'

export default function Profile({ params }) {
  return (
    <main className='profile_page_section flex h-full p4 gap-4'>
      <UserProfileWrapper params={params} />
    </main>
  )
}


