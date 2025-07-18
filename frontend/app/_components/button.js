import React from 'react'
import './components.css'


export default function Button({ children, variant = "btn-primary", type, onClick, ref, disabled, style}) {

  return (
    <button type={type} ref={ref} onClick={onClick} className={variant} style={style} >
      {children}
    </button>
  );
}
