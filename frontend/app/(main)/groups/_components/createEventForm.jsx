import { createGroupEventAction } from '@/app/_actions/group';
import { useActionState, useState } from 'react'
import styles from "@/app/page.module.css"
import Button from '@/app/_components/button';


const event = {
    title : "",
    description : "",
    date: ""
}

export default function CreateEventForm({groupId}) {
    console.log(groupId)
    const [state, action] = useActionState(createGroupEventAction, {});
    const [eventData, setEventData] = useState(event)
    return (
        <form noValidate action={action} className={`${styles.form} glass-bg w-full`}>
            <input type="hidden" name="groupId" value={groupId} />
            <div className="flex flex-col gap-3">
                <div className={styles.formGrp}>
                    <label htmlFor="content">Event title:</label>
                    <input
                        className={styles.input}
                        name="title"
                        id="title"
                        value={eventData.title}
                        onChange={(e) => setEventData((prev) => ({ ...prev, title: e.target.value }))}
                        placeholder="Enter post title"
                    />
                    {state.errors?.title && <span className="field-error">{state.errors.title}</span>}
                </div>
                <div className={styles.formGrp}>
                    <label htmlFor="description">Event description:</label>
                    <textarea
                        className={styles.input}
                        name="description"
                        id="description"
                        rows={5}
                        value={eventData.description}
                        onChange={(e) => setEventData((prev) => ({ ...prev, description: e.target.value }))}
                    />
                    {state.errors?.description && <span className="field-error">{state.errors.description}</span>}
                </div>
                <div className={styles.formGrp}>
                    <label htmlFor="date">Event Date:</label>
                    <input
                        className={styles.input}
                        name="date"
                        id="date"
                        type='datetime-local'
                        value={eventData.date}
                        onChange={(e) => setEventData((prev) => ({ ...prev, date: e.target.value }))}
                    />
                    {state.errors?.date && <span className="field-error">{state.errors.date}</span>}
                </div>

                <input type="text" name="groupId" id="groupId" defaultValue={groupId} hidden />
                <Button>
                    Submit
                </Button>
                {state.error && <span className="field-error">{state.error}</span>}
                {state.message && <span className="field-success">{state.message}</span>}
            </div>
        </form>
    )
}

const initialPostData = {}
