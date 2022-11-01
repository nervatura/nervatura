import { APP_MODULE, SIDE_VISIBILITY, MENU_EVENT, MODAL_EVENT } from '../config/enums.js'

export class MenuController {
  constructor(host, app) {
    this.host = host;
    this.store = app.store;
    this.app = app;
    host.addController(this);
  }

  async deleteBookmark({ bookmark_id, menubar }){
    const { inputBox } = this.host
    const { data, setData, msg } = this.store
    const login = data[APP_MODULE.LOGIN]
    const  modalForm = inputBox({ 
      title: msg("", { id: "msg_warning" }), 
      message: msg("", { id: "msg_delete_text" }),
      infoText: msg("", { id: "msg_delete_info" }),
      labelCancel: msg("", { id: "msg_cancel" }),
      labelOK: msg("", { id: "msg_ok" }),
      onEvent: {
        onModalEvent: async ({ key }) => {
          setData("current", { modalForm: null })
          if(key === MODAL_EVENT.CANCEL){
            return menubar._onBookmark()
          }
          const result = await this.app.requestData("/ui_userconfig", 
            { method: "DELETE", query: { id: bookmark_id } })
          if(result && result.error){
            return this.app.resultError(result)
          }
          const bookmark = await this.app.loadBookmark({user_id: login.data.employee.id})
          return menubar._onBookmark(bookmark)
        }
      }
    })
    setData("current", { modalForm })
  }

  onModalEvent({key, data}){
    const { setData } = this.store
    switch (key) {
      case MODAL_EVENT.CANCEL:
        setData("current", { modalForm: null })
        break;

      case MODAL_EVENT.SELECTED:
        if((data.view === "bookmark") && (data.row.cfgroup === "browser")){
          const search_data = {
            ...this.store.data.search,
            filters: {
              ...this.store.data.search.filters,
              [data.row.view]: data.row.filters
            },
            columns: {
              ...this.store.data.search.columns,
              [data.row.view]: data.row.columns
            }
          }
          setData("current", {
            modalForm: null, 
            module: APP_MODULE.SEARCH, 
            content: {[APP_MODULE.SEARCH]: [data.row.vkey, data.row.view, search_data]} 
          })
        } else {
          setData("current", { 
            modalForm: null,
            module: APP_MODULE.EDIT, 
            content: {[APP_MODULE.EDIT]: {ntype: data.row.ntype, ttype: data.row.transtype, id: data.row.id}}
          })
        }
        break;

      case MODAL_EVENT.DELETE:
        this.deleteBookmark(data)
        break;
    
      default:
        break;
    }
  }

  onMenuEvent({key, data}){
    const { setData } = this.store
    const { current, setting } = this.store.data
    switch (key) {
      case MENU_EVENT.SIDEBAR:
        setData("current", {
          side: (current.side === SIDE_VISIBILITY.SHOW) ? SIDE_VISIBILITY.HIDE : SIDE_VISIBILITY.SHOW 
        })
        break;

      case MENU_EVENT.MODULE:
        switch (data.value) {
          case APP_MODULE.LOGIN:
            this.app.signOut()
            break;
          
          case APP_MODULE.HELP:
            this.app.showHelp("")
            break;
    
          case APP_MODULE.BOOKMARK:
            setData("current", { modalForm: data.modalForm })
            break;
        
          default:
            const values = { module: data.value, menu: "" }
            if(data.value === APP_MODULE.SETTING && !setting.group_key){
              setData(APP_MODULE.SETTING, { group_key: "group_admin" })
              values.content = { type: APP_MODULE.SETTING }
            }
            setData("current", { ...values })
            break;
        }
        break;

      case MENU_EVENT.SCROLL:
        window.scrollTo(0,0);
        break;

      default:
        break;
    }
  }
}