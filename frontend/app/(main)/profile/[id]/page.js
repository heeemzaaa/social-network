import UserProfileWrapper from '../../_components/profile/user_profile_wrapper'

export default async function Profile({ params }) {
  let {id} = await params
  return (
    <main className='profile_page_section flex h-full p4 gap-4'>
      <UserProfileWrapper id={id} />
    </main>
  )
}


