"use client"
import styles from "../auth.module.css"
import { HiMiniUser, HiLockClosed } from "react-icons/hi2";
import { loginUser } from '@/app/actions/user';
import { useActionState, useState } from "react";
import SubmitButton from "@/app/_components/subimtButton";

export default function LoginForm() {
    const [state, action] = useActionState(loginUser, {});
    const [data, setData] = useState({
        username: "",
        password: ""
    })

    return (
        <form action={action} className={`${styles.form} glass-bg`}>
            <div className={`${styles.formGrp}`}>
                <label htmlFor='username'>
                    <HiMiniUser />
                    <span>
                        User Name:
                    </span>
                </label>
                <input className={`${styles.input}`}
                    id='username'
                    name='username'
                    type='text'
                    placeholder='User Name ...'
                    value={data.username}
                    onChange={(e) => setData(prev => ({ ...prev, username: e.target.value }))}
                />
                {state.errors?.username && <span className='field-error'> {state.errors.username} </span>}
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
                    type='text'
                    placeholder='Password ...'
                    value={data.password}
                    onChange={(e) => setData(prev => ({ ...prev, password: e.target.value }))}
                />
                {state.errors?.password && <span className='field-error'> {state.errors.password} </span>}
            </div>
            {state.error && <span className='field-error'> {state.errors} </span>}
            {state.message && <span className='field-error'> {state.message} </span>}
            <SubmitButton className='btn-primary' />
        </form>
    )
}
