import Logo from '../../_components/logo'

import LoginForm from './loginForm';
import Link from 'next/link';

export default function Login() {
  return (
    <div className='login'>
      <Logo />
      <LoginForm />
      <span>Don&apost have an account ?  <Link href={"/register"}> Register </Link></span>
    </div>
  )
}
