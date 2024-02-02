/**
 * @author: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyright:  2024 
*/

import {
    insertNewAccount,
    getDataUserEmail
} from '../../api/data';
import CryptoJS from 'crypto-js';
import Cookies from 'js-cookie';

const _0x36547a=_0x19d1;function _0x19d1($,x){let _=_0x2da2();return(_0x19d1=function($,x){return _[$-=442]})($,x)}!function($,x){let _=_0x19d1,t=$();for(;;)try{let n=-parseInt(_(447))/1+-parseInt(_(445))/2+-parseInt(_(442))/3+parseInt(_(446))/4+parseInt(_(451))/5*(-parseInt(_(448))/6)+parseInt(_(444))/7*(parseInt(_(452))/8)+-parseInt(_(443))/9*(-parseInt(_(450))/10);if(233795===n)break;t.push(t.shift())}catch(c){t.push(t.shift())}}(_0x2da2,233795);const secrecToken=_0x36547a(449);function _0x2da2(){let $=["49xjEEIm","275918eviGjP","593724kFiWoF","15413OBLEFE","294VlUZwC","489438948sdjkdjdssnnayuvtyyiugyuy","1418430RUwjHx","13605XCCByt","125688iyraMZ","915852KJbrCu","36OSqUYy"];return(_0x2da2=function(){return $})()}

/***
 *  Validate login or create account
*/
export const loginCreateOrlogin = async ({data_decode}) => {
    return new Promise(async (resolve, reject) => {
        try {
            const 
            defaultArray = [],
            isRegisterUser = await getDataUserEmail({
                email: data_decode.email
            });
       
            if (isRegisterUser.status) {
                createCookkieLogin({
                    data: (isRegisterUser?.data)[0] ?? {}
                })
                defaultArray.push((isRegisterUser?.data)[0])
            } else {
                const result = await createUserAccount(data_decode);
                defaultArray.push(result)
            }

            resolve(defaultArray);
        } catch (error) {
            reject(error);
        }
    });
}

/***
 * Create User 
*/
const createUserAccount = async (data_decode) => {
    const data_ = []
    const response = await insertNewAccount({
        email: data_decode.email, 
        password: data_decode.jti.substring(0, 60), 
        nickname: data_decode.jti.substring(0, 50), 
        firstname: data_decode?.given_name, 
        lastname: data_decode?.family_name, 
        img: data_decode?.picture
    })

    if (response.status) {
        const dataUser = await getDataUserEmail({
            email: data_decode.email
        });
        data_.push((dataUser?.data)[0])
    }

    return data_
}

/***
 * Create Cokkie
*/
const createCookkieLogin = ({data}) => {
    const 
    data_ = JSON.stringify({ user: data  }),
    encryptedData = CryptoJS.AES.encrypt(data_, secrecToken).toString(),
    expirationDate = new Date();

    // set date in 15 days
    expirationDate.setDate(expirationDate.getDate() + 15);
    //create cokkie
    Cookies.set('_login', encryptedData, { expires: expirationDate });
}

/**
 * get Cokkie
*/
export const getCokkie = ({cokkie}) => {
    let
    encryptedData = Cookies.get(cokkie),
    jsonParse = '{}',
    decryptedBytes;

    if (encryptedData) {
        decryptedBytes = CryptoJS.AES.decrypt(encryptedData, secrecToken)
        jsonParse = decryptedBytes.toString(CryptoJS.enc.Utf8) ?? '{}';
    }

    return JSON.parse(jsonParse);
}