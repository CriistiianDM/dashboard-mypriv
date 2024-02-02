/**
 * @author: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyright:  2024 
*/

import { getAllDataUser } from "../api/data";
import { useEffect , useState } from 'react';
import RenderImg from "../shared/utils/RenderImg";
import styles from '../css/header.module.css'

// User Accounts 
export default () => {
   const [allData , setAllData] = useState([]);

   const getData = async () => {
        const response = await getAllDataUser();
        setAllData(response?.data)
   }

   useEffect(() => {
     if (allData[0] == undefined) {
        getData()
     }
   }, []);

   return (
        <aside className={styles.containerProfiles}>
            {
             allData.length > 0 && 
             allData.map((e,index) => (                    
                <div key={index}>
                    <RenderImg url={e.profile_picture} alt={index} />
                </div>
             ))
            }
        </aside>
   );
}