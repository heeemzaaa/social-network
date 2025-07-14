import React, { useActionState, useState } from 'react';
import styles from "../../../(auth)/auth.module.css"
const initialPostData = {
    title: '',
    content: '',
};
// type two cases first post for user and post for groups
export default function CreatePost({ type, postAction }) {
    const [state, action] = useActionState(postAction, {});
    const [data, setData] = useState(initialPostData);

    const handleChange = (e) => {
        setData(prev => ({
            ...prev,
            [e.target.name]: e.target.value,
        }));
    };

    return (
        <form noValidate action={action} className={`${styles.form} glass-bg`}>
            <div className="flex gap-3">
                <div className="flex-col gap-1">

                    {/* Title */}
                    <div className={styles.formGrp}>
                        <label htmlFor="title">Title:</label>
                        <input
                            className={styles.input}
                            type="text"
                            name="title"
                            id="title"
                            value={data.title}
                            onChange={handleChange}
                        />
                        {state.errors?.title && <span className="field-error">{state.errors.title}</span>}
                    </div>

                    {/* Content */}
                    <div className={styles.formGrp}>
                        <label htmlFor="content">Content:</label>
                        <textarea
                            className={styles.input}
                            name="content"
                            id="content"
                            rows={5}
                            value={data.content}
                            onChange={handleChange}
                            placeholder="Write your post here..."
                        />
                        {state.errors?.content && <span className="field-error">{state.errors.content}</span>}
                    </div>

                    {/* Privacy */}
                    <div className={styles.formGrp}>
                        <label htmlFor="privacy">Privacy:</label>
                        <select
                            className={styles.input}
                            name="privacy"
                            id="privacy"
                            value={data.privacy}
                            onChange={handleChange}
                        >
                            <option value="public">Public</option>
                            <option value="private">Private</option>
                            <option value="friends">Friends Only</option>
                        </select>
                        {state.errors?.privacy && <span className="field-error">{state.errors.privacy}</span>}
                    </div>

                </div>

                <div className="flex-col gap-1">
                    {/* Image Upload */}
                    <div className={styles.formGrp}>
                        <label htmlFor="img">Image (Optional):</label>
                        <input
                            className={styles.input}
                            type="file"
                            name="img"
                            id="img"
                            accept="image/*"
                        />
                        {state.errors?.img && <span className="field-error">{state.errors.img}</span>}
                    </div>
                </div>
            </div>

            {/* Submit Button */}
            <button type="submit" className="btn-primary" disabled={state.pending}>
                {state.pending ? 'Submitting...' : 'Submit'}
            </button>

            {/* General messages */}
            {state.error && <span className="field-error">{state.error}</span>}
            {state.message && <span className="field-success">{state.message}</span>}
        </form>

    );
}

