import React from 'react'
import './components.css'


export default function Button({children, variant = "btn-primary", onClick}, ref) {

   return (
    <button ref={ref} onClick={onClick} className={variant}>
      {children}
    </button>
  );
}
