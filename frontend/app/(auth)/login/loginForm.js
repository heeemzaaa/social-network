"use client"
import styles from "@/app/page.module.css"
import { HiMiniUser, HiLockClosed } from "react-icons/hi2";
import { loginUser } from '@/app/_actions/user';
import { useActionState, useState } from "react";
import SubmitButton from "@/app/_components/subimtButton";


export default function LoginForm() {
    const [state, action] = useActionState(loginUser, {});
    const [data, setData] = useState({
        login: "",
        password: ""
    })

    return (
        <form action={action} className={`${styles.form} glass-bg`}>
            <div className={`${styles.formGrp}`}>
                <label htmlFor='login'>
                    <HiMiniUser />
                    <span>
                        User Name:
                    </span>
                </label>
                <input className={`${styles.input}`}
                    id='login'
                    name='login'
                    type='text'
                    placeholder='User Name ...'
                    value={data.login}
                    onChange={(e) => setData(prev => ({ ...prev, login: e.target.value }))}
                />
                {state.errors?.login && <span className='field-error'> {state.errors.login} </span>}
            </div>
            <div className={`${styles.formGrp}`}>
                <label htmlFor='password'>
                    <HiLockClosed />
                    <span>
                        Password:
                    </span>
                </label>
                <input className={`${styles.input}`}
                    id='password'
                    name='password'
                    type='password'
                    placeholder='Password ...'
                    value={data.password}
                    onChange={(e) => setData(prev => ({ ...prev, password: e.target.value }))}
                />
                {state.errors?.password && <span className='field-error'> {state.errors.password} </span>}
            </div>
            {state.error && <span className='field-error'> {state.error} </span>}
            {state.message && <span className='field-success'> {state.message} </span>}
            <SubmitButton className='btn-primary' />
        </form>
    )
}
