import React from 'react'
import './components.css'
import useDebouncedCallback from '../_hooks/useDebouncedCallback'

export default function Button({ children, variant = "btn-primary", type, onClick, ref, disabled = false, style }) {
  const debouncedFn = useDebouncedCallback(onClick)
  return (
    <button type={type} ref={ref} onClick={debouncedFn} className={variant} disabled={disabled} style={style}>
      {children}
    </button>
  );
}
