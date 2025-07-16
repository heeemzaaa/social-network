// components/Tabs.jsx
import "./style.css";
import React, { useEffect, useState } from "react";

export default function Tabs({ children, className }) {
    console.log(`rendering tabs (instance: ${Math.random()})`)
    const [activeTab, setActiveTab] = useState(0);

    const handleTabClick = (index) => {
        setActiveTab(index);
    };

    const tabs = [];
    const contents = [];

    // useEffect(()=>{

        React.Children.forEach(children, (child) => {
            if (child && child.type && child.type.name === "Tab") {
                tabs.push(child);
            } else if (child && child.type && child.type.name === "TabContent") {
                contents.push(child);
            }
        })
    // })

    return (
        <div className={`tab-container ${className}`}>
            <h3>tab list: </h3>
            <div className="tab-list">
                
                {tabs.map((tab, index) =>
                    React.cloneElement(tab, {
                        isActive: index === activeTab,
                        onClick: () => handleTabClick(index),
                        key: index,
                    })
                )}
            </div>
            <h3>tab content</h3>
            <div className="tab-content">``
                {contents[activeTab] || <div>No content available for this tab</div>}
            </div>
        </div>
    );
}