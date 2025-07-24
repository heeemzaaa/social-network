import React from 'react'
import Avatar from '../../_components/avatar';

export default function UserCard({ user, isSelected, onSelect }) {
    return (
        <div
            style={{
                display: 'flex',
                alignItems: 'center',
                padding: '12px',
                borderRadius: '8px',
                border: isSelected ? '2px solid #C7DB7A' : '1px solid #e5e7eb',
                backgroundColor: isSelected ? 'rgba(199, 219, 122, .2)' : '#ffffff',
                cursor: 'pointer',
                transition: 'background-color 0.2s',
                marginBottom: '8px',
                width:'300px'
            }}

            onClick={() => onSelect(user.id)}
            onMouseOver={(e) => !isSelected && (e.currentTarget.style.backgroundColor = '#f9fafb')}
            onMouseOut={(e) => !isSelected && (e.currentTarget.style.backgroundColor = '#ffffff')}
        >
            <input
                name='userIds'
                value={user.id}
                type="checkbox"
                checked={isSelected}
                onChange={() => onSelect(user.id)}
                style={{ marginRight: '12px' }}
                hidden
            />

            <Avatar size={42} img={user.avatar} />
            <div>
                <p style={{ color: '#1f2937', fontWeight: '500', fontSize: '16px', marginLeft: "5px" }}>{user.firstname} {user.lastname}</p>
                <p>@{user.nickname}</p>
            </div>
        </div>
    );
};
