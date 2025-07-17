import Button from "@/app/_components/button";
import Tag from "../../_components/tag";
import { HiMiniUsers } from "react-icons/hi2";
import "./style.css"
import { useRouter } from "next/navigation";
export default function GroupCard({
    type,
    group_id,
    image_path,
    title,
    description,
    total_members
}) {

    const router = useRouter()

    const handleJoingGrp = (groupId) => {
        console.log("join grp:) ")
    }


    const navigateToGroup = (groupId) => {
        router.push(`/groups/${groupId}`);
    }

    return (
        <div className="grp-card w-quarter" onClick={() => navigateToGroup(group_id)}>
            <div className="grp-card-img-holder glass-bg">
                <div className="grp-card-img" style={{ backgroundImage: `url(${'/no-profile.png'})` }}></div>
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
                        <Button className={"text-center"} onClick={(e) => {
                            e.stopPropagation()
                            handleJoingGrp(group_id)
                        }}>
                            Join
                        </Button>
                        :
                        <Button className={"text-center"}>Go to</Button>
                }
            </div>
        </div>
    )
}