import { memo } from 'react';
import PropTypes from 'prop-types';

import MenuBar from 'components/MenuBar'

export const MenuBarComponent = ({
  data,
  getText, loadModule, sideBar, setScroll,
}) => {
  return <MenuBar
    side={data.side} scrollTop={data.scrollTop} module={data.module}
    loadModule={loadModule} getText={getText} sideBar={sideBar} setScroll={setScroll} />
}

MenuBarComponent.propTypes = {
  data: PropTypes.shape({
    side: MenuBar.propTypes.side,
    scrollTop: MenuBar.propTypes.scrollTop,
    module: MenuBar.propTypes.module,
  }).isRequired,
  loadModule: PropTypes.func, 
  getText: PropTypes.func, 
  sideBar: PropTypes.func, 
  setScroll: PropTypes.func,
}

MenuBarComponent.defaultProps = {
  data: {
    side: MenuBar.defaultProps.side,
    scrollTop: MenuBar.defaultProps.scrollTop,
    module: MenuBar.defaultProps.module,
  },
  loadModule: undefined, 
  getText: undefined, 
  sideBar: undefined, 
  setScroll: undefined,
}

export default memo(MenuBarComponent, (prevProps, nextProps) => {
  return (
    (prevProps.data.side === nextProps.data.side) &&
    (prevProps.data.scrollTop === nextProps.data.scrollTop) &&
    (prevProps.data.module === nextProps.data.module)
  )
})
