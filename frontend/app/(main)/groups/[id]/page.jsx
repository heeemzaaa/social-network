"use client"; // Required for client-side rendering in Next.js

import React, { useEffect, useState } from "react";
import { HiMiniUsers } from "react-icons/hi2";
import Tag from "../../_components/tag";
import Tabs from "../../_components/tab/tabs";
import Tab from "../../_components/tab/tab";
import TabContent from "../../_components/tab/tabContent";
import GroupPostCardList from "../_components/groupPostCardList";
import GroupEventCardList from "../_components/groupEventCardList";

export default function GroupPage({ params }) {
  // const [groupId, setGroupId] = useState(null)
  const [data, setData] = useState({});
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState(null);

  const resolvedParams = React.use(params);
  const groupId = resolvedParams.id;

  // Fetch group data
  useEffect(() => {
    const getId = async () => {
      let { id } = await params
      return id
    }
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
        setData(result); // Set fetched data (e.g., { title, description, followers_number })
      } catch (error) {
        console.error("Error fetching group:", error);
        setError(error.message);
      } finally {
        setIsLoading(false);
      }
    };

    // getId().then(id => {
    //   getGroupData(id)
    //   setGroupId(id)
    //   console.log()
    // })

    getGroupData(groupId)
  }, []);


  // Render item for InfiniteList (customize based on your data structure)
  const renderItem = ({ item, index }) => (
    <div key={item.id || index} className="post-item p-4 border rounded">
      <h5>{item.title}</h5>
      <p>{item.content}</p>
    </div>
  );

  if (isLoading) return <p className="text-center">Loading...</p>;
  if (error) return <p className="text-danger text-center">Error: {error}</p>;

  return (
    <main className="group-page-section flex gap-2">
      <div className="col w-quarter" >
        {/* Group Info */}
        <div className="grp-info-container">
          <div className="grp-img w-full" >
            <img
              className="w-full"
              src={"/no-profile.png"}
              alt={data.title || "Group"}
            />
          </div>
          <div className="grp-info flex-col">
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
      <div className="flex-grow">
        <Tabs>
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