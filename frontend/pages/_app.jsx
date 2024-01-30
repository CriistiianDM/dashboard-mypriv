/**
 * @author: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyright:  2024 
*/

// Imports
import '../css/_general.css';

// Component Wrapper
const Wrapper = ({ 
  Component, 
  pageProps 
}) => 
{
  return <Component {...pageProps} />;
}

// Export
export default Wrapper;