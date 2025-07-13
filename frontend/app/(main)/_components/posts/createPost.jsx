import React, { useActionState, useState } from 'react'


// this component is for post creation form
// @type : can be ( user or grp ) 
// @action : the server action that will handle the form validation
// and send the validated data to the right endpoint
export default function CreatePost({ type, postAction }) {
    const [state, action] = useActionState(postAction, {});
    const [data, setData] = useState(initialPostData)

    return (
        <form action={action}>
            Here goes the creation post form <br />
            Can be used for either user or group posts <br />
            {action}
        </form>
    )
}

const initialPostData = {}