"use client"
import styles from "@/app/page.module.css"
import { useActionState, useState } from "react";
import { registerUser } from "@/app/_actions/user";
import SubmitButton from "@/app/_components/subimtButton";
import Link from 'next/link';


const initialData = {
    email: "",
    password: "",
    firstname: "",
    lastname: "",
    birthdate: "",
    avatar: "",
    nickname: "",
    about_me: ""
};

export default function RegisterForm() {
    const [state, action] = useActionState(registerUser, {});
    const [data, setData] = useState(initialData);

    return (
        <form noValidate action={action} className={`${styles.form}`}>
            <div className="flex gap-3">
                <div className="flex-col gap-1">
                    {/* first name */}
                    <div className={styles.formGrp}>
                        <label htmlFor="firstname">First Name:</label>
                        <input
                            className={styles.input}
                            type="text"
                            name="firstname"
                            id="firstname"
                            value={data.firstname}
                            onChange={(e) => setData(prev => ({ ...prev, firstname: e.target.value }))}
                        />
                        {state.errors?.firstname && <span className="field-error">{state.errors.firstname}</span>}
                    </div>

                    {/* last name */}
                    <div className={styles.formGrp}>
                        <label htmlFor="lastname">Last Name:</label>
                        <input
                            className={styles.input}
                            type="text"
                            name="lastname"
                            id="lastname"
                            value={data.lastname}
                            onChange={(e) => setData(prev => ({ ...prev, lastname: e.target.value }))}
                        />
                        {state.errors?.lastname && <span className="field-error">{state.errors.lastname}</span>}

                    </div>

                    {/* birth date */}
                    <div className={styles.formGrp}>
                        <label htmlFor="birthdate">Date of Birth:</label>
                        <input
                            className={styles.input}
                            type="date"
                            name="birthdate"
                            id="birthdate"
                            value={data.birthdate}
                            onChange={(e) => setData(prev => ({ ...prev, birthdate: e.target.value }))}
                        />
                        {state.errors?.birthdate && <span className="field-error">{state.errors.birthdate}</span>}

                    </div>

                    {/* email */}
                    <div className={styles.formGrp}>
                        <label htmlFor="email">Email:</label>
                        <input
                            className={styles.input}
                            type="email"
                            name="email"
                            id="email"
                            value={data.email}
                            onChange={(e) => setData(prev => ({ ...prev, email: e.target.value }))}
                        />
                        {state.errors?.email && <span className="field-error">{state.errors.email}</span>}
                    </div>

                    {/* password */}
                    <div className={styles.formGrp}>
                        <label htmlFor="password">Password:</label>
                        <input
                            className={styles.input}
                            type="password"
                            name="password"
                            id="password"
                            value={data.password}
                            onChange={(e) => setData(prev => ({ ...prev, password: e.target.value }))}
                        />
                        {state.errors?.password && <span className="field-error">{state.errors.password}</span>}
                    </div>
                </div>

                {/* optional fields */}
                <div className="flex-col gap-1">
                    {/* avatar image */}
                    <div className={styles.formGrp}>
                        <label htmlFor="avatar">Avatar (Optional):</label>
                        <input
                            className={styles.input}
                            type="file"
                            name="avatar"
                            id="avatar"
                            accept="image/*"
                        />
                        {state.errors?.avatar && <span className="field-error">{state.errors.avatar}</span>}
                    </div>

                    {/* nickname */}
                    <div className={styles.formGrp}>
                        <label htmlFor="nickname">Nickname (Optional):</label>
                        <input
                            className={styles.input}
                            type="text"
                            name="nickname"
                            id="nickname"
                            value={data.nickname}
                            onChange={(e) => setData(prev => ({ ...prev, nickname: e.target.value }))}
                        />
                        {state.errors?.nickname && <span className="field-error">{state.errors.nickname}</span>}
                    </div>

                    {/* about me */}
                    <div className={styles.formGrp}>
                        <label htmlFor="about_me">About Me (Optional):</label>
                        <textarea
                            className={styles.input}
                            name="about_me"
                            id="about_me"
                            rows={4}
                            value={data.about_me}
                            onChange={(e) => setData(prev => ({ ...prev, about_me: e.target.value }))}
                        />
                        {state.errors?.about_me && <span className="field-error">{state.errors.about_me}</span>}
                    </div>
                </div>

            </div>
            <SubmitButton />
            {state.error && <span className="field-error">{state.error}</span>}
            {state.message && <span className="field-success">{state.message}</span>}
            <div className="text-sm font-medium ">
                <span>Don't have an account ? </span>
                <Link href={"/register"} style={{ color: "var(--color-primary)", textDecoration: "underline" }}
                > Register Now </Link>
            </div>
        </form>
    );
}
