import "./chat.css"

export default function UserList({ users }) {
  return (
    <div className="pi3">
      {users.map((user, index) => (
        <>
          <div key={index} className="user_item p2 gap-1">
            <img src={user.img || "/no-profile.png"} />
            <p className="text-md">{user.username}</p>
          </div>
          <div className="sep"></div>
        </>
      ))}
    </div>
  );
}
