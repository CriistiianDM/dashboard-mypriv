/**
 * @author: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyright:  2024 
*/

// Imports
import React from 'react';
import { GlobalStateProvider } from '../hooks/contex';
import '../css/_general.css';

// _app
export default ({ Component, pageProps }) => {
  return (
    <GlobalStateProvider>
      <Component {...pageProps} />
    </GlobalStateProvider>
  );
};