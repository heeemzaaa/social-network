'use client'
import Button from "@/app/_components/button";
import Tag from "../../_components/tag";
import { HiMiniUsers } from "react-icons/hi2";
import "./style.css"
import { useRouter } from "next/navigation";
import { useState } from "react";
import { JoinGroupAction } from "@/app/_actions/group";
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

            if (!res.ok) console.error("Failed to send the request")

            requestState === 0 ? setRequestState(1) : setRequestState(0)
        } catch (err) {
            console.log(err);
        }
    }

    const navigateToGroup = (groupId) => {
        router.push(`/groups/${groupId}`);
    }


    return (
        <div className="grp-card w-quarter" onClick={() => {
            navigateToGroup(group_id)
        }}>
            <div className="grp-card-img-holder glass-bg">
                <div className="grp-card-img"
                    style={{ backgroundImage: image_path ? `url(http://localhost:8080/static/${image_path})` : `url('/no-profile.png')` }}

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
                                <Button className={"text-center"} onClick={(e) => handleJoingGrp()}>
                                    Join
                                </Button>
                            </div>
                            :
                            <div onClick={e => e.stopPropagation()}>
                                <Button  variant = 'btn-danger' className={"text-center"} onClick={(e) => handleJoingGrp(e)}>
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