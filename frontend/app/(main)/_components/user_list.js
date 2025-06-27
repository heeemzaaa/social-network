export default function UserList({ users }) {
  return (
    <div>
      {users.map((user, index) => (
        <div key={index}>
          <img src={user.img || "/no-profile"} />
          <p className="text-md">{user.username}</p>
        </div>
      ))}
    </div>
  );
}
