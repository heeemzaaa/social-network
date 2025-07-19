import React from 'react'
import './components.css'


export default function Button({ children, variant = "btn-primary", type, onClick, ref, disabled = false, style }) {
  console.log(disabled)

  return (
    <button type={type} ref={ref} onClick={onClick} className={variant} disabled={disabled} style={style}>
      {children}
    </button>
  );
}
