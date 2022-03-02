import { useContext, useState } from 'react';
import update from 'immutability-helper';
import PropTypes from 'prop-types';

import AppStore from 'containers/App/context'
import { appActions } from 'containers/App/actions'
import Bookmark from 'components/Modal/Bookmark'
import InputBox from 'components/Modal/InputBox'

import MenuBarMemo, { MenuBarComponent } from './MenuBar';

const MenuBar = (props) => {
  const { data, setData } = useContext(AppStore);
  const app = appActions(data, setData)

  const [state] = useState(update(props, {data: {$merge: {
    ...data[props.key]
  }}}))

  state.data = update(state.data, {$merge: { ...data[state.key] }})
  state.bookmark = update(state.bookmark, {$merge: { ...data.bookmark }})

  state.getText = (key, defValue) => {
    return app.getText(key, defValue)
  }

  state.sideBar = () => {
    setData(state.key, { side: (state.data.side === "show") ? "hide" : "show" })
  }

  state.setScroll = () => {
    window.scrollTo(0,0);
  }

  state.loadModule = (key) => {
    switch (key) {
      case "login":
        return app.signOut()

      case "help":
        return app.showHelp("")

      case "bookmark":
        return state.showBookmarks()

      default:
        setData(state.key, { module: key, menu: "" }, ()=>{
          if(key === "setting" && !data.setting.group_key){
            setData(state.data.module, { group_key: "group_admin" }, ()=>{
              setData("current", { module: "setting", content: { type: 'setting' } })
            })
          }
        })
    }
  }

  state.showBookmarks = () => {
    setData(state.key, { modalForm:
      <Bookmark bookmark={state.bookmark} tabView={state.bookmarkView}
        getText={app.getText} 
        onSelect={(view, row) => {
          setData(state.key, { modalForm: null }, async () => {
            if((view === "bookmark") && (row.cfgroup === "browser")){
              let search_data = update(data.search, {
                filters: {$merge: {
                  [row.view]: row.filters
                }
              }})
              search_data = update(search_data, {
                columns: {$merge: {
                  [row.view]: row.columns
                }
              }})
              setData("current", { 
                module: "search", 
                content: [row.vkey, row.view, search_data]
              })
            } else {
              setData("current", { 
                module: "edit", 
                content: {
                  ntype: row.ntype, 
                  ttype: row.transtype, 
                  id: row.id,
                }
              })
            }
          })
        }}
        onDelete={(id) => {
          setData(state.key, { modalForm: 
            <InputBox 
              title={app.getText("msg_warning")}
              message={app.getText("msg_delete_text")}
              infoText={app.getText("msg_delete_info")}
              labelOK={app.getText("msg_ok")}
              labelCancel={app.getText("msg_cancel")}
              onCancel={() => {
                state.showBookmarks()
              }}
              onOK={async () => {
                const result = await app.requestData("/ui_userconfig", 
                  { method: "DELETE", query: { id: id } })
                if(result && result.error){
                  return app.resultError(result)
                }
                app.loadBookmark({user_id: data.login.data.employee.id, callback: ()=>{
                  state.showBookmarks()
                }})
              }}
            /> })
        }}
        onClose={()=>{
          setData(state.key, { modalForm: null })
        }}
      />
    })
  }

  return <MenuBarMemo {...state} />
}

MenuBar.propTypes = {
  key: PropTypes.string.isRequired,
  ...MenuBarComponent.propTypes,
  bookmark: PropTypes.object,
  bookmarkView: PropTypes.string.isRequired,
  showBookmarks: PropTypes.func,
}

MenuBar.defaultProps = {
  key: "current",
  ...MenuBarComponent.defaultProps,
  bookmark: {},
  bookmarkView: "bookmark",
  showBookmarks: undefined,
}

export default MenuBar;