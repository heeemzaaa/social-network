"use client"; // Required for client-side rendering in Next.js

import "./style.css"
import Tag from "../../_components/tag";
import Tab from "../../_components/tab/tab";
import Tabs from "../../_components/tab/tabs";
import Avatar from "../../_components/avatar";
import Button from "@/app/_components/button";
import TabContent from "../../_components/tab/tabContent";
import { useModal } from "../../_context/ModalContext";
import CreatePostForm from "../_components/createPostForm";
import CreateEventForm from "../_components/createEventForm";
import InviteFriendForm from "../_components/inviteFriendsForm";
import GroupPostCardList from "../_components/groupPostCardList";
import GroupEventCardList from "../_components/groupEventCardList";
import { LuCalendarPlus } from "react-icons/lu";
import React, {
  useEffect,
  useState
} from "react";
import {
  HiMiniUsers,
  HiOutlineDocumentPlus,
  HiOutlineUserPlus
} from "react-icons/hi2";

export default function GroupPage({ params }) {
  const [data, setData] = useState({});
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState(null);
  const resolvedParams = React.use(params);
  const groupId = resolvedParams.id;
  const [isAccessible, setIsAccessible] = useState(null)

  const { openModal } = useModal()
  const actionButtons = [
    {
      label: "Invite Friend",
      icon: <HiOutlineUserPlus size={24} />,
      onClick: () => openModal(<InviteFriendForm groupId={groupId} />)
    },
    {
      label: "Add Post",
      icon: <HiOutlineDocumentPlus size={24} />,
      onClick: () => openModal(<CreatePostForm groupId={groupId} />)
    },
    {
      label: "Add Event",
      icon: <LuCalendarPlus size={"24"} />,
      onClick: () => openModal(<CreateEventForm groupId={groupId} />)
    }
  ]

  // Fetch group data
  useEffect(() => {
    const getGroupData = async (id) => {
      try {
        const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/groups/${id}`, {
          credentials: "include",
        });
        if (!response.ok) {
          throw new Error(`Failed to fetch group data: ${response.status}`);
        }
        const result = await response.json();
        setData(result); // Set fetched data (e.g., { title, description, followers_number })
      } catch (error) {
        console.error("Error fetching group:", error);
        setError(error.message);
      } finally {
        setIsLoading(false);
      }
    };
    getGroupData(groupId)
  }, []);

  if (isLoading) return <main className="text-center">Loading...</main>;
  if (error) return <p className="text-danger text-center">Error: {error}</p>;

  return (
    <main className="group-page-section flex gap-1">
      <div className="col" >

        {/* Group Info */}
        <Avatar img={data.image_path || null} size={250} />
        <div className="grp-info-container">
          <div className="grp-info flex-col align-center ">
            <h4 className="grp-title">
              {data.title}
            </h4>
            <p className="grp-description">
              {data.description}
            </p>
            <Tag className="glass-bg">
              <HiMiniUsers className="w-5 h-5" />
              {data.total_members || 0} {data.total_members > 1 ? "Members" : "Member"}
            </Tag>
          </div>
        </div>
      </div>

      {/* tabs for posts and  */}
      <div className="flex-grow h-full ">
        <div className="flex gap-1">
          {isAccessible?.status != 403 && actionButtons.map((button, index) =>
            <Button onClick={button.onClick} key={index}>
              {button.icon}
              <span style={{ marginLeft: "5px" }}>{button.label}</span>
            </Button>
          )}
        </div>

        <Tabs className={"h-full"}>
          <Tab label={"Posts"} />
          <Tab label={"Events"} />
          <TabContent>
            <GroupPostCardList groupId={groupId} setIsAccessible={setIsAccessible} isAccessible={isAccessible} />
          </TabContent>
          <TabContent>
            <GroupEventCardList groupId={groupId} setIsAccessible={setIsAccessible} isAccessible={isAccessible} />
          </TabContent>
        </Tabs>
      </div>
    </main>
  );
}