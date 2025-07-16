import React, { useActionState } from 'react'

export default function CreateEventForm() {
    const [state, action] = useActionState(postAction, {});
    const [data, setData] = useState(initialPostData)

    return (
        <form action={action}>
            Here goes the creation group event form <br />
        </form>
    )
}

const initialPostData = {}
