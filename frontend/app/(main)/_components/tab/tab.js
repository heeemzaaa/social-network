import "./style.css"
export default function  Tab({ label, isActive, onClick}) {
    return (
        <button
            className={`tab-button ${isActive ? 'active' : ''}`}
            onClick={onClick}
        >
            {label}
        </button>
    );
};
