import React from 'react'
import './components.css'


export default function Button({children, variant = "btn-primary", onClick}) {

   return (
    <button onClick={onClick} className={variant}>
      {children}
    </button>
  );
}
