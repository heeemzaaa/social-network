"use client"; // Required for client-side rendering in Next.js

import React, { useEffect, useState } from "react";
import { HiMiniUsers, HiOutlineDocumentPlus, HiOutlineUserPlus } from "react-icons/hi2";
import { LuCalendarPlus } from "react-icons/lu";
import Tag from "../../_components/tag";
import Tabs from "../../_components/tab/tabs";
import Tab from "../../_components/tab/tab";
import TabContent from "../../_components/tab/tabContent";
import GroupPostCardList from "../_components/groupPostCardList";
import GroupEventCardList from "../_components/groupEventCardList";
import Avatar from "../../_components/avatar";
import Button from "@/app/_components/button";
import { useModal } from "../../_context/ModalContext";
import CreatePostForm from "../_components/createPostForm";
import CreateEventForm from "../_components/createEventForm";
import InviteFriendForm from "../_components/inviteFriendsForm";
import "./style.css"

export default function GroupPage({ params }) {
  const [data, setData] = useState({});
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState(null);
  const resolvedParams = React.use(params);
  const groupId = resolvedParams.id;

  console.log('inside the gtttt ', groupId);

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
        const response = await fetch(`http://localhost:8080/api/groups/${id}`, {
          credentials: "include",
        });
        if (!response.ok) {
          console.log("getting the data: ", await response.json())
          throw new Error(`Failed to fetch group data: ${response.status}`);
        }
        const result = await response.json();
        console.log(result)
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

  if (isLoading) return <p className="text-center">Loading...</p>;
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
          {actionButtons.map((button, index) =>
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
            <GroupPostCardList groupId={groupId} />
          </TabContent>
          <TabContent>
            <GroupEventCardList groupId={groupId} />
          </TabContent>
        </Tabs>
      </div>
    </main>
  );
}


created_at: "2025-07-19T11:55:34Z"
description: "darkmethoss-grp3"
group_creator: "ayoub nachti"
group_creator_id: "2ca83dd3-6e1a-4452-8653-fdea0b4d4417"
group_creator_nickname: "DarkMethoss"
group_id: "14aa33ae-aea8-43f0-99e9-ed6f6b197f03"
image_path: "/uploads/groups/1752926134233169971.jpg"
last_interaction: "0001-01-01T00:00:00Z"
title: "darkmethoss-grp3"
total_members: 1