export default function GroupList({ groups, onGroupClick }) {
  return (
    <div className="pi3">
      {groups.map((grp, index) => (
        <>
          <div key={index} className="user_item p2 gap-1" onClick={() => onGroupClick(grp)} style={{ cursor: "pointer" }}>
            <img src={grp.image_path || "/no-profile.png"} />
            <p className="text-md">{grp.title}</p>
          </div>
          <div className="sep">

          </div>
        </>
      ))}
    </div>
  );
}