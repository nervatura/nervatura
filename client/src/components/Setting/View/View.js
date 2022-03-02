import { Fragment } from 'react';
import PropTypes from 'prop-types';
import update from 'immutability-helper';

import Icon from 'components/Form/Icon'
import Label from 'components/Form/Label'
import Table from 'components/Form/Table';
import List from 'components/Form/List';

import styles from './View.module.css';
import { getSetting } from 'config/app'

export const View = ({ 
  data, className,
  paginationPage, dateFormat, timeFormat,
  getText, onEvent,
  ...props 
}) => {
  const { caption, icon, view, actions, page } = data
  let fields = {}
  if(view.type === "table"){
    if(actions.edit){
      fields = update(fields, {$merge: {
        edit: { columnDef: {
          id: "edit",
          Header: "",
          headerStyle: {},
          Cell: ({ row, value }) => {
            return <Fragment>
              <div className={`${"cell"} ${styles.editCol}`} >
                <Icon id={"edit_"+row.original["id"]}
                  iconKey="Edit" width={24} height={21.3} 
                  onClick={(event) => {
                    event.stopPropagation();
                    onEvent("setViewActions", [actions.edit, row.original])
                  }}
                  className={styles.editCol} />
              </div>
            </Fragment>
          },
          cellStyle: { width: 30, padding: "7px 3px 3px 8px" }
        }}
      }})
    }
    fields = update(fields, {$merge: {...view.fields}})
  }
  const onPage = (page) => {
    onEvent("changeData",["page", page])
  }
  return (
    <div {...props} className={`${styles.width800} ${className}`}>
      <div className={`${styles.panel}`} >
        <div className={`${styles.panelTitle} ${"primary"}`}>
          <Label value={caption} 
            leftIcon={<Icon iconKey={icon} />} iconWidth="20px" />
        </div>
        <div className={`${"section-small"} ${styles.settingPanel}`} >
          <div className="row full container section-small-bottom" >
            <div className={`${styles.viewPanel}`} >
              <div className="row full" >
                {(view.type === "table")?
                <Table rowKey="id"
                  fields={fields} rows={view.result} tableFilter={true}
                  filterPlaceholder={getText("placeholder_filter")}
                  onAddItem={(actions.new) ? () => onEvent("setViewActions", [actions.new]) : undefined}
                  onRowSelected={(actions.edit) ? (row) => onEvent("setViewActions", [actions.edit, row]) : undefined}
                  labelAdd={getText("label_new")} labelYes={getText("label_yes")} labelNo={getText("label_no")} 
                  dateFormat={dateFormat} timeFormat={timeFormat} 
                  paginationPage={paginationPage} paginationTop={true}
                  currentPage={page} onCurrentPage={onPage} />:
                <List rows={view.result}
                  listFilter={true} filterPlaceholder={getText("placeholder_filter")}
                  onAddItem={(actions.new) ? () => onEvent("setViewActions", [actions.new]) : undefined}
                  labelAdd={getText("label_new")}
                  paginationPage={paginationPage} paginationTop={true} 
                  onEdit={(actions.edit) ? (row) => onEvent("setViewActions", [actions.edit, row]) : undefined} 
                  onDelete={(actions.delete) ? (row) => onEvent("setViewActions", [actions.delete, row]) : undefined}
                  currentPage={page} onCurrentPage={onPage} />}
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

View.propTypes = {
  data: PropTypes.shape({
    caption: PropTypes.string.isRequired,
    icon: PropTypes.string.isRequired,
    view: PropTypes.object.isRequired,
    actions: PropTypes.object.isRequired,
  }).isRequired,
  paginationPage: PropTypes.number.isRequired, 
  dateFormat: PropTypes.string, 
  timeFormat: PropTypes.string,
  className: PropTypes.string,
  onEvent: PropTypes.func,
  getText: PropTypes.func,
}

View.defaultProps = {
  data: {
    caption: "",
    icon: "",
    view: {},
    actions: {},
  },
  paginationPage: getSetting("paginationPage"),
  dateFormat: getSetting("dateFormat"),
  timeFormat: getSetting("timeFormat"),
  className: "",
  onEvent: undefined,
  getText: undefined,
}

export default View;