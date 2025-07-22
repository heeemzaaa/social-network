import Image from 'next/image'
import React from 'react'

export default function Avatar({ img, size }) {

  let imgContainer = {
    width: `${size}px`,
    height: `${size}px`,
    borderRadius: "var(--border-radius-full)",
    overflow: "hidden",
    display: "flex",
    backgroundSize: "cover",
    backgroundRepeat: "no-repeat",
    backgroundPosition: "center",
  }

  let avatar = {
    padding: "3px",
    width: "max-content",
    height: "max-content",
    border: "solid 2px var(--color-primary)",
    borderRadius: "var(--border-radius-full)"
  }

  return (
    <div className='flex align-center justify-center glass-bg' style={avatar} >
      <div style={{ ...imgContainer, backgroundImage: img ? `url(http://localhost:8080/static${img})` : "url(/no-profile.png)" }} >
        {/* <img src={`http://localhost:8080/static${img}` || "/no-profile.png"} alt="" className="w-full" /> */}
      </div>
    </div>
  )
}




