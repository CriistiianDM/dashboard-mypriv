/**
 * @author: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyright:  2024 
*/

import SearchHorizontal from '../../ui/SearchHorizontal';
import styles from '../../css/header.module.css';
import UserAccounts from '../../ui/UserAccounts';
import RenderImg from '../utils/RenderImg';

// Header
export default () => {
    return (
        <header className={styles.containerHeader}>
            <section type="logo">
                <RenderImg alt={"logo"}/>
            </section>
            <section type="content">
                <SearchHorizontal />
                <UserAccounts />
            </section>
        </header>
    );
}