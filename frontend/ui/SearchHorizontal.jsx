/**
 * @author: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyright:  2024 
*/

import RenderImg from "../shared/utils/RenderImg";
import styles from '../css/search.module.css'
import { useState } from 'react';

// Search Horizontal
export default () => {

    const [isOpen , setOpen] = useState(false)
    const handleClickIcon = () => {
        setOpen(!isOpen)
    }

    // Render Component
    return (
        <aside className={styles.containerSearchHorizontal}>
            { isOpen && 
              <input type="text" placeholder="@nickname or email"/>
            }
            <div type="icon" onClick={handleClickIcon}>
                <RenderImg 
                    url="/assets/ei--search.svg"
                    alt="Search"
                />
            </div>
        </aside>
    );
}