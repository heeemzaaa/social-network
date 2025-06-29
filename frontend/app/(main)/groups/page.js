'use client'

import Button from '@/app/_components/button'
import { HiOutlinePlus } from "react-icons/hi";
import { CreateGroup } from "./_components/create_group"
import { useState, useEffect, useRef } from 'react';


export default function Groups() {
  const [menu, showMenu] = useState(false)



  function HandleShowMenu() {
    console.log(menu, "heere");
    showMenu(!menu)
  }

   const myElementRef = useRef(null);

      useEffect(() => {
        const element = myElementRef.current; 
        console.log("heere inside the useeffect", element);
       

        if (element) {
          const handleClick = () => {
            HandleShowMenu()
          };

          element.addEventListener('click', handleClick);
          return () => {
            element.removeEventListener('click', handleClick);
          };
        }
      }, []); 
  return (
    <main >
      <Button className="align-center" onClick={HandleShowMenu}>
        <HiOutlinePlus />create group</Button>
      {menu && <CreateGroup ref={myElementRef} />}


    </main>
  )
}




