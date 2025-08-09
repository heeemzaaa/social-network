import "./chat.css";
import React from 'react'


export default function UserList({ users, onUserClick }) {
  console.log("usssssssssssssers: ", users)
  return (
    <div className="pi3">
      {users.length > 0
        ? users?.map((user, index) => (
          <React.Fragment key={index}>
            <div key={index} className="user_item p2 gap-1" onClick={() => onUserClick(user)} style={{ cursor: "pointer" }}>
              <img src={user.img || "/no-profile.png"} />
              <p className="text-md">{user.username}</p>
            </div>
            <div className="sep"></div>
          </React.Fragment>
        ))
        : <div style={{width:"100%", maxWidth:"150px",  display:"flex", flexDirection:"column", margin:"2rem auto", opacity:".8", gap:"10px", }}>
            <img src="/search.png" />
            <p className="font-semibold" style={{textAlign:"center"}} >You have no friends, you need to follow others.</p>
        </div>
      }
    </div>
  );
}


