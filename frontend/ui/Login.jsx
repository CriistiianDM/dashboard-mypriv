/**
 * @author: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyright:  2024 
*/

import styles from "../css/home.module.css";

// Export Login
export default () => {
    return (
        <aside className={styles._containe_login}>
            <a type="content" id="google-login">
                <strong>Google</strong>
            </a>
            <h1 type="title">Inicia Sesion</h1>
        </aside>
    );
}