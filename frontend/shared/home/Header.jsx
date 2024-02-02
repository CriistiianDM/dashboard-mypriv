/**
 * @author: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyright:  2024 
*/

import SearchHorizontal from '../../ui/SearchHorizontal';
import styles from '../../css/header.module.css';
import UserAccounts from '../../ui/UserAccounts';

// Header
export default () => {
    return (
        <header className={styles.containerHeader}>
            <SearchHorizontal />
            <UserAccounts />
        </header>
    );
}