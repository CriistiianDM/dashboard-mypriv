/**
 * @author: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyright:  2024 
*/

import json from '../json/endpoints.json';
import {
    fetchPostGeneral
} from './fetch'

/**
 * Create New User
 * 
 * @param {*} param0 
 * @returns 
 */
export const insertNewAccount = (data) => {
    return fetchPostGeneral({
        dataSend: data,
        urlEndPoint: json.createUser
    })
}

/**
 * Get data of user per email
 * 
 * @param {*} param0 
 * @returns 
 */
export const getDataUserEmail = (data) => {
    return fetchPostGeneral({
        dataSend: data,
        urlEndPoint: json.getUserEmail
    })
}

/**
 * Get data of user per email
 * 
 * @param {*} param0 
 * @returns 
 */
export const getAllDataUser = () => {
    return fetchPostGeneral({
        dataSend: {},
        urlEndPoint: json.getAllUser
    })
}