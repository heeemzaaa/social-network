import Logo from '@/app/_components/logo'

import React from 'react'
import RegisterForm from './registerForm';
import Link from 'next/link';

export default function Register() {
  return (
    <div className='register'>
      <Logo />
      <RegisterForm />
      <span>Already have an accout ?  <Link href={"/login"} className='color-primary'> Login </Link></span>
    </div>
  )
}
