import UserProfileWrapper from '../../_components/user_profile_wrapper'

export default async function Profile({ params }) {
  let id = await params.id
  return (
    <main className='profile_page_section flex h-full p4 gap-4'>
      <UserProfileWrapper id={id} />
    </main>
  )
}


