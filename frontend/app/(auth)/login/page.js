import Logo from '@/app/_components/logo'

import LoginForm from './loginForm';

export default function Login() {
  return (
    <main className='login flex-col justify-center align-center' style={{borderRadius : "unset", margin: "unset"}}>
      <Logo />
      <LoginForm />
    </main>
  )
}
