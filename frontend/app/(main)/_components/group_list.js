export default function GroupList({ groups }) {
  return (
    <div className="pi3">
      {groups.map((grp, index) => (
        <>
          <div key={index} className="user_item p2 gap-1">
            <img src={grp.img || "/no-profile.png"} />
            <p className="text-md">{grp.name}</p>
          </div>
          <div className="sep">

          </div>
        </>
      ))}
    </div>
  );
}