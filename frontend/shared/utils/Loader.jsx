/**
 * @author: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyright:  2024 
*/

import { useGlobalState } from '../../hooks/contex';
import styles from '../../css/loader.module.css';

export default () => {
    const { globalState } = useGlobalState();

    return (
        !globalState.isLoad &&
        <main className={styles.container_loader}>
           <span className={styles.loader}></span> 
        </main>
    );
}