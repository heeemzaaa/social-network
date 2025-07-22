import Logo from '@/app/_components/logo'

import React from 'react'
import RegisterForm from './registerForm';

export default function Register() {
  return (
    <main className='register flex-col justify-center align-center' style={{borderRadius : "unset", margin: "unset"}}>
      <Logo />
      <RegisterForm />
    </main>
  )
}