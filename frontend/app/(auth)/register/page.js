import Logo from '@/app/_components/logo'
import styles from "../auth.module.css"
import { HiMiniUser,HiLockClosed } from "react-icons/hi2";
import Button from '@/app/_components/button'

import React from 'react'

export default function Register() {
  return (
    <div className='register'>
      <Logo />
      <form className={`${styles.form} glass-bg`}>
        <div className={`${styles.formGrp}`}>
          <label htmlFor='username'>
            <HiMiniUser />
            <span>
              User Name:
            </span>
          </label>
          <input className={`${styles.input}`} id='username' name='username' type='text' placeholder='User Name ...' />
          <span className='field-error'></span>
        </div>
        <div className={`${styles.formGrp}`}>
          <label htmlFor='password'>
            <HiLockClosed />
            <span>
              Password:
            </span>
          </label>
          <input className={`${styles.input}`} id='password' name='password' type='text' placeholder='Password ...' />
          <span className='field-error'></span>
        </div>

        <Button className="justify-center">Submit</Button>
      </form>
    </div>
  )
}
