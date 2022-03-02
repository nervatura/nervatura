import React, { useContext, useState, useEffect } from 'react';
import PropTypes from 'prop-types';
import update from 'immutability-helper';

import AppStore from 'containers/App/context'
import { getSql, saveToDisk, appActions } from 'containers/App/actions'
import { searchActions } from './actions'
import { Queries } from 'containers/Controller/Queries'
import { Quick } from 'containers/Controller/Quick'
import Total from 'components/Modal/Total'
import { getSetting } from 'config/app'

import SearchMemo, { SearchView } from './Search';

const Search = (props) => {
  const { data, setData } = useContext(AppStore);
  const search = searchActions(data, setData)
  const app = appActions(data, setData)
  
  const [state] = useState(update(props, {$merge: {
    engine: data.login.data.engine,
    login: data.login.data,
    queries: Queries({ getText: app.getText }),
    quick: Quick({ getText: app.getText }),
    ui: getSetting("ui"),
    showHelp: app.showHelp
  }}))

  state.data = update(state.data, {$merge: { ...data[state.key] }})
  state.current = update(state.current, {$merge: { ...data.current }})

  useEffect(() => {
    if(state.current && state.current.content){
      const content = state.current.content
      setData("current", { content: null }, () => {
        search.showBrowser(...content)
      })
    }
  }, [setData, search, state]);

  state.getText = (key, defValue) => {
    return app.getText(key, defValue)
  }

  state.onEvent = (fname, params) => {
    state[fname](...params)
  }

  state.checkEditor = (params) => {
    setData("current", { module: "edit", content: params })
  }

  state.quickView = (qview) => {
    setData(state.key, { seltype: "selector",
      result: [], qview: qview, qfilter: "", page: 0 })
    setData("current", { side: "hide" })
  }

  state.editRow = (row, rowIndex) => {
    const params = row.id.split("/")
    const options = update({},{$set: { 
      ntype: params[0], 
      ttype: params[1], 
      id: parseInt(params[2],10), item:row 
    }})
    if (options.ntype === "servercmd") {
      search.showServerCmd(options.id)
    } else {
      setData("current", { module: "edit", content: options })
    }
  }

  state.setFormActions = (options) => {
    options.nextKey = "FORM_ACTIONS"
    setData("current", { module: "edit", content: options })
  }

  state.onEdit = ( fieldname, value, row ) => {
    const params = value.split("/")
    let options = update({},{$set: { 
      ntype: params[0], 
      ttype: params[1], 
      id: params[2] 
    }})
    if(fieldname === "id"){
      options = update(options, {$merge: {
        form: row.form,
        form_id: row.form_id
      }})
    }
    setData("current", { module: "edit", content: options })
  }

  state.onPage = (page) => {
    setData(state.key, { page: page })
  }

  state.changeData = (fieldname, value) => {
    setData(state.key, { [fieldname]: value })
  }

  state.setColumns = (fieldname, value) => {
    let columns = update(state.data.columns, {})
    if(value){
      columns = update(columns, { [state.data.view] : {$merge: {
        [fieldname]: true
      }} })
    } else {
      delete columns[state.data.view][fieldname]
      columns = update(columns, {$merge: {
        [state.data.view] : columns[state.data.view]
      } })
    }
    setData(state.key, { columns: columns, update: new Date().getTime() })
  }

  state.addFilter = () => {
    const viewDef = state.queries[state.data.vkey]()[state.data.view]
    const frow = viewDef.fields[Object.keys(viewDef.fields)[0]]
    let filters = update(state.data.filters, {})
    filters = update(filters, { [state.data.view]: {$push: [{
      id: new Date().getTime().toString(),
      fieldtype: frow.fieldtype,
      fieldname: Object.keys(viewDef.fields)[0],
      sqlstr: frow.sqlstr,
      wheretype: frow.wheretype,
      filtertype: "===",
      value: search.defaultFilterValue(frow.fieldtype)
    }]}})
    setData(state.key, { filters: filters })
  }

  state.deleteFilter = (index) => {
    let filters = update(state.data.filters, {})
    filters = update(filters, { [state.data.view]: {
      $splice: [[index, 1]]
    } })
    setData(state.key, { filters: filters })
  }

  state.editFilter = (index, fieldname, value) => {
    const viewDef = state.queries[state.data.vkey]()[state.data.view]
    let filters = update(state.data.filters, {})
    if((fieldname === "filtertype") || (fieldname === "value")){
      filters = update(filters, { [state.data.view]: {
        [index]: {$merge: { [fieldname]: value } }
      } })
    }
    if(fieldname === "fieldname"){
      if (Object.keys(viewDef.fields).includes(value)) {
        const frow = viewDef.fields[value]
        filters = update(filters, { [state.data.view]: {
          [index]: {$merge: { 
            fieldname: value, fieldtype: frow.fieldtype,
            sqlstr: frow.sqlstr, wheretype: frow.wheretype, filtertype: "===",
            value: search.defaultFilterValue(frow.fieldtype)
          }}
        }})
      } else {
        const deffield = state.data.deffield.filter((df)=>(df.fieldname === value))[0]
        const deftype = {
          bool: {
            fieldtype: "bool",
            sqlstr: "fg.groupvalue='bool' and case when fv.value='true' then 1 else 0 end "
          },
          integer: {
            fieldtype: "integer",
            sqlstr: "fg.groupvalue='integer' and {FMSF_NUMBER} {CAS_INT}fv.value {CAE_INT} {FMEF_CONVERT} "
          },
          float: {
            fieldtype: "float",
            sqlstr: "fg.groupvalue='float' and {FMSF_NUMBER} {CAS_FLOAT}fv.value {CAE_FLOAT} {FMEF_CONVERT} "
          },
          date: {
            fieldtype: "date",
            sqlstr: "fg.groupvalue='date' and {FMSF_DATE} {CASF_DATE}fv.value{CAEF_DATE} {FMEF_CONVERT} "
          },
          customer: {
            fieldtype: "string",
            sqlstr: "rf_customer.custname "
          },
          tool: {
            fieldtype: "string",
            sqlstr: "rf_tool.serial "
          },
          product: {
            fieldtype: "string",
            sqlstr: "rf_product.partnumber "
          },
          trans: {
            fieldtype: "string",
            sqlstr: "rf_trans.transnumber "
          },
          transitem: {
            fieldtype: "string",
            sqlstr: "rf_trans.transnumber "
          },
          transmovement: {
            fieldtype: "string",
            sqlstr: "rf_trans.transnumber "
          },
          transpayment: {
            fieldtype: "string",
            sqlstr: "rf_trans.transnumber "
          },
          project: {
            fieldtype: "string",
            sqlstr: "rf_project.pronumber "
          },
          employee: {
            fieldtype: "string",
            sqlstr: "rf_employee.empnumber "
          },
          place: {
            fieldtype: "string",
            sqlstr: "rf_place.planumber "
          }
        }
        let fieldtype = "string"; let sqlstr = "fv.value ";
        if(deftype[deffield.fieldtype]){
          fieldtype = deftype[deffield.fieldtype].fieldtype
          sqlstr = deftype[deffield.fieldtype].sqlstr
        }
        filters = update(filters, { [state.data.view]: {
          [index]: {$merge: { 
            fieldname: deffield.fieldname,
            fieldlimit: ["and","fv.fieldname","=","'"+deffield.fieldname+"'"],
            fieldtype: fieldtype, sqlstr: sqlstr,
            wheretype: "where", filtertype: "===", value: ""
          }}
        }})
      }
    }
    setData(state.key, { filters: filters })
  }

  state.showBrowser = (vkey, view) => {
    search.showBrowser(vkey, view, state.data)
  }

  state.quickSearch = async (filter) => {
    const view = await app.quickSearch(state.data.qview, filter)
    if(view.error){
      return app.resultError(view)
    }
    setData(state.key, { result: view.result, qfilter: filter, page: 0 })
  }

  state.browserView = async () => {
    const query = state.queries[state.data.vkey]()[state.data.view]
    let _sql = update({}, {$set: query.sql})
    let params = []
    let _where = []
    state.data.filters[state.data.view].filter((filter)=>(filter.wheretype === "where")).forEach(filter => {
      _where.push(search.getFilterWhere(filter))
      if(filter.filtertype !== "==N"){
        params.push(filter.value)
      }
    });
    if(_where.length > 0){
      _sql = update(_sql, { where: {$push: [..._where]}})
    }

    _where = []
    state.data.filters[state.data.view].filter((filter)=>(filter.wheretype === "having")).forEach(filter => {
      _where.push(search.getFilterWhere(filter))
      if(filter.filtertype !== "==N"){
        params.push(filter.value)
      }
    });
    if(_where.length > 0){
      _sql = update(_sql, { having: {$push: [..._where]}})
    }

    _where = app.getDataFilter(state.data.vkey, [], state.data.view)
    if(_where.length > 0){
      _sql = update(_sql, { where: {$push: [_where]}})
    }

    let userFilter = app.getUserFilter(state.data.vkey)
    if(userFilter.where.length > 0){
      _sql = update(_sql, { where: {$push: userFilter.where}})
      params = params.concat(userFilter.params)
    }

    let views = [
      { key: "result",
        text: getSql(state.engine, _sql).sql,
        values: params 
      }
    ]
    let options = { method: "POST", data: views }
    let view = await app.requestData("/view", options)
    if(view.error){
      return app.resultError(view)
    }
    setData(state.key, { result: view.result, dropdown: "", page: 0 })
  }

  state.showTotal = (fields, total) => {
    const deffield = Object.keys(fields).includes("deffield_value")
    const getValidValue = (value) => {
      if(isNaN(parseFloat(value))) {
        return 0
      } else {
        return parseFloat(value)
      }
    }
    state.data.result.forEach(row => {
      if (deffield) {
        if (typeof total.totalFields[row.fieldname] !== "undefined") {
          total = update(total, { 
            totalFields: { $merge: { 
              [row.fieldname]: total.totalFields[row.fieldname] + getValidValue(row.export_deffield_value) }}
          })
        }
      } else {
        Object.keys(total.totalFields).forEach(fieldname => {
          if (typeof row["export_"+fieldname] !== "undefined") {
            total = update(total, { 
              totalFields: { $merge: { 
                [fieldname]: total.totalFields[fieldname] + getValidValue(row["export_"+fieldname]) }}
            })
          } else {
            total = update(total, { 
              totalFields: { $merge: { 
                [fieldname]: total.totalFields[fieldname] + getValidValue(row[fieldname]) }}
            })
          }
        });
      }
    });
    setData("current", { modalForm: 
      <Total 
        total={total}
        getText={app.getText}
        onClose={() => {
          setData("current", { modalForm: null })
        }}
      /> 
    })
  }

  state.exportResult = (fields) => {
    const filename = state.data.view+".csv"
    let data = ""
    const labels = Object.keys(fields).map(fieldname => {
      return fields[fieldname].label
    })
    data += labels.join(state.ui.export_sep) + "\n"
    state.data.result.forEach(row => {
      const cols = Object.keys(fields).map(fieldname => {
        return (typeof row["export_"+fieldname] != "undefined") ? row["export_"+fieldname] : row[fieldname]
      })
      data += cols.join(state.ui.export_sep) + "\n"
    });
    const csvUrl = URL.createObjectURL(new Blob([data], {type : 'text/csv;charset=utf-8;'}))
    saveToDisk(csvUrl, filename)
  }

  state.bookmarkSave = () => {
    app.saveBookmark(['browser',state.queries[data.search.vkey]()[data.search.view].label])
  }

  return <SearchMemo {...state} />
}

Search.propTypes = {
  key: PropTypes.string.isRequired,
  ...SearchView.propTypes,
  showHelp: PropTypes.func,
  setFormActions: PropTypes.func,
  onEdit: PropTypes.func,
  changeData: PropTypes.func,
  setColumns: PropTypes.func,
  defaultFilterValue: PropTypes.func,
  addFilter: PropTypes.func,
  deleteFilter: PropTypes.func,
  editFilter: PropTypes.func,
  showBrowser: PropTypes.func,
  quickSearch: PropTypes.func,
  browserView: PropTypes.func,
  showTotal: PropTypes.func,
  exportResult: PropTypes.func,
  bookmarkSave: PropTypes.func
}

Search.defaultProps = {
  key: "search",
  ...SearchView.defaultProps,
  showHelp: undefined,
  setFormActions: undefined,
  onEdit: undefined,
  changeData: undefined,
  setColumns: undefined,
  defaultFilterValue: undefined,
  addFilter: undefined,
  deleteFilter: undefined,
  editFilter: undefined,
  showBrowser: undefined,
  quickSearch: undefined,
  browserView: undefined,
  showTotal: undefined,
  exportResult: undefined,
  bookmarkSave: undefined
  
}

export default Search;
