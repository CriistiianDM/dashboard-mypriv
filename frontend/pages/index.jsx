/**
 * @author: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyright:  2024 
*/

import RenderImg from "../shared/utils/RenderImg";
import Login from "../ui/Login";
import styles from "../css/home.module.css";
import GoogleLogin from "../shared/google/GoogleLogin";
import Loader from "../shared/utils/Loader";
import { useEffect } from 'react';
import { useGlobalState } from '../hooks/contex';

/**
 * Route Page Init
 */
export default () => {
   //const
   const { setGlobalState } = useGlobalState();

   // Load Windows
   useEffect(() => {
      setGlobalState(prevState => ({
         ...prevState,
         isLoad: true,
       }));
   }, []);

   return (
      <main className={styles._container_home_primary}>
         <Loader />
         <RenderImg alt="LOGO" />
         <Login/>
         <GoogleLogin/>
      </main>
   );
}