import React from "react"
import Avatar from "./avatar"


export default function GroupList({ groups, onGroupClick }) {
  return (
    <div className="pi3">
      {groups.map((grp, index) => (
        <React.Fragment key={grp.id || index}>
          <div
            className="user_item p2 gap-1"
            onClick={() => onGroupClick(grp)}
            style={{ cursor: "pointer" }}
          >
            <Avatar img={grp.image_path} size="42" />
            <p className="text-md">{grp.title}</p>
          </div>
          <div className="sep"></div>
        </React.Fragment>
      ))}
    </div>
  );
}
