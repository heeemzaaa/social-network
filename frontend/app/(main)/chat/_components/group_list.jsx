import React from "react"
import Avatar from "../../_components/avatar";

export default function GroupList({ groups, onGroupClick }) {
  return (
    <div className="pi3">
      {groups.length > 0
        ? groups.map((grp, index) => (
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
        ))
        : <div style={{ width: "100%", maxWidth: "150px", display: "flex", flexDirection: "column", margin: "2rem auto", opacity: ".8", gap: "10px", }}>
          <img src="/search.png" />
          <p className="font-semibold" style={{ textAlign: "center" }} >You have no Groups, join some group.</p>
        </div>
      }
    </div>
  );
}
