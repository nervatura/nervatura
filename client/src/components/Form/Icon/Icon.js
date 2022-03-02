import PropTypes from 'prop-types';

import Icons from './IconData.json'

export const ICON_KEY = Object.keys(Icons).sort()

export const Icon = ({
  iconKey, width, height, color,
  ...props 
}) => {
  return (
    <svg xmlns="http://www.w3.org/2000/svg" viewBox={Icons[iconKey].viewBox}
      {...props}
      width={width||Icons[iconKey].width} height={height||Icons[iconKey].height}>
      <g fill={color}>
        <path d={Icons[iconKey].path}/>
      </g>
    </svg>
  )
}


Icon.propTypes = {
  iconKey: PropTypes.oneOf(ICON_KEY).isRequired,
  width: PropTypes.number, 
  height: PropTypes.number, 
  color: PropTypes.string,
}

Icon.defaultProps = {
  iconKey: "ExclamationTriangle",
  width: undefined, 
  height: undefined, 
  color: undefined,
}

export default Icon;