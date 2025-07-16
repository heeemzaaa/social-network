
"use client";

import React, { useState } from "react";
import { useActionState } from "react-dom";
import { createPost } from "../_actions/post";
// import SubmitButton from "../_components/SubmitButton";
import styles from "../auth.module.css";
import Button from "@/app/_components/button";

const initialData = {
    title: "",
    content: "",
};

export default function CreatePostForm({ groupId }) {
    const [state, action] = useActionState(createPost, {});
    const [data, setData] = useState(initialData);

    return (
        <form noValidate action={action} className={`${styles.form} glass-bg`}>
            <input type="hidden" name="groupId" value={groupId} />
            <div className="flex flex-col gap-3">
                <div className={styles.formGrp}>
                    <label htmlFor="title">Post Title:</label>
                    <input
                        className={styles.input}
                        type="text"
                        name="title"
                        id="title"
                        value={data.title}
                        onChange={(e) => setData((prev) => ({ ...prev, title: e.target.value }))}
                        placeholder="Enter post title"
                    />
                    {state.errors?.title && <span className="field-error">{state.errors.title}</span>}
                </div>
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
                <Button>
                    Submit
                </Button>
                {state.error && <span className="field-error">{state.error}</span>}
                {state.message && <span className="field-success">{state.message}</span>}
            </div>
        </form>
    );
}