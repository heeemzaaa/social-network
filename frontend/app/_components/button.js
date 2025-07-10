import React from 'react'
import './components.css'


export default function Button({children, variant = "btn-primary",type, onClick}, ref) {

   return (
    <button type={type} ref={ref} onClick={onClick} className={variant}>
      {children}
    </button>
  );
}
