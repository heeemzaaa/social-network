import React, { useActionState, useState } from 'react';

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
        <form action={action} >
            title : <input
                name="title"
                value={data.title}
                onChange={handleChange}
                placeholder="Title"
                className="input"
            />
            <br />
            {state.errors?.title && <span>{state.errors.title}</span>}
            <br />
            <label htmlFor="upload">GIF / IMG</label>
            <input name="img" id="upload" type="file" />
            <br />
            content : <textarea
                name="content"
                value={data.content}
                onChange={handleChange}
                placeholder="Write your post here..."
                className="textarea"
            />
            {state.errors?.content && <span>{state.errors.content}</span>}

            <button type="submit" className="btn-primary" disabled={state.pending}>
                {state.pending ? 'Submitting...' : 'Submit'}
            </button>

            {state.errors?.title && <p className="error">{state.errors.title}</p>}
            {state.errors?.content && <p className="error">{state.errors.content}</p>}
            {state.error && <p className="error">{state.error}</p>}
            {state.message && <p className="success">{state.message}</p>}
        </form>
    );
}

