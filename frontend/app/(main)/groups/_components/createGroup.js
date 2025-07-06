import styles from "../group.module.css"
import Button from '@/app/_components/button'
import { HiMiniPencilSquare, HiOutlinePencil } from "react-icons/hi2";


export  function CreateGroup() {
    return (
        <div className='group-form'>
            <form className={`${styles.form} glass-bg`}>
                <div className={`${styles.formGrp}`}>
                    <label htmlFor='title'>
                         <HiOutlinePencil />
                        <span>
                            Title:
                        </span>
                    </label>
                    <input className={`${styles.input}`} id='title' name='title' type='text' placeholder='Title...' />
                    <span className='field-error'></span>
                </div>
                <div className={`${styles.formGrp}`}>
                    <label htmlFor='description'>
                        <HiMiniPencilSquare />
                        <span>
                            Description:
                        </span>
                    </label>
                    <textarea className={`${styles.input}`} id='description' name='description' type='text' placeholder='Description...' />
                    <span className='field-error'></span>
                </div>
                <Button className="justify-center">Submit</Button>
            </form>
        </div>
    );

}