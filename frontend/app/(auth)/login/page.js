import Logo from '@/app/_components/logo'
import styles from  "../auth.module.css"
import Button from '@/app/_components/button'
import { HiMiniUser,HiLockClosed } from "react-icons/hi2";

export default function Login() {
  return (
    <div className='login '>
      <Logo />
      <form className={`${styles.form} glass-bg`}>
        <div className={`${styles.formGrp}`}>
          <label> 
            <HiMiniUser/>
            <span>
            User Name:
            </span>
            
            </label>
          <input className={`${styles.input}`} name='userName' type='text' placeholder='User Name ...'/>
          <span className='field-error'></span>
        </div>
        <div className={`${styles.formGrp}`}>
          <label for='password'>
            <HiLockClosed/>
            <span>
            Password:  
            </span>
            </label>
          <input className={`${styles.input}`} id='password' name='password' type='text' placeholder='Password ...'/>
          <span className='field-error'></span>
        </div>

        <Button className="justify-center">Submit</Button>
      </form>
    </div>
  )
}
