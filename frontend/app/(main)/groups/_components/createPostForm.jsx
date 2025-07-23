
"use client";

import React, { useActionState, useEffect, useState } from "react";
import styles from "@/app/page.module.css"
import Button from "@/app/_components/button";
import { createGroupPostAction } from "@/app/_actions/group";
import { useModal } from "../../_context/ModalContext";

const initialData = {
    title: "",
    content: "",
};

export default function CreatePostForm({ groupId }) {
    const [state, action] = useActionState(createGroupPostAction, {});
    const [data, setData] = useState(initialData);
    const { setModalData, closeModal } = useModal()

    useEffect(() => {
        if (state.message) {
            state.data.type = "groupPost"
            setModalData(state.data)
            closeModal()
        }
    }, [state])


    return (
        <form noValidate action={action} className={`${styles.form} glass-bg`}>
            <input type="hidden" name="groupId" value={groupId} />
            <div className="flex flex-col gap-3">
                <div className={styles.formGrp}>
                    <label htmlFor="content">Content:</label>
                    <textarea
                        className={styles.input}
                        name="content"
                        id="content"
                        rows={5}
                        value={data.content}
                        onChange={(e) => setData((prev) => ({ ...prev, content: e.target.value }))}
                        placeholder="Enter post content"
                    />
                    {state.errors?.content && <span className="field-error">{state.errors.content}</span>}
                </div>
                <div className={styles.formGrp}>
                    <label htmlFor="image">Image (Optional):</label>
                    <input
                        className={styles.input}
                        type="file"
                        name="image"
                        id="image"
                        accept="image/*"
                    />
                    {state.errors?.image && <span className="field-error">{state.errors.avatar}</span>}
                </div>
                <input type="text" name="groupId" id="groupId" defaultValue={groupId} hidden />
                <Button>
                    Submit
                </Button>
                {state.error && <span className="field-error">{state.error}</span>}
                {state.message && <span className="field-success">{state.message}</span>}
            </div>
        </form>
    );
}