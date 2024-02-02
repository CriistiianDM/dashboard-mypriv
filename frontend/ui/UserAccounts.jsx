/**
 * @author: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyright:  2024 
*/

import { getAllDataUser } from "../api/data";
import { useEffect , useState } from 'react';

// User Accounts 
export default () => {
   const [allData , setAllData] = useState([]);

   const getData = async () => {
        const response = await getAllDataUser();
        setAllData(response?.data)
        console.log(response)
   }

   useEffect(() => {
     if (allData[0] == undefined) {
        getData()
     }
   }, []);

   return (
        <aside>
        </aside>
   );
}