'use client'

import Button from '@/app/_components/button'
import { HiOutlinePlus } from "react-icons/hi";
import styles from  "./group.module.css"
import { CreateGroup } from "./_components/createGroup"
import  {CardList, myItems} from "./_components/cardLists"
import { useState, useEffect, useRef } from 'react';


export default function Groups() {
  const [menu, showMenu] = useState(false)



  function HandleShowMenu() {
    console.log(menu, "before");
    showMenu(!menu)
    console.log(menu, "after");
  }

  // const myElementRef = useRef(null);

  // useEffect(() => {
  //   const element = myElementRef.current;
  //   console.log("hnaaaaaa", myElementRef.current);
  //   if (element) {
  //     const handleClick = () => {
  //       HandleShowMenu()
  //     };

  //     element.addEventListener('click', handleClick);
  //     return () => {
  //       element.removeEventListener('click', handleClick);
  //     };
  //   }
  // }, []);
  return (
    <main>
      <div className={`${styles.container_1}`}>
        <CardList  title={"My groups"} items={myItems} />
      </div>
      <div className={`${styles.container_2}`}>
        <div className={`${styles.create_bouton}`}>
          <Button className="justify-end" onClick={HandleShowMenu}>
            <HiOutlinePlus />
            create group
          </Button>
          {menu && <CreateGroup />}
        </div>
        <div className={`${styles.groups_unjoined}`}>
          
        </div>
      </div>
    </main>
  )
}




