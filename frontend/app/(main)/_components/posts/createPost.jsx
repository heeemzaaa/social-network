import React, { useActionState, useState, useEffect, use } from 'react';
import styles from "@/app/page.module.css"
import { useModal } from '../../_context/ModalContext';
import { useUserContext } from '../../_context/userContext';

// the intial post data
const initialPostData = {
    title: '',
    content: '',
    privacy: 'public',
    selectedFollowers: []
};
// component for creating posts
// use action state for form submissiom pecialized for running asynchronous actions (like postAction) and tracking their state
// use state to track  changed things
// use state for tracking user update  
// set data for setting daata whenever the user change somthing
export default function CreatePost({postAction}) {
    const [state, action] = useActionState(postAction, {});
    const [data, setData] = useState(initialPostData);
    const [followers, setFollowers] = useState([]);
    const [loadingFollowers, setLoadingFollowers] = useState(true);
    const {authenticatedUser} = useUserContext()

    // stores data that the modal will show when opened.
    const { setModalData, closeModal } = useModal()

    useEffect(() => {
        if (!state.data) return
        state.data.type = 'post';
        setModalData(state.data)
        closeModal()
    }, [state])

    // useEffect -> fetch followes olny when first rendring 
    useEffect(() => {
        const fetchFollowers = async () => {
            try {
                const res = await fetch(`http://localhost:8080/api/profile/${authenticatedUser.id}/connections/followers`, {
                    method: 'GET',
                    credentials: 'include',
                });
                const data = await res.json();
                setFollowers(data);
            } catch (err) {
                console.error("Error loading followers:", err);
            } finally {
                setLoadingFollowers(false);
            }
        };
        fetchFollowers();
    }, []);
    // set data when user change somthing
    const handleChange = (e) => {
        setData(prev => ({
            ...prev,
            [e.target.name]: e.target.value,
        }));
    };
    // If the follower is already selected, remove them (filter removes the id)
    // If the follower is not selected, add them ([...prev.selectedFollowers, followerId]).
    const handleFollowerToggle = (followerId) => {
        setData(prev => ({
            ...prev,
            selectedFollowers: prev.selectedFollowers.includes(followerId)
                ? prev.selectedFollowers.filter(id => id !== followerId)
                : [...prev.selectedFollowers, followerId]
        }));
    };
    // se;ect all deselect all
    const handleSelectAllFollowers = () => {
        if (!Array.isArray(followers)) return;
        setData(prev => ({
            ...prev,
            selectedFollowers: prev.selectedFollowers.length === followers.length
                ? []
                : followers.map(f => f.id)
        }));
    };
    // return the form 
    return (
        <form noValidate action={action} className={`${styles.form} glass-bg`}>
            <div className="flex gap-3">
                <div className="flex-col gap-1">
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
                            <option value="public">Public (All users can see)</option>
                            <option value="almost private">Almost Private (Only followers can see)</option>
                            <option value="private">Private (Selected followers only)</option>
                        </select>
                        {state.errors?.privacy && <span className="field-error">{state.errors.privacy}</span>}
                    </div>

                    {/* Follower Selection */}
                    {data.privacy === 'private' && (
                        <div className={styles.formGrp}>
                            <label>Select Followers:</label>
                            <div style={{
                                maxHeight: '200px',
                                overflowY: 'auto',
                                border: '1px solid #ddd',
                                borderRadius: '4px',
                                padding: '8px',
                                marginTop: '4px'
                            }}>
                                {loadingFollowers ? (
                                    <p style={{ fontStyle: 'italic', color: '#666' }}>Loading followers...</p>
                                ) : (
                                    <>
                                        {/* Select All */}
                                        <div style={{ borderBottom: '1px solid #eee', paddingBottom: '8px', marginBottom: '8px' }}>
                                            <label style={{ display: 'flex', alignItems: 'center', gap: '8px', cursor: 'pointer' }}>
                                                <input
                                                    type="checkbox"
                                                    checked={data.selectedFollowers.length === followers.length}
                                                    onChange={handleSelectAllFollowers}
                                                    style={{ cursor: 'pointer' }}
                                                />
                                                <strong>Select All ({followers.length})</strong>
                                            </label>
                                        </div>
                                        {/* Individual Followers */}
                                        {Array.isArray(followers) && followers.length > 0 && followers.map(follower => (
                                            <div key={follower.id} style={{ marginBottom: '8px' }}>
                                                <label style={{ display: 'flex', alignItems: 'center', gap: '8px', cursor: 'pointer' }}>
                                                    <input
                                                        type="checkbox"
                                                        checked={data.selectedFollowers.includes(follower.id)}
                                                        onChange={() => handleFollowerToggle(follower.id)}
                                                        style={{ cursor: 'pointer' }}
                                                    />
                                                    <div>
                                                        <div style={{ fontWeight: '500' }}>
                                                            {follower.fullname}
                                                        </div>
                                                        <div style={{ fontSize: '0.8em', color: '#666' }}>
                                                            {follower.nickname}
                                                        </div>
                                                    </div>
                                                </label>
                                            </div>
                                        ))}

                                        {state.errors?.selectedFollowers && <span className="field-error">{state.errors.selectedFollowers}</span>}
                                    </>
                                )}
                                {state.errors?.privacy && <span className="field-error">{state.errors.privacy}</span>}
                            </div>
                            {data.selectedFollowers.length > 0 && !loadingFollowers && (
                                <div style={{ marginTop: '8px', fontSize: '0.9em', color: '#666' }}>
                                    {data.selectedFollowers.length} follower(s) selected
                                </div>
                            )}
                        </div>
                    )}
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

            {/* Hidden input */}
            <input
                type="hidden"
                name="selectedFollowers"
                value={JSON.stringify(data.selectedFollowers)}
            />

            {/* Submit */}
            <button type="submit" className="btn-primary" disabled={state.pending}>
                {state.pending ? 'Submitting...' : 'Submit'}
            </button>

            {/* Messages */}
            {state.error && <span className="field-error">{state.error}</span>}
            {state.message && <span className="field-success">{state.message}</span>}
        </form>
    );
}
