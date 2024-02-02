/**
 * @author: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyright:  2024 
*/

import Header from '../shared/home/Header';
import styles from '../css/home.page.module.css';

// Page Home
export default () => {
    return (
        <main className={styles.containerHomePage}>
            <Header />
        </main>
    );
}