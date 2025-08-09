"use client"

import Button from "@/app/_components/button";
import { useRef, useState } from "react";

export default function Popover({ trigger, children }) {
  const [isOpen, setIsOpen] = useState(false)
  const triggerRef = useRef(null)
  const popoverRef = useRef(null)

  // Toggle popover visibility
  const togglePopover = () => {
    setIsOpen((prev) => {
      const newState = !prev;
      return newState;
    });
  };

  return (
    <div className="relative inline-block">
      <Button ref={triggerRef} variant="btn-icon" onClick={togglePopover}>
        {trigger}
      </Button>
      {isOpen && (
        <div
          ref={popoverRef}
          className="popover-content glass-bg flex-col gap-0 absolute"
        >
          {children}
        </div>
      )}
    </div>
  )
}
