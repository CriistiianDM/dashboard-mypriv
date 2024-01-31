/**
 * @author: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyright:  2024 
*/

// Render Img
const RenderImg = ({
    url,
    alt = "default"
}) => {
    const defaultImg = '/assets/white-logo.svg'
    return (
       <figure className="_default_img">
          <img src={url ?? defaultImg } alt={alt}/>
        </figure>
    );
}

// Export
export default RenderImg;