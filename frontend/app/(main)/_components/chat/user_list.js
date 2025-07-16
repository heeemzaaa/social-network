import "./chat.css";
import React from 'react'


export default function UserList({ users , onUserClick}) {
  return (
    <div className="pi3">
      {users.map((user, index) => (
        <React.Fragment key={index}>
          <div key={index} className="user_item p2 gap-1" onClick={() => onUserClick(user)} style={{ cursor: "pointer" }}>
            <img src={user.img || "/no-profile.png"} />
            <p className="text-md">{user.username}</p>
          </div>
          <div className="sep"></div>
        </React.Fragment>
      ))}
    </div>
  );
}


