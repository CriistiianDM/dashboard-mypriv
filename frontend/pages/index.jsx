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
import { useRouter } from 'next/router';

/**
 * Route Page Init
 */
export default () => {
   //const
   const 
   { globalState , setGlobalState } = useGlobalState(),
   router = useRouter();

   // Load Windows
   useEffect(() => {
      // Validate if is loggued
      if (globalState.isLogged) {
          router.push('/home');
      }

      // Quit Loader
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