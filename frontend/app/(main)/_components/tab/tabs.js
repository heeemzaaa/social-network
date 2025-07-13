// components/Tabs.jsx
import "./style.css";
import React, { useState } from "react";

export default function Tabs({ children, className }) {
    const [activeTab, setActiveTab] = useState(0);

    const handleTabClick = (index) => {
        setActiveTab(index);
    };

    const tabs = [];
    const contents = [];
    React.Children.forEach(children, (child) => {
        if (child && child.type && child.type.name === "Tab") {
            tabs.push(child);
        } else if (child && child.type && child.type.name === "TabContent") {
            contents.push(child);
        }
    });

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
            <div className="tab-content ">
                {contents[activeTab] || <div>No content available for this tab</div>}
            </div>
        </div>
    );
}