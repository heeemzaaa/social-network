"use client"
// import styles from "@/app/(auth)/auth.module.css"
import { HiMiniDocumentText } from "react-icons/hi2";
import { useActionState, useState } from "react";
import { createGroupAction } from "@/app/_actions/group";
import Button from "@/app/_components/button";
import styles from "@/app/page.module.css"

export default function CreateGroupForm() {
    const [state, action] = useActionState(createGroupAction, {});
    const [data, setData] = useState({
        title: "",
        description: "",
        img: null
    });

    const handleFileChange = (e) => {
        setData(prev => ({ ...prev, img: e.target.files[0] }));
    };

    return (
        <form action={action} className={`${styles.form} glass-bg`}>
            <div className={`${styles.formGrp}`}>
                <label htmlFor='title'>
                    <HiMiniDocumentText />
                    <span>Group Title:</span>
                </label>
                <input
                    className={`${styles.input}`}
                    id='title' 
                    name='title'
                    type='text'
                    placeholder='Group Title ...'
                    value={data.title}
                    onChange={(e) => setData(prev => ({ ...prev, title: e.target.value }))}
                />
                {state.errors?.title && <span className='field-error'>{state.errors.title}</span>}
            </div>
            <div className={`${styles.formGrp}`}>
                <label htmlFor='description'>
                    <HiMiniDocumentText />
                    <span>Description:</span>
                </label>
                <textarea
                    className={`${styles.input}`}
                    id='description'
                    name='description'
                    placeholder='Group Description ...'
                    value={data.description}
                    onChange={(e) => setData(prev => ({ ...prev, description: e.target.value }))}
                />
                {state.errors?.description && <span className='field-error'>{state.errors.description}</span>}
            </div>
            <div className={`${styles.formGrp}`}>
                <label htmlFor='img'>
                    {/* <HiMiniPhotograph /> */}
                    <span>Group Image (Optional):</span>
                </label>
                <input
                    className={`${styles.input}`}
                    id='img'
                    name='img'
                    type='file'
                    accept='image/jpeg,image/png,image/gif'
                    onChange={handleFileChange}
                />
                {state.errors?.img && <span className='field-error'>{state.errors.img}</span>}
            </div>
            <Button type={"submit"}>Create Group</Button>
            {state.error && <span className='field-error'>{state.error}</span>}
            {state.message && <span className='field-success'>{state.message}</span>}
        </form>
    )
}
