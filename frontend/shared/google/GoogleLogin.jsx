/**
 * @author: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyright:  2024 
*/

import { useEffect } from 'react';
import { decodeToken } from "react-jwt";
import { useGlobalState } from '../../hooks/contex';

function _0x416d($,x){let _=_0x7704();return(_0x416d=function($,x){return _[$-=488]})($,x)}const _0x180e23=_0x416d;function _0x7704(){let $=["7508473JsKwdt","05vgd0jgh3c69t","5718356pWKwGR","3057748jeCnaW","3346110qiFIdv","1139905RQVWjX","794518930528","6CDsTeh","13670noECOt","2098672bYhubb","39879wnPcET","1wVpQhp"];return(_0x7704=function(){return $})()}!function($,x){let _=_0x416d,e=$();for(;;)try{let t=-parseInt(_(488))/1*(parseInt(_(492))/2)+-parseInt(_(493))/3+-parseInt(_(491))/4+-parseInt(_(494))/5+-parseInt(_(496))/6*(parseInt(_(489))/7)+parseInt(_(498))/8+parseInt(_(499))/9*(parseInt(_(497))/10);if(945058===t)break;e.push(e.shift())}catch(n){e.push(e.shift())}}(_0x7704,945058);const PanconQueso=[_0x180e23(495),"-do01fjptc1i7dvoc70",_0x180e23(490)];

// Google Components
export default ({}) => {
    const _get_auth = async () => {
        try {
            const btnGoogle = document.getElementById("google-login");
            google.accounts.id.initialize({
            client_id: `${PanconQueso.join('')}.apps.googleusercontent.com`,
            callback: (response) => handleCredentialResponse(response),
            })
    
            google.accounts.id.renderButton(
            btnGoogle,
            { theme: "outline", size: "small" , text: "Google"}  
            );
    
            google.accounts.id.prompt();
        } catch (error) {
            console.log('error' , error);
        }
    }
    
    const onLoad = () => {
        alert("ha")
       //const
       const { setGlobalState } = useGlobalState();
       // Load Windows
        setGlobalState(prevState => ({
        ...prevState,
        isLoad: false,
        }));
    }
    
    const handleCredentialResponse = async (response) => {
        alert("antes")
        const data_decode = decodeToken(response.credential);
        console.log(data_decode)
    }

    useEffect( () => {
        const _root = document.body
  
        const _script = document.createElement('script');
        _script.src = 'https://accounts.google.com/gsi/client';
        _script.async = true;
        _script.id = 'google-login';
        _script.defer = true;
        _root.appendChild(_script);

        // Script Load
        _script.onload = () => {
            _get_auth();
        };
    }, []);

    return null;
}