import React from 'react'

export default function Tag({ children, className}) {
    return (
        <div style={style} className={`${className}`} >
            {children}
        </div>
    )
}

const style = {
    // background : "#ddd",
    fontSize : "1rem",
    color: "black",
    width : "max-content",
    padding : ".2rem .5rem",
    display : "flex",
    alignItems :"center",
    gap : ".2rem",
    borderRadius : ".4rem"
}