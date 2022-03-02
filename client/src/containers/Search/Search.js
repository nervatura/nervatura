import { memo, Fragment } from 'react';
import PropTypes from 'prop-types';

import SideBar from 'components/SideBar/Search'
import Selector from 'components/Modal/Selector'
import Browser from 'components/Browser'

export const SearchView = ({
  data, queries, quick, ui, login, current,
  onEvent, editRow, quickSearch, onPage, getText
}) => {
  return (
    <Fragment>
      <SideBar side={current.side} 
        groupKey={data.group_key} auditFilter={login.audit_filter} 
        onEvent={onEvent} getText={getText} />
      <div className={`${"page padding-normal"} ${current.theme}`} >
        {(data.seltype === "browser") ? 
          <Browser id={"browser_"+data.vkey} data={data}
            keyMap={queries[data.vkey]()} viewDef={queries[data.vkey]()[data.view]}
            paginationPage={ui.paginationPage} dateFormat={ui.dateFormat} 
            timeFormat={ui.timeFormat} filter_opt_1={ui.filter_opt_1} filter_opt_2={ui.filter_opt_2}
            onEvent={onEvent} getText={getText}/> : 
          <Selector id={"selector_"+data.qview} className="margin0"
            view={data.qview} columns={quick[data.qview]().columns}
            result={data.result} filter={data.qfilter} 
            currentPage={data.page} onCurrentPage={onPage}
            getText={getText} onClose={null} onSelect={editRow} onSearch={quickSearch} />}
      </div>
    </Fragment>
  )
}

SearchView.propTypes = {
  data: PropTypes.shape({
    ...Browser.propTypes.data,
    qview: PropTypes.string.isRequired,
    qfilter: PropTypes.string.isRequired
  }).isRequired,
  current: PropTypes.object.isRequired,
  queries: PropTypes.object.isRequired,
  quick: PropTypes.object.isRequired,
  ui: PropTypes.object.isRequired,
  login: PropTypes.object.isRequired,
  onEvent: PropTypes.func,
  editRow: PropTypes.func,
  quickSearch: PropTypes.func,
  onPage: PropTypes.func,
  getText: PropTypes.func
}

SearchView.defaultProps = {
  data: {
    ...Browser.defaultProps.data,
    qview: "",
    qfilter: ""
  }, 
  current: {},
  queries: {},
  quick: {},
  ui: {},
  login: {},
  onEvent: undefined,
  editRow: undefined,
  quickSearch: undefined,
  onPage: undefined,
  getText: undefined
}

export default memo(SearchView, (prevProps, nextProps) => {
  return (
    (prevProps.data === nextProps.data) &&
    (prevProps.current.side === nextProps.current.side) &&
    (prevProps.data.update === nextProps.data.update)
  )
})