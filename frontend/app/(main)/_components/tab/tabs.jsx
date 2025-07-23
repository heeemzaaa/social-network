// components/Tabs.jsx
import "./style.css";
import React, { useEffect, useState } from "react";
import Tab from "./tab";
import TabContent from "./tabContent";

export default function Tabs({ children, className }) {
    const [activeTab, setActiveTab] = useState(0);

    const handleTabClick = (index) => {
        setActiveTab(index);
    };

    const tabs = [];
    const contents = [];

    // useEffect(()=>{

        React.Children.forEach(children, (child) => {
            if (child && child.type === Tab) {
                tabs.push(child);
            } else if (child && child.type === TabContent) {
                contents.push(child);
            }
        })
    // })

    return (
        <div className={`tab-container ${className}`}>
            <div className="tab-list">
                
                {tabs.map((tab, index) =>
                    React.cloneElement(tab, {
                        isActive: index === activeTab,
                        onClick: () => handleTabClick(index),
                        key: index,
                    })
                )}
            </div>
            <div className="tab-content">
                {contents[activeTab] || <div>No content available for this tab</div>}
            </div>
        </div>
    );
}
