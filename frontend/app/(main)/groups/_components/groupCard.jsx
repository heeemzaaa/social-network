import Button from "@/app/_components/button";
import Tag from "../../_components/tag";
import { HiMiniUsers } from "react-icons/hi2";
import "./style.css"
export default function GroupCard({
    img,
    title,
    description,
    membersCount
}) {
    return (
        <div className="grp-card w-quarter">
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
                        {membersCount}
                    </Tag>
                </div>
                {/* <div>
                    <br /> */}
                    <Button className={"text-center"}>Join</Button>
                {/* </div> */}
            </div>
        </div>
    )
};

