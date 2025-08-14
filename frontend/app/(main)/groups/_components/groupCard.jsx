'use client'
import Button from "@/app/_components/button";
import Tag from "../../_components/tag";
import { HiMiniUsers } from "react-icons/hi2";
import "./style.css"
import { useRouter } from "next/navigation";
import { useState } from "react";
import { JoinGroupAction } from "@/app/_actions/group"; // should be used

import { useNotification } from "../../_context/NotificationContext";
import { useUserContext } from "../../_context/userContext";

export default function GroupCard({
    type,
    group_id,
    image_path,
    title,
    description,
    total_members,
    requested
}) {
    const router = useRouter()
    const [requestState, setRequestState] = useState(requested)
    const [error, setError] = useState(null)
    const { showNotification } = useNotification();
    const { sendSocketMessage } = useUserContext()


    // let's create here the function that toggles the state of the button with the same
    // way as hamza  
    async function handleJoingGrp() {
        let endpoint = `http://localhost:8080/api/groups/${group_id}/join-request`
        let method = requestState === 0 ? 'POST' : 'DELETE'
        try {
            const res = await fetch(endpoint, {
                method: method,
                headers: { 'Content-Type': 'application/json' },
                credentials: 'include',
            })

            const data = await res.json()
            if (!res.ok) {
                if (data.status === 401) {
                    router.push("/login")
                    return
                }
            }

            if (data.error === 'Already a member!') {
                showNotification({ Content: `You are already a member!`, Status: "warn" });
                console.warn(`You are already a member!`)
                router.push(`/groups/${group_id}`)
                return

            } else if (data.Message === 'ERROR!! Invitation not found') { // should .error
                console.warn(`Invitation not found !!`)
                showNotification({ Content: `Invitation not found`, Status: "warn" });
            }

            if (requestState === 0) {
                console.log("response data after join request", data)
                sendSocketMessage({
                    type: "notification",
                    notification: data,
                })
            }
            setRequestState(requestState === 0 ? 1 : 0)
        } catch (error) {
            setError(error)
            console.error(error);
        }
    }

    const navigateToGroup = (groupId) => {
        router.push(`/groups/${groupId}`);
    }

    if (error) {
        return (
            <Error error={error} />
        )
    }
    
    return (
        <div className="grp-card w-quarter" onClick={() => {
            navigateToGroup(group_id)
        }}>
            <div className="grp-card-img-holder glass-bg">
                <div className="grp-card-img"
                    style={{ backgroundImage: image_path ? `url(http://localhost:8080/static/${image_path})` : `url('/no-group.svg')` }}

                ></div>
            </div>
            <div className="grp-card-body flex-col justify-between gap-2">
                <div className="flex-col justify-evenly flex-grow">
                    <h3 className="grp-title">
                        {title}
                    </h3>
                    <p className="grp-description">
                        {description}
                    </p>
                    <Tag className={"glass-bg align-end"}>
                        <HiMiniUsers />
                        {total_members}
                    </Tag>
                </div>

                {
                    type === "available" ?
                        requestState == 0 ?
                            <div onClick={e => e.stopPropagation()}>
                                <Button className={"text-center"} onClick={(e) => handleJoingGrp(e)}>
                                    Join
                                </Button>
                            </div>
                            :
                            <div onClick={e => e.stopPropagation()}>
                                <Button variant='btn-danger' className={"text-center"} onClick={(e) => handleJoingGrp(e)}>
                                    Cancel
                                </Button>
                            </div>
                        :
                        <Button className={"text-center"}>Go to</Button>
                }
            </div>
        </div>
    )
}